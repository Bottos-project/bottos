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
 * file description:  account role
 * @Author: may luo
 * @Date:   2017-12-12
 * @Last Modified by:
 * @Last Modified time:
 */
package role

import (
	"math/big"

	"encoding/json"
	"strconv"

	"github.com/bottos-project/bottos/db"
	log "github.com/cihub/seelog"
)

// TransitVotesObjectName is definition of delegate vote object name in transition period
const TransitVotesObjectName string = "transitvotes"

// TransitVotesObjectKeyName is definition of delegate vote object key name in transition period
const TransitVotesObjectKeyName string = "producer_account"

// TransitVotesObjectIndexVote is definition of delegate vote object index name in transition period
const TransitVotesObjectIndexVote string = "transit_votes"
const TransitVotesObjectIndexVoteJSON string = "transit_votes"

//TransitVotes is definition of delegate votes
type TransitVotes struct {
	ProducerAccount string   `json:"producer_account"`
	TransitVotes    *big.Int `json:"transit_votes"`
}

// CreateTransitVotesRole is to save initial delegate votes in transition period
func CreateTransitVotesRole(ldb *db.DBService) error {
	err := ldb.CreatObjectIndex(TransitVotesObjectName, TransitVotesObjectKeyName, TransitVotesObjectKeyName)
	if err != nil {
		return err
	}
	err = ldb.CreatObjectMultiIndex(TransitVotesObjectName, TransitVotesObjectIndexVote, TransitVotesObjectIndexVoteJSON, TransitVotesObjectKeyName)
	if err != nil {
		return err
	}
	ldb.AddObject(TransitVotesObjectName)
	return nil
}

// SetTransitVotesRole is to save delegate votes in transition period
func SetTransitVotesRole(ldb *db.DBService, key string, value *TransitVotes) error {
	jsonvalue, err := json.Marshal(value)
	if err != nil {
		log.Error("ROLE Marshal failed ", err)
		return err
	}

	return ldb.SetObject(TransitVotesObjectName, key, string(jsonvalue))
}


// GetTransitVotesRole is to get delegate votes by account name
func GetTransitVotesRole(ldb *db.DBService, key string) (*TransitVotes, error) {

	value, err := ldb.GetObject(TransitVotesObjectName, key)
	if err != nil {
		return nil, err
	}
	res := &TransitVotes{}
	err = json.Unmarshal([]byte(value), res)
	if err != nil {
		log.Error("ROLE Unmarshal failed ", err)
		return nil, err
	}

	return res, nil

}