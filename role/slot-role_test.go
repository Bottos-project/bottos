﻿// Copyright 2017~2022 The Bottos Authors
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
 * file description: producer
 * @Author: May Luo
 * @Date:   2017-12-11
 * @Last Modified by:
 * @Last Modified time:
 */
package role

import (
	"testing"
	"time"

	log "github.com/cihub/seelog"

	"github.com/bottos-project/bottos/common"
	"github.com/bottos-project/bottos/db"
)

func startup() RoleInterface {
	//config.LoadConfig()
	dbInst := db.NewDbService("./temp/db", "./temp/codedb")
	if dbInst == nil {
		log.Error("Create DB service fail")
	}
	roleIntf := NewRole(dbInst)

	return roleIntf
}

func TestReporter_GetSlotAtTime(t *testing.T) {
	ins := startup()
	cbegin := time.Time{}
	slot := ins.GetSlotAtTime(uint64(cbegin.Unix()))
	log.Info(slot)
	cUnix := cbegin.Unix()
	log.Info(cUnix)
	//	slot = ins.GetSlotAtTime(cUnix)
	//	log.Info(slot)
	now := common.NowToSeconds()
	slot = ins.GetSlotAtTime(now)
	log.Info(slot)

	nowMicroSec := common.NowToSlotSec(time.Now(), 500000)
	slot = ins.GetSlotAtTime(nowMicroSec)
	log.Info(slot)

}
