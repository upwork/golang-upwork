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
package team

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

// List all oTask/Activity records within a team
func (r a) GetList(company string, team string, params ...map[string]string) (*http.Response, []byte) {
    var p map[string]string
    if params != nil {
	p = params[0]
    }
    return r.getByType(company, team, "", p)
}

// List all oTask/Activity records within a team by specified code(s)
func (r a) GetSpecificList(company string, team string, code string, params ...map[string]string) (*http.Response, []byte) {
    var p map[string]string
    if params != nil {
	p = params[0]
    }
    return r.getByType(company, team, code, p)
}

// Create an oTask/Activity record within a team
func (r a) AddActivity(company string, team string, params map[string]string) (*http.Response, []byte) {
    return r.client.Post("/otask/v1/tasks/companies/" + company + "/teams/" + team + "/tasks", params)
}

// Update specific oTask/Activity record within a team
func (r a) UpdateActivity(company string, team string, code string, params map[string]string) (*http.Response, []byte) {
    return r.client.Put("/otask/v1/tasks/companies/" + company + "/teams/" + team + "/tasks/" + code, params)
}

// Archive specific oTask/Activity record within a team
func (r a) ArchiveActivity(company string, team string, code string) (*http.Response, []byte) {
    return r.client.Put("/otask/v1/tasks/companies/" + company + "/teams/" + team + "/archive/" + code, nil)
}

// Unarchive specific oTask/Activity record within a team
func (r a) UnarchiveActivity(company string, team string, code string) (*http.Response, []byte) {
    return r.client.Put("/otask/v1/tasks/companies/" + company + "/teams/" + team + "/unarchive/" + code, nil)
}

// Update a group of oTask/Activity records
func (r a) UpdateBatch(company string, params map[string]string) (*http.Response, []byte) {
    return r.client.Put("/otask/v1/tasks/companies/" + company + "/tasks/batch", params)
}

// Get by type 
func (r a) getByType(company string, team string, code string, params map[string]string) (*http.Response, []byte) {
    url := ""
    if code != "" {
        url = "/" + code;
    }

    return r.client.Get("/otask/v1/tasks/companies/" + company + "/teams/" + team + "/tasks" + url, params)
}
