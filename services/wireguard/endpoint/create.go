// +build !linux linux,android

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

package endpoint

import (
	wg "github.com/mysteriumnetwork/node/services/wireguard"
	"github.com/mysteriumnetwork/node/services/wireguard/endpoint/userspace"
	"github.com/mysteriumnetwork/node/services/wireguard/resources"
)

// NewConnectionEndpoint creates new wireguard connection endpoint.
func NewConnectionEndpoint(publicIP string, resourceAllocator *resources.Allocator) (wg.ConnectionEndpoint, error) {
	client, err := userspace.NewWireguardClient()
	return &connectionEndpoint{
		wgClient:          client,
		publicIP:          publicIP,
		resourceAllocator: resourceAllocator,
	}, err
}
