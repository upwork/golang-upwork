package api

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestClientConstants(t *testing.T) {
    assert.Equal(t, "https://www.upwork.com/", BaseHost)
    assert.Equal(t, "api", DefaultEpoint)
    assert.Equal(t, "https://www.upwork.com/api/auth/v1/oauth/token/request", RequestTokenEP)
    assert.Equal(t, "https://www.upwork.com/api/auth/v1/oauth/token/access", AccessTokenEP)
    assert.Equal(t, "https://www.upwork.com/services/api/auth", AuthorizationEP)
    assert.Equal(t, "json", DataFormat)
    assert.Equal(t, "http_method", OverloadParam)
}

func TestSetup(t *testing.T) {
    client := Setup(ReadConfig("../example/config.json"))
    if assert.NotNil(t, client) {
        assert.NotNil(t, client.client)
        assert.NotNil(t, client.atoken)
    }
}

func TestSetEntryPoint(t *testing.T) {
    client := Setup(ReadConfig("../example/config.json"))
    client.SetEntryPoint("gds")

    assert.Equal(t, "gds", client.ep)
}