// Copyright (c) 2019 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"context"

	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/model"
)

type testModelWithHookOnApplyClientConnChanges struct {
	model.Model
	signalChannel     chan struct{}
	waitChannel       chan struct{}
	neededChangeState model.ClientConnectionState
	applyChangesFunc  func(context.Context, string, func(*model.ClientConnection)) *model.ClientConnection
}

func (m *testModelWithHookOnApplyClientConnChanges) ApplyClientConnectionChanges(ctx context.Context, connectionID string, changeFunc func(*model.ClientConnection)) *model.ClientConnection {
	rv := m.Model.ApplyClientConnectionChanges(ctx, connectionID, changeFunc)

	// Check whether changeFunc changes connection state to neededChangeState. If so, signal to the channel.
	dummyConnection := &model.ClientConnection{
		ConnectionState: model.ClientConnectionReady,
	}

	changeFunc(dummyConnection)

	if dummyConnection.ConnectionState == m.neededChangeState {
		close(m.signalChannel)
		<-m.waitChannel
	}

	return rv
}
