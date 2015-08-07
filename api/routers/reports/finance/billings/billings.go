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
package billings

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

// Generate Billing Reports for a Specific Freelancer
func (r a) GetByFreelancer(freelancerReference string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/finreports/v2/providers/" + freelancerReference + "/billings", params)
}

// Generate Billing Reports for a Specific Freelancer's Team
func (r a) GetByFreelancersTeam(freelancerTeamReference string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/finreports/v2/provider_teams/" + freelancerTeamReference + "/billings", params)
}

// Generate Billing Reports for a Specific Freelancer's Company
func (r a) GetByFreelancersCompany(freelancerCompanyReference string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/finreports/v2/provider_companies/" + freelancerCompanyReference + "/billings", params)
}

// Generate Billing Reports for a Specific Buyer's Team
func (r a) GetByBuyersTeam(buyerTeamReference string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/finreports/v2/buyer_teams/" + buyerTeamReference + "/billings", params)
}

// Generate Billing Reports for a Specific Buyer's Company
func (r a) GetByBuyersCompany(buyerCompanyReference string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/finreports/v2/buyer_companies/" + buyerCompanyReference + "/billings", params)
}