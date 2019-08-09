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
 * file description:  account role
 * @Author: Gong Zibin
 * @Date:   2017-12-12
 * @Last Modified by:
 * @Last Modified time:
 */

package role

import (
	"encoding/json"
	"github.com/bottos-project/bottos/db"
	log "github.com/cihub/seelog"
)

const (
	//AccountObjectName is the table name of user account
	AccountObjectName string = "account"
	//MsignAccountObjectName is the table name of user multisign account
	MsignAccountObjectName string = "msignaccount"
)

// Account is definition of user account
type Account struct {
	AccountName  string                                      `json:"account_name"`
	PublicKey    []byte                                      `json:"public_key"`
	CreateTime   uint64                                      `json:"create_date"`
	GsPermission bool                                        `json:"gs_permission"`
	ContractName []string 									 `json:"contract_name"`
}

type MsignAccountAuthority struct {
	AuthorAccount string `json:"author_account"`
	Weight        uint32 `json:"weight"`
}

// MsignAccount is definition of user multisign account
type MsignAccount struct {
	MsignAccountName string                  `json:"msign_account_name"`
	Authority        []MsignAccountAuthority `json:"authority"`
	Threshold        uint32                  `json:"threshold"`
}

// CreateAccountRole is create account role
func CreateAccountRole(ldb *db.DBService) error {
	ldb.AddObject(AccountObjectName)
	return nil
}

// CreateMsignAccountRole is create account role
func CreateMsignAccountRole(ldb *db.DBService) error {
	ldb.AddObject(MsignAccountObjectName)
	return nil
}

func accountNameToKey(name string) string {
	return name
}

// SetAccountRole is common func to set role for account
func SetAccountRole(ldb *db.DBService, accountName string, value *Account) error {
	key := accountNameToKey(accountName)
	jsonvalue, err := json.Marshal(value)
	if err != nil {
		log.Errorf("DB Marshal failed %v,%v", accountName, err)
		return err
	}
	return ldb.SetObject(AccountObjectName, key, string(jsonvalue))
}

// GetAccountRole is common func to get role for account
func GetAccountRole(ldb *db.DBService, accountName string) (*Account, error) {
	key := accountNameToKey(accountName)
	value, err := ldb.GetObject(AccountObjectName, key)
	if err != nil {
		return nil, err
	}

	res := &Account{}
	err = json.Unmarshal([]byte(value), res)
	if err != nil {
		log.Errorf("ROLE Unmarshal failed %v,%v", accountName, err)
		return nil, err
	}

	return res, nil
}

// SetMsignAccountRole is common func to set role for multisign account
func SetMsignAccountRole(ldb *db.DBService, accountName string, value *MsignAccount) error {
	key := accountNameToKey(accountName)
	jsonvalue, err := json.Marshal(value)
	if err != nil {
		log.Errorf("DB Marshal failed %v,%v", accountName, err)
		return err
	}
	return ldb.SetObject(MsignAccountObjectName, key, string(jsonvalue))
}

// GetMsignAccountRole is common func to get role for multisign account
func GetMsignAccountRole(ldb *db.DBService, accountName string) (*MsignAccount, error) {
	key := accountNameToKey(accountName)
	value, err := ldb.GetObject(MsignAccountObjectName, key)
	if err != nil {
		return nil, err
	}

	res := &MsignAccount{}
	err = json.Unmarshal([]byte(value), res)
	if err != nil {
		log.Errorf("ROLE Unmarshal failed %v,%v", accountName, err)
		return nil, err
	}

	return res, nil
}
