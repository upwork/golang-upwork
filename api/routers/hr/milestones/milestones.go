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
package milestones

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

// Get active Milestone for the Contract
func (r a) GetActiveMilestone(contractId string) (*http.Response, []byte) {
    return r.client.Get("/hr/v3/fp/milestones/statuses/active/contracts/" + contractId, nil)
}

// Get all submissions for the active Milestone
func (r a) GetSubmissions(milestoneId string) (*http.Response, []byte) {
    return r.client.Get("/hr/v3/fp/milestones/" + milestoneId + "/submissions", nil)
}

// Create a new Milestone
func (r a) Create(params map[string]string) (*http.Response, []byte) {
    return r.client.Post("/hr/v3/fp/milestones", params)
}

// Edit an existing Milestone
func (r a) Edit(milestoneId string, params map[string]string) (*http.Response, []byte) {
    return r.client.Put("/hr/v3/fp/milestones/" + milestoneId, params)
}

// Activate an existing Milestone
func (r a) Activate(milestoneId string, params map[string]string) (*http.Response, []byte) {
    return r.client.Put("/hr/v3/fp/milestones/" + milestoneId + "/activate", params)
}

// Approve an existing Milestone
func (r a) Approve(milestoneId string, params map[string]string) (*http.Response, []byte) {
    return r.client.Put("/hr/v3/fp/milestones/" + milestoneId + "/approve", params)
}

// Delete an existing Milestone
func (r a) Delete(milestoneId string) (*http.Response, []byte) {
    return r.client.Delete("/hr/v3/fp/milestones/" + milestoneId , nil)
}