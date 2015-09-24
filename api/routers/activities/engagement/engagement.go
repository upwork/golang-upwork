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
package engagement

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

// List activities for specific engagement
func (r a) GetSpecific(engagementRef string) (*http.Response, []byte) {
    return r.client.Get("/tasks/v2/tasks/contracts/" + engagementRef, nil)
}

// Assign engagements to the list of activities
func (r a) Assign(company string, team string, engagement string, params map[string]string) (*http.Response, []byte) {
    return r.client.Put("/otask/v1/tasks/companies/" + company + "/teams/" + team + "/engagements/" + engagement + "/tasks", params)
}

// Assign to specific engagement the list of activities
func (r a) AssignToEngagement(engagementRef string, params map[string]string) (*http.Response, []byte) {
    return r.client.Put("/tasks/v2/tasks/contracts/" + engagementRef, params)
}
