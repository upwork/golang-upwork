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
package offers

import (
	"github.com/upwork/golang-upwork/api"
	"net/http"
)

const (
	EntryPoint = "api"
)

type a struct {
	client api.ApiClient
}

// Constructor
func New(c api.ApiClient) a {
	var r a
	c.SetEntryPoint(EntryPoint)
	r.client = c

	return r
}

// Get list of offers
func (r a) GetList(params map[string]string) (*http.Response, []byte) {
	return r.client.Get("/offers/v1/contractors/offers", params)
}

// Get specific offer
func (r a) GetSpecific(reference string) (*http.Response, []byte) {
	return r.client.Get("/offers/v1/contractors/offers/"+reference, nil)
}

// Run a specific action
func (r a) MakeOffer(reference string, params map[string]string) (*http.Response, []byte) {
	return r.client.Post("/offers/v1/contractors/offers/"+reference, params)
}
