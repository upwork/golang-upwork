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
package accounts

import (
    "net/http"
    "github.com/upwork/golang-upwork/api"
)

const (
    EntryPoint = "gds"
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

// Generate Financial Reports for an owned Account
func (r a) GetOwned(freelancerReference string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/finreports/v2/financial_account_owner/" + freelancerReference, params)
}

// Generate Financial Reports for a Specific Account
func (r a) GetSpecific(entityReference string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/finreports/v2/financial_accounts/" + entityReference, params)
}