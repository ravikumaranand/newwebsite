/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package connection

import (
	"time"

	"github.com/mysteriumnetwork/node/client/stats"
	"github.com/mysteriumnetwork/node/core/storage"
	"github.com/mysteriumnetwork/node/session"
)

// SessionStorage describes functions for storing session objects
type SessionStorage interface {
	Save(Session) error
	Update(session.ID, time.Time, stats.SessionStats) error
	GetAll() ([]Session, error)
}

type sessionStorage struct {
	storage storage.Storage
}

// NewSessionStorage creates session repository with given dependencies
func NewSessionStorage(storage storage.Storage) SessionStorage {
	return &sessionStorage{
		storage: storage,
	}
}

// Save saves a new session
func (repo *sessionStorage) Save(se Session) error {
	return repo.storage.GetDB().Save(&se)
}

// Update updates specified fields of existing session by id
func (repo *sessionStorage) Update(sessionID session.ID, TimeUpdated time.Time, dataStats stats.SessionStats) error {
	// update two fields by sessionID
	se := Session{SessionID: sessionID, TimeUpdated: TimeUpdated, DataStats: dataStats}
	return repo.storage.GetDB().Update(&se)
}

// GetAll returns array of all sessions
func (repo *sessionStorage) GetAll() ([]Session, error) {
	var sessions []Session
	err := repo.storage.GetDB().All(&sessions)
	if err != nil {
		return nil, err
	}
	return sessions, nil
}
