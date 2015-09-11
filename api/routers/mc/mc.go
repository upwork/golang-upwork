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
package mc

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

// Get trays
func (r a) GetTrays() (*http.Response, []byte) {
	return r.client.Get("/mc/v1/trays", nil)
}

// Get tray by type
func (r a) GetTrayByType(username string, tType string) (*http.Response, []byte) {
	return r.client.Get("/mc/v1/trays/"+username+"/"+tType, nil)
}

// List thread details based on thread id
func (r a) GetThreadDetails(username string, threadId string, params map[string]string) (*http.Response, []byte) {
	return r.client.Get("/mc/v1/threads/"+username+"/"+threadId, params)
}

// Get a specific thread by "Interview" context
func (r a) GetThreadByContext(username string, jobKey string, applicationId string, context string, params map[string]string) (*http.Response, []byte) {
	if context == "" {
		context = "Interviews"
	}
	return r.client.Get("/mc/v1/contexts/"+username+"/"+context+":"+jobKey+":"+applicationId, params)
}

// Get a specific thread by "Interview" context
func (r a) GetThreadByContextLastPosts(username string, jobKey string, applicationId string, context string, params map[string]string) (*http.Response, []byte) {
	if context == "" {
		context = "Interviews"
	}
	return r.client.Get("/mc/v1/contexts/"+username+"/"+context+":"+jobKey+":"+applicationId+"/last_posts", params)
}

// Send new message
func (r a) StartNewThread(username string, params map[string]string) (*http.Response, []byte) {
	return r.client.Post("/mc/v1/threads/"+username, params)
}

// Reply to existent thread
func (r a) ReplyToThread(username string, threadId string, params map[string]string) (*http.Response, []byte) {
	return r.client.Post("/mc/v1/threads/"+username+"/"+threadId, params)
}

// Update threads based on user actions
func (r a) MarkThread(username string, threadId string, params map[string]string) (*http.Response, []byte) {
	return r.client.Put("/mc/v1/threads/"+username+"/"+threadId, params)
}
