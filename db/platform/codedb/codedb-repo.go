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
 * file description: code database interface
 * @Author: May Luo
 * @Date:   2017-12-05
 * @Last Modified by:
 * @Last Modified time:
 */

package codedb

import (
	"errors"
        "sync"

	log "github.com/cihub/seelog"
	"github.com/tidwall/buntdb"
)

//CodeDbRepository is to build code db
type CodeDbRepository struct {
	fn string     // filename for reporting
	db *buntdb.DB // LevelDB instance
	tx *buntdb.Tx
        globalSignal   sync.RWMutex // global signal for all state
}

//NewCodeDbRepository is to create new code db
func NewCodeDbRepository(file string) (*CodeDbRepository, error) {
	codedb, err := buntdb.Open(file)
	if err != nil {
		return nil, err
	}
	return &CodeDbRepository{
		fn: file,
		db: codedb,
	}, nil
}
//CallGlobalLock is to lock
func (m *CodeDbRepository) CallGlobalLock() {
	log.Info("CallGlobalLock")
	m.globalSignal.Lock()
}
//CallGlobalUnLock is to unlock
func (m *CodeDbRepository) CallGlobalUnLock() {
	log.Info("CallGlobalUnLock")
	m.globalSignal.Unlock()
}
//CallStartUndoSession is to start undo session
func (k *CodeDbRepository) CallStartUndoSession(writable bool) {
	k.tx, _ = k.db.Begin(true)
}

//CallCreatObjectIndex is to create object index
func (k *CodeDbRepository) CallCreatObjectIndex(objectName string, indexName string, indexJson string) error {
	if k.tx == nil {

		return k.db.CreateIndex(indexName, objectName+"*", buntdb.IndexJSON(indexJson))
	}

	return k.tx.CreateIndex(indexName, objectName+"*", buntdb.IndexJSON(indexJson))
}
func (m *MultindexDB) CallClose() {
	m.db.Close()
	m.undoList = nil
	m.session = nil
	m.subsession = nil
	m.sessionEx = nil
	m.revision = uint64(0)
	m.commitRevision = uint64(0)
	m.ai = nil

}
func (m *MultindexDB) CallGlobalLock() {
	m.globalSignal.Lock()
}

func (m *MultindexDB) CallGlobalUnLock() {
	m.globalSignal.Unlock()
}

func (m *MultindexDB) CallLock() {
	m.signal.Lock()
}

func (m *MultindexDB) CallUnLock() {
	m.signal.Unlock()
}
