// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with bottos.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  producer entry
 * @Author:
 * @Date:   2017-12-06
 * @Last Modified by:
 * @Last Modified time:
 */

package producer

import (
	"math"

	log "github.com/cihub/seelog"

	"github.com/bottos-project/bottos/common"
	"github.com/bottos-project/bottos/config"
	"github.com/bottos-project/bottos/context"
)

//ReportState is recording the state of reporters
type ReportState struct {
	ScheduledTime     uint64
	ScheduledReporter string
	PubKey            string
	IsReporting       bool
	CheckFlag         uint32
	ProtocolInterface context.ProtocolInterface
}

//IsReady is check if repoter state
func (r *Reporter) IsReady() bool {
	now := GetReportTimeNow()
	r.state.SetCheckFlag(1)

	if config.BtoConfig.Delegate.Solo == false && r.roleIntf.IsChainActivated() == false {
		//log.Debug("PRODUCER node is not produce transfered")
		return false
	} else if config.BtoConfig.Delegate.Solo == false {
		if r.IsSynced() == false {
			log.Debugf("PRODUCER p2p is syncing")
			return false
		}
	}

	slot := r.roleIntf.GetSlotAtTime(now)
	if slot == 0 {
		return false
	}

	object, err := r.roleIntf.GetChainState()
	if err != nil {
		log.Errorf("PRODUCER GetChainState failed %v", err)
		return false
	}
	if (now < object.LastBlockTime+uint64(config.DEFAULT_BLOCK_INTERVAL)) && object.LastBlockNum != 0 {
		log.Debugf("PRODUCER time not ready", now, object.LastBlockTime, uint64(config.DEFAULT_BLOCK_INTERVAL))
		return false
	}
	if r.IsMyTurn(now, slot) == false {
		return false
	}
	return true

}

//GetReportTimeNow is to count reporter's time
func GetReportTimeNow() uint64 {
	systemNow := common.NowToMicroseconds()
	nowMicro := common.Microsecond{}
	nowMicro.Count = (systemNow + config.DEFALT_SLOT_CHECK_INTERVAL)
	now := common.ToSeconds(nowMicro)
	return now
}

//StartReport is to start
func (r *Reporter) StartReport() {
	r.state.IsReporting = true
}

//EndReport is to stop report
func (r *Reporter) EndReport() {
	r.state.IsReporting = false
}

//SetCheckFlag is to set check flags
func (r *ReportState) SetCheckFlag(flag uint32) {
	r.CheckFlag |= flag
}

//IsSynced is to check synced flags
func (r *Reporter) IsSynced() bool {
	if r.state.ProtocolInterface.GetBlockSyncState() == true {
		return true
	}
	return false
}

//IsMyTurn is to check if is my turn
func (r *Reporter) IsMyTurn(startTime uint64, slot uint64) bool {
	accountName, err := r.roleIntf.GetCandidateBySlot(slot)
	if err != nil {
		log.Debugf("PRODUCER cannot get delegate by slot:%v,err:%v", slot, err)
		return false
	}
	if r.roleIntf.IsAccountExist(accountName) == false {
		log.Debugf("PRODUCER account not exist, account %s", accountName)
		return false
	}

	scheduledTime := r.roleIntf.GetSlotTime(slot)
	delegate, err := r.roleIntf.GetDelegateByAccountName(accountName)
	if err != nil {
		log.Debugf("PRODUCER find delegate by account failed %s", accountName)
		return false
	}
	_, err = config.GetDelegateSignKey(delegate.ReportKey)
	if err != nil {
		return false
	}

	if accountName != config.BtoConfig.Delegate.Account {
		return false
	}

	prate := r.roleIntf.GetDelegateParticipationRate(accountName)

	if prate < config.DELEGATE_PATICIPATION {
		return false
	}

	if math.Abs(float64(scheduledTime)*1000-float64(startTime)*1000) > float64(config.PRODUCER_TIME_OUT*10) {
		log.Debugf("PRODUCER delegate is too slow %s,%v,%v", accountName, scheduledTime, startTime)
		return false
	}
	r.state.ScheduledTime = scheduledTime
	r.state.ScheduledReporter = accountName
	r.state.PubKey = delegate.ReportKey

	return true
}
