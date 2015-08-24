package api

import (
    "net/http"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestImports(t *testing.T) {
    if assert.Equal(t, 1, 1) != true {
        t.Error("Something is wrong.")
    }
}

func TestNewConfig(t *testing.T) {
    settings := map[string]string{
        "consumer_key": "consumerkey",
        "consumer_secret": "consumersecret",
        "access_token": "accesstoken",
        "access_secret": "accesssecret",
        "debug": "on",
    }
    config := NewConfig(settings)

    if assert.NotNil(t, config) {
        assert.Equal(t, "consumerkey", config.ConsumerKey)
        assert.Equal(t, "consumersecret", config.ConsumerSecret)
        assert.Equal(t, "accesstoken", config.AccessToken)
        assert.Equal(t, "accesssecret", config.AccessSecret)
        assert.True(t, true, config.Debug)
    }
}

func TestReadConfig(t *testing.T) {
    config := ReadConfig("../example/config.json")

    if assert.NotNil(t, config) {
        assert.Equal(t, "YOUR_CONSUMER_KEY", config.ConsumerKey)
        assert.Equal(t, "YOUR_CONSUMER_SECRET", config.ConsumerSecret)
        assert.Equal(t, "", config.AccessToken)
        assert.Equal(t, "", config.AccessSecret)
    }
}

func TestSetCustomHttpClient(t *testing.T) {
    c := &http.Client{}
    config := ReadConfig("../example/config.json")
    config.SetCustomHttpClient(c)
    
    if assert.NotNil(t, config) {
        assert.NotNil(t, config.CustomHttpClient)
    }
}