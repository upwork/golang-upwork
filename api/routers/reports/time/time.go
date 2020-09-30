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
package time

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

// Generate Time Reports for a Specific Team (with financial info)
func (r a) GetByTeamFull(company string, team string, params map[string]string) (*http.Response, []byte) {
    return r.getByType(company, team, "", params, false)
}

// Generate Time Reports for a Specific Team (hide financial info)
func (r a) GetByTeamFullLimited(company string, team string, params map[string]string) (*http.Response, []byte) {
    return r.getByType(company, team, "", params, true)
}

// Generating Agency Specific Reports
func (r a) GetByAgency(company string, agency string, params map[string]string) (*http.Response, []byte) {
    return r.getByType(company, "", agency, params, false)
}

// Generating Company Wide Reports
func (r a) GetByCompany(company string, params map[string]string) (*http.Response, []byte) {
    return r.getByType(company, "", "", params, false)
}

// Generating Freelancer's Specific Reports (with financial info)
func (r a) GetByFreelancerFull(freelancerId string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/timereports/v1/providers/" + freelancerId, params)
}

// Generating Freelancer's Specific Reports (hide financial info)
func (r a) GetByFreelancerLimited(freelancerId string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/timereports/v1/providers/" + freelancerId + "/hours", params)
}

// Get by type 
func (r a) getByType(company string, team string, agency string, params map[string]string, hideFinDetails bool) (*http.Response, []byte) {
    url := ""
    if team != "" {
        url = "/teams/" + team
        if hideFinDetails {
            url = url + "/hours"
        }
    } else if agency != "" {
        url = "/agencies/" + agency
    }

    return r.client.Get("/timereports/v1/companies/" + company + url, params)
}
