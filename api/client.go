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
    "strings"
    "fmt"
    "bytes"
    "net/http"
    "net/url"
    "io/ioutil"

    "github.com/mnovozhylov/oauth" // this is a forked dependency, to avoid unexpected behavior because of any changes
)

// Define end points
const (
    BaseHost = "https://www.upwork.com/"
    DefaultEpoint = "api"
    RequestTokenEP = BaseHost + DefaultEpoint + "/auth/v1/oauth/token/request"
    AuthorizationEP = BaseHost + "services/api/auth"
    AccessTokenEP = BaseHost + DefaultEpoint + "/auth/v1/oauth/token/access"
    DataFormat = "json"
    OverloadParam = "http_method"
)

// Api client
type ApiClient struct {
    client *oauth.Consumer
    atoken *oauth.AccessToken
    rtoken *oauth.RequestToken
    oclient *http.Client
    ep string
}

// Setup client using specific config
func Setup(config *Config) (client ApiClient) {
    var c ApiClient

    c.client = setupNewConsumer(config, config.CustomHttpClient)

    if config.Debug {
        c.client.Debug(true)
    }

    c.atoken = &oauth.AccessToken{
        Token: config.AccessToken,
        Secret: config.AccessSecret,
    }

    return c
}

// Set entry point, e.g requested from a router
func (c *ApiClient) SetEntryPoint(ep string) {
    c.ep = ep
}

// Receive a request token/secret pair
// and get authorization url with a verifier
func (c *ApiClient) GetAuthorizationUrl(callback string) (authzUrl string) {
    if callback == "" {
        callback = "oob"
    }

    requestToken, u, err := c.client.GetRequestTokenAndUrl(callback)
    if err != nil {
        log.Fatal(err)
    }

    c.rtoken = requestToken

    return u
}

// Get access token using a specific verifier
func (c *ApiClient) GetAccessToken(verifier string) (creds *oauth.AccessToken) {
    accessToken, err := c.client.AuthorizeToken(c.rtoken, strings.Trim(verifier, "\n"))
    if err != nil {
        log.Fatal(err)
    }

    c.atoken = accessToken
    c.setupOauthClient()

    return c.atoken
}

// Check if client contains already a token/secret pair
func (c *ApiClient) HasAccessToken() (bool) {
    has := (c.atoken.Token != "" && c.atoken.Secret != "")
    if has {
        c.setupOauthClient()
    }
    return has
}

// GET method for client
func (c *ApiClient) Get(uri string, params map[string]string) (resp *http.Response, rb []byte) {
    // parameters must be encoded according to RFC 3986
    // hmmm, it seems to be the easiest trick?
    qstr := ""
    if params != nil {
        for k, v := range params {
            qstr += fmt.Sprintf("%s=%s&", k, v)
        }
        qstr = qstr[0:len(qstr)-1]
    }
    // use Path on request from the user
    // ./ will be later replaced
    u := &url.URL{Path: qstr}

    // https://github.com/mrjones/oauth/issues/34
    encQuery := strings.Replace(u.String(), ";", "%3B", -1)
    encQuery = strings.Replace(encQuery, "./", "?", 1) // see URL.String method to understand when "./" is returned

    // non-empty string may miss "?"
    if encQuery !="" && encQuery[:1] != "?" {
        encQuery = "?" + encQuery
    }

    response, err := c.oclient.Get(formatUri(uri, c.ep) + encQuery)

    return formatResponse(response, err)
}

// POST method for client
func (c *ApiClient) Post(uri string, params map[string]string) (r *http.Response, rb []byte) {
    return c.sendPostRequest(uri, params)
}

// PUT method for client
func (c *ApiClient) Put(uri string, params map[string]string) (r *http.Response, rb []byte) {
    return c.sendPostRequest(uri, addOverloadParam(params, "put"))
}

// DELETE method for client
func (c *ApiClient) Delete(uri string, params map[string]string) (r *http.Response, rb []byte) {
    return c.sendPostRequest(uri, addOverloadParam(params, "delete"))
}

// setup/save authorized oauth client, based on received or provided access token credentials
func (c *ApiClient) setupOauthClient() {
    // setup authorized oauth client
    client, err := c.client.MakeHttpClient(c.atoken)
    if err != nil {
        log.Fatal(err)
    }

    c.oclient = client
}

// run post/put/delete requests
func (c *ApiClient) sendPostRequest(uri string, params map[string]string) (r *http.Response, rb []byte) {
    var jsonStr = []byte("{}")
    if params != nil {
        str := ""
        for k, v := range params {
            str += fmt.Sprintf("\"%s\": \"%s\",", k, v)
        }
        jsonStr = []byte(fmt.Sprintf("{%s}", str[0:len(str)-1]))
    }

    response, err := c.oclient.Post(formatUri(uri, c.ep), "application/json", bytes.NewBuffer(jsonStr))

    return formatResponse(response, err)
}

// Create new OAuth consumer, based on setup config and possibly a custom http client
func setupNewConsumer(config *Config, httpClient *http.Client) *oauth.Consumer {
    if (httpClient == nil) {
        return oauth.NewConsumer(
                config.ConsumerKey,
                config.ConsumerSecret,
                oauth.ServiceProvider{
                    RequestTokenUrl:   RequestTokenEP,
                    AuthorizeTokenUrl: AuthorizationEP,
                    AccessTokenUrl:    AccessTokenEP,
                    HttpMethod:    "POST",
                })
    } else {
        return oauth.NewCustomHttpClientConsumer(
                config.ConsumerKey,
                config.ConsumerSecret,
                oauth.ServiceProvider{
                    RequestTokenUrl:   RequestTokenEP,
                    AuthorizeTokenUrl: AuthorizationEP,
                    AccessTokenUrl:    AccessTokenEP,
                    HttpMethod:    "POST",
                }, httpClient)
    }
}

// Check and format (preparate a byte body) http response routine
func formatResponse(resp *http.Response, err error) (r *http.Response, rb []byte) {
    if err != nil {
        log.Fatal("Can not execute the request, " + err.Error())
    }

    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        // do not exit, it can be a normal response
        // it's up to client/requester's side decide what to do
    }
    // read json http response
    jsonDataFromHttp, _ := ioutil.ReadAll(resp.Body)

    return resp, jsonDataFromHttp
}

// Create a path to a specific resource
func formatUri(uri string, ep string) (string) {
    format := ""
    if ep == DefaultEpoint {
        format += "." + DataFormat
    }
    return BaseHost + ep + uri + format
}

// add overload parameter to the map of parameters
func addOverloadParam(params map[string]string, op string) map[string]string {
    if params == nil {
        params = make(map[string]string)
    }
    params[OverloadParam] = op
    return params
}
