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
 * file description:  net actor
 * @Author:
 * @Date:   2017-12-06
 * @Last Modified by:
 * @Last Modified time:
 */

package netactor

import (
	log "github.com/cihub/seelog"
	//"encoding/json"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/bottos-project/bottos/action/env"
	"github.com/bottos-project/bottos/action/message"
	"github.com/bottos-project/bottos/config"
	"github.com/bottos-project/bottos/context"
	netprotocol "github.com/bottos-project/bottos/protocol"
)

type NetActor struct {
	actorEnv *env.ActorEnv
	protocol context.ProtocolInstance
}

var netactor *NetActor

func NewNetActor(env *env.ActorEnv) *actor.PID {

	netactor = &NetActor{
		actorEnv: env,
	}

	props := actor.FromProducer(func() actor.Actor { return netactor })

	pid, err := actor.SpawnNamed(props, "NetActor")
	if err == nil {

		netactor.protocol = netprotocol.MakeProtocol(config.Param, env.Chain)
		netactor.protocol.Start()

		env.Protocol = netactor.protocol
		return pid
	} else {
		panic(log.Errorf("NetActor SpawnNamed error: ", err))
	}

	return nil
}

//main loop
func (n *NetActor) handleSystemMsg(context actor.Context) {
	switch msg := context.Message().(type) {

	case *actor.Started:
		log.Infof("NetActor received started msg", msg)

	case *actor.Stopping:
		log.Info("NetActor received stopping msg")

	case *actor.Restart:
		log.Info("NetActor received restart msg")

	case *actor.Restarting:
		log.Info("NetActor received restarting msg")

	case *message.NotifyTrx:
		n.protocol.ProcessNewTrx(msg)

	case *message.NotifyBlock:
		n.protocol.ProcessNewBlock(msg)

	default:
		log.Error("netactor receive unknown message")
	}

}

func (n *NetActor) Receive(context actor.Context) {
	n.handleSystemMsg(context)
}

func (n *NetActor) setChainActor(tpid *actor.PID) {
	n.protocol.SetChainActor(tpid)
}

func (n *NetActor) setTrxActor(tpid *actor.PID) {
	n.protocol.SetTrxActor(tpid)
}

func (n *NetActor) setProducerActor(tpid *actor.PID) {
	n.protocol.SetTrxActor(tpid)
}

func SetChainActorPid(tpid *actor.PID) {
	netactor.setChainActor(tpid)
}

func SetTrxActorPid(tpid *actor.PID) {
	netactor.setTrxActor(tpid)
}

func SetProducerActorPid(tpid *actor.PID) {
	netactor.setProducerActor(tpid)
}
