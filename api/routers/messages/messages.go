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
// Copyright:: Copyright 2016(c) Upwork.com
// License::   See LICENSE.txt and TOS - https://developers.upwork.com/api-tos.html
package messages

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

// Retrieve rooms information
func (r a) GetRooms(company string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/messages/v3/" + company + "/rooms", params)
}

// Get a specific room information
func (r a) GetRoomDetails(company string, roomId string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/messages/v3/" + company + "/rooms/" + roomId, params)
}

// Get messages from a specific room
func (r a) GetRoomMessages(company string, roomId string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/messages/v3/" + company + "/rooms/" + roomId + "/stories", params)
}

// Get a specific room by offer ID
func (r a) GetRoomByOffer(company string, offerId string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/messages/v3/" + company + "/rooms/offers/" + offerId, params)
}

// Get a specific room by application ID
func (r a) GetRoomByApplication(company string, applicationId string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/messages/v3/" + company + "/rooms/applications/" + applicationId, params)
}

// Get a specific room by contract ID
func (r a) GetRoomByContract(company string, contractId string, params map[string]string) (*http.Response, []byte) {
    return r.client.Get("/messages/v3/" + company + "/rooms/contracts/" + contractId, params)
}

// Create a new room
func (r a) CreateRoom(company string, params map[string]string) (*http.Response, []byte) {
    return r.client.Post("/messages/v3/" + company + "/rooms", params)
}

// Send a message to a room
func (r a) SendMessageToRoom(company string, roomId string, params map[string]string) (*http.Response, []byte) {
    return r.client.Post("/messages/v3/" + company + "/rooms/" + roomId + "/stories", params)
}

// Send a message to a batch of rooms
func (r a) SendMessageToRooms(company string, params map[string]string) (*http.Response, []byte) {
    return r.client.Post("/messages/v3/" + company + "/stories/batch", params)
}

// Update a room settings
func (r a) UpdateRoomSettings(company string, roomId string, username string, params map[string]string) (*http.Response, []byte) {
    return r.client.Put("/messages/v3/" + company + "/rooms/" + roomId + "/users/" + username, params)
}

// Update the metadata of a room
func (r a) UpdateRoomMetadata(company string, roomId string, params map[string]string) (*http.Response, []byte) {
    return r.client.Put("/messages/v3/" + company + "/rooms/" + roomId, params)
}
