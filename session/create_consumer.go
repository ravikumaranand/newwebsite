/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
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

package session

import (
	"encoding/json"

	"github.com/mysteriumnetwork/node/communication"
	"github.com/mysteriumnetwork/node/identity"
)

// createConsumer processes session create requests from communication channel.
type createConsumer struct {
	sessionCreator Creator
	peerID         identity.Identity
	configProvider ConfigProvider
}

// Creator defines method for session creation
type Creator interface {
	Create(consumerID identity.Identity, proposalID int, config ServiceConfiguration, destroyCallback DestroyCallback) (Session, error)
}

// GetMessageEndpoint returns endpoint there to receive messages
func (consumer *createConsumer) GetRequestEndpoint() communication.RequestEndpoint {
	return endpointSessionCreate
}

// NewRequest creates struct where request from endpoint will be serialized
func (consumer *createConsumer) NewRequest() (requestPtr interface{}) {
	var request CreateRequest
	return &request
}

// Consume handles requests from endpoint and replies with response
func (consumer *createConsumer) Consume(requestPtr interface{}) (response interface{}, err error) {
	request := requestPtr.(*CreateRequest)

	config, destroyCallback, err := consumer.configProvider(request.Config)
	if err != nil {
		return responseInternalError, err
	}

	// TODO: consume the consumerInfo here from the request
	// Session is the obvious choice for this
	sessionInstance, err := consumer.sessionCreator.Create(consumer.peerID, request.ProposalId, config, destroyCallback)
	switch err {
	case nil:
		// TODO: the last promise info should also be loaded here somehow
		lastPromise := LastPromise{
			SequenceID: 1,
			Amount:     0,
		}
		return responseWithSession(sessionInstance, lastPromise), nil
	case ErrorInvalidProposal:
		return responseInvalidProposal, nil
	default:
		return responseInternalError, nil
	}
}

func responseWithSession(sessionInstance Session, lp LastPromise) CreateResponse {
	serializedConfig, err := json.Marshal(sessionInstance.Config)
	if err != nil {
		// Failed to serialize session
		// TODO Cant expose error to response, some logging should be here
		return responseInternalError
	}

	return CreateResponse{
		Success: true,
		Session: SessionDto{
			ID:     sessionInstance.ID,
			Config: serializedConfig,
		},
		LastPromise: &lp,
	}
}
