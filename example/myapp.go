// Example shows how to work with Upwork API
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
package main

import (
    "fmt"
    "bufio"
    "os"
    //"net/http" //uncomment if you need to setup a custom http client
    
    "github.com/upwork/golang-upwork/api"
    "github.com/upwork/golang-upwork/api/routers/auth"
    //"github.com/upwork/golang-upwork/api/routers/mc" // uncomment to test mc examples
)

const cfgFile = "config.json" // update the path to your config file, or provide properties directly in your code

func main() {
/* it is possible to set up properties from code
    settings := map[string]string{
        "consumer_key": "consumerkey",
        "consumer_secret": "consumersecret",
    }
    config := api.NewConfig(settings)

    //or read them from a specific configuration file
    config := api.ReadConfig(cfgFile)
    config.Print()
*/
/* it is possible to setup a custom http client if needed
    c := &http.Client{}
    config := api.ReadConfig(cfgFile)
    config.SetCustomHttpClient(c)
    client := api.Setup(config)
*/
    
    client := api.Setup(api.ReadConfig(cfgFile))
    // we need an access token/secret pair in case we haven't received it yet
    if !client.HasAccessToken() {
        aurl := client.GetAuthorizationUrl("")

        // read verifier
        reader := bufio.NewReader(os.Stdin)
        fmt.Println("Visit the authorization url and provide oauth_verifier for further authorization")
        fmt.Println(aurl)
        verifier, _ := reader.ReadString('\n')

        // get access token
        token := client.GetAccessToken(verifier)
        fmt.Println(token)
    }

    // http.Response and []byte will be return, you can use any
    _, jsonDataFromHttp1 := auth.New(client).GetUserInfo()
    
    // here you can Unmarshal received json string, or do any other action(s)
    fmt.Println(string(jsonDataFromHttp1))

    // getting reports example
    //params := make(map[string]string)
    //params["tq"] = "select memo where worked_on >= '05-08-2015'"
    //params["tqx"] = "out:json"
    //_, jsonDataFromHttp4 := timereports.New(client).GetByFreelancerFull(params)
    //fmt.Println(string(jsonDataFromHttp4))
}
