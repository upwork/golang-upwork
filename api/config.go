// Package implements access to Upwork API
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
package api

import (
    "log"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

// Config
type Config struct {
    ConsumerKey string
    ConsumerSecret string
    AccessToken string
    AccessSecret string
    Debug bool
    CustomHttpClient *http.Client
}

// List of required configuration keys
var requiredKeys = [2]string{"consumer_key", "consumer_secret"}

// Create a new config
func NewConfig(data map[string]string) (settings *Config) {
    cfg := &Config{
        ConsumerKey: data["consumer_key"],
        ConsumerSecret: data["consumer_secret"],
    }

    // save access token if defined
    if atoken, ok := data["access_token"]; ok {
        cfg.AccessToken = atoken
    }

    // save access token secret if defined
    if asecret, ok := data["access_secret"]; ok {
        cfg.AccessSecret = asecret
    }

    // save debug flag if defined
    if debug, ok := data["debug"]; ok && debug == "on" {
        cfg.Debug = true
    }

    return cfg
}

// Read a specific configuration (json) file
func ReadConfig(fn string) (settings *Config) {
    // read from config file if exists
    b, err := ioutil.ReadFile(fn)
    if err != nil {
        log.Fatal("config file: ", err)
    }

    // parse json config
    var data map[string]interface{}
    if err := json.Unmarshal(b, &data); err != nil {
        log.Fatal("config file: ", err)
    }
    
    // test required properties
    for _, v := range requiredKeys {
        _, ok := data[v]
        if !ok {
            log.Fatal("config file: " + v + " is missing in " + fn)
        }
    }

    // convert
    config := make(map[string]string)
    for k, v := range data {
        config[k] = v.(string)
    }

    return NewConfig(config)
}

// Configure for usage with custom http client
func (cfg *Config) SetCustomHttpClient(c *http.Client) *Config {
    cfg.CustomHttpClient = c
    return cfg
}

// Test print of found/assigned key
func (cfg *Config) Print() {
    fmt.Println("assigned key:", cfg.ConsumerKey)
}