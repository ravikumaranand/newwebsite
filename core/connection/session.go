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
	"github.com/mysteriumnetwork/node/identity"
	"github.com/mysteriumnetwork/node/session"
)

// Session holds structure for saving session history
type Session struct {
	SessionID       session.ID `storm:"id"`
	ProviderID      identity.Identity
	ServiceType     string
	ProviderCountry string
	TimeStarted     time.Time
	// these will be updated on disconnect event
	TimeUpdated time.Time // Duration can be calculated: TimeUpdated - TimeStarted
	DataStats   stats.SessionStats
}
