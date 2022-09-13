package rocketfuel

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	version            = "0.0.1"
	defaultHTTPTimeout = 60 * time.Second
	baseURL            = "https://api.rocketfuelblockchain.com"
	userAgent          = "rocketfuel-go" + version
)

type service struct {
	client *Client
}
type Client struct {
	common    service
	client    *http.Client
	publicKey string
	baseURL   *url.URL
	logger    Logger
}

type Logger interface {
	Printf(format string, v ...interface{})
}
type Response map[string]interface{}

type RequestValues url.Values

func (v RequestValues) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 3)
	for k, val := range v {
		m[k] = val[0]
	}
	return json.Marshal(m)
}

func init() {

}
