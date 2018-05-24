// Router for Upwork API
//
// Licensed under the Upwork's API Terms of Use;
// you may not use this file except in compliance with the Terms.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author::    Maksym Novozhylov (mnovozhilov@upwork.com)
// Copyright:: Copyright 2015(c) Upwork.com
// License::   See LICENSE.txt and TOS - https://developers.upwork.com/api-tos.html
package snapshot

import (
    "net/http"
    "github.com/upwork/golang-upwork/api"
)

const (
    EntryPoint = "api"
)

type a struct {
    client api.ApiClient
}

// Constructor
func New(c api.ApiClient) (a) {
    var r a
    c.SetEntryPoint(EntryPoint)
    r.client = c

    return r
}

// Get snapshot info by specific contract
func (r a) GetByContract(contractId string, ts string) (*http.Response, []byte) {
    return r.client.Get("/team/v3/snapshots/contracts/" + contractId + "/" + ts, nil)
}

// Update snapshot by specific contract
func (r a) UpdateByContract(contractId string, ts string, params map[string]string) (*http.Response, []byte) {
    return r.client.Post("/team/v3/snapshots/contracts/" + contractId + "/" + ts, params)
}

// Delete snapshot by specific contract
func (r a) DeleteByContract(contractId string, ts string) (*http.Response, []byte) {
    return r.client.Delete("/team/v3/snapshots/contracts/" + contractId + "/" + ts, nil)
}
