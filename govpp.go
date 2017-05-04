// Copyright (c) 2017 Cisco and/or its affiliates.
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

package govpp

import (
	"gerrit.fd.io/r/govpp.git/adapter"
	"gerrit.fd.io/r/govpp.git/adapter/vppapiclient"
	"gerrit.fd.io/r/govpp.git/core"
)

var vppAdapter adapter.VppAdapter // VPP Adapter that will be used in the subsequent Connect calls

// Connect connects the govpp core to VPP either using the default VPP Adapter, or using the adapter previously
// set by SetAdapter (useful mostly just for unit/integration tests with mocked VPP adapter).
func Connect() (*core.Connection, error) {
	if vppAdapter == nil {
		vppAdapter = vppapiclient.NewVppAdapter()
	}
	return core.Connect(vppAdapter)
}

// SetAdapter sets the adapter that will be used for connections to VPP in the subsequent `Connect` calls.
func SetAdapter(ad adapter.VppAdapter) {
	vppAdapter = ad
}