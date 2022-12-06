package rocketfuel

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

const (
	version            = "0.0.2"
	defaultHTTPTimeout = 60 * time.Second
	baseURL            = "https://api.rocketfuelblockchain.com/api"
	userAgent          = "rocketfuel-go" + version
)

type service struct {
	client *Client
}
type Client struct {
	common         service
	client         *http.Client
	options        *Options
	baseURL        *url.URL
	logger         Logger
	Key            string
	Log            Logger
	LoggingEnabled bool
	Authorization  *AuthorizationService
	Service        *Service
	HostedPage     *HostedPageService
	Update         *UpdateService
}
type Options struct {
	Environment  string `json:"environment,omitempty"`
	PublicKey    string `json:"publicKey,omitempty"`
	MerchantId   string `json:"merchantId,omitempty"`
	ClientId     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
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
func GetBaseUrl(env string) string {
	prodUrl := "https://app.rocketfuelblockchain.com/api"

	if env == "" {
		return prodUrl
	}

	environment_data := map[string]string{
		"prod":    prodUrl,
		"dev":     "https://dev-app.rocketdemo.net/api",
		"qa":      "https://qa-app.rocketdemo.net/api",
		"preprod": "https://preprod-app.rocketdemo.net/api",
		"sandbox": "https://app-sandbox.rocketfuelblockchain.com/api",
	}

	return environment_data[env]
}

// NewClient creates a new Rocketfuel API client with the given API key
// and HTTP client, allowing overriding of the HTTP client to use.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
func NewClient(options *Options, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultHTTPTimeout}
	}

	u, _ := url.Parse(GetBaseUrl(options.Environment))
	c := &Client{
		client:         httpClient,
		options:        options,
		baseURL:        u,
		LoggingEnabled: true,
		Log:            log.New(os.Stderr, "", log.LstdFlags),
	}

	c.common.client = c
	c.Authorization = (*AuthorizationService)(&c.common)
	c.HostedPage = (*HostedPageService)(&c.common)
	c.Update = (*UpdateService)(&c.common)

	return c
}

func (c *Client) Call(method, path string, body string, v interface{}) error {
	var buf *strings.Reader
	c.Log.Printf("body: %v\n", body)

	if body != "" {

		buf = strings.NewReader(body)

	}

	u, _ := c.baseURL.Parse("api" + path)

	if c.LoggingEnabled {
		c.Log.Printf("buf: %v\n ", buf)
	}

	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		if c.LoggingEnabled {
			c.Log.Printf("Cannot create Rocketfuel request: %v\n", err)
		}
		return err
	}

	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.Key != "" {
		req.Header.Set("Authorization", "Bearer "+c.Key)
	}

	req.Header.Set("User-Agent", userAgent)

	if c.LoggingEnabled {
		c.Log.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
		c.Log.Printf("POST request data %v\n", buf)
	}

	start := time.Now()

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if c.LoggingEnabled {
		c.Log.Printf("Completed in %v\n", time.Since(start))
	}

	defer resp.Body.Close()
	return c.decodeResponse(resp, v)
}

func GetOptions(environment string, publicKey string, clientId string, merchantId string, clientSecret string) *Options {

	options := &Options{
		Environment:  environment,
		PublicKey:    publicKey,
		ClientId:     clientId,
		MerchantId:   merchantId,
		ClientSecret: clientSecret,
	}

	return options
}
func (c *Client) getMerchantCred() string {

	toEncrypt := marshalize(map[string]string{
		"merchantId": c.options.MerchantId, "totp": ""})

	mapB := marshalize(map[string]string{"clientId": c.options.ClientId, "encryptedPayload": encrypt(string(toEncrypt),
		c.options.ClientSecret)})

	return string(mapB)
}

// decodeResponse decodes the JSON response
// The actual response will be written to the `v` parameter
func (c *Client) decodeResponse(httpResp *http.Response, v interface{}) error {
	var resp Response
	respBody, err := ioutil.ReadAll(httpResp.Body)
	json.Unmarshal(respBody, &resp)

	if resp["ok"] != true || httpResp.StatusCode != 200 {
		if c.LoggingEnabled {
			c.Log.Printf("Rocketfuel error: %+v", err)
			c.Log.Printf("HTTP Response: %+v", resp)
		}
		return newAPIError(httpResp)
	}

	if c.LoggingEnabled {
		c.Log.Printf("Rocketfuel response: %v\n", resp)
	}
	if data, ok := resp["result"]; ok {

		switch t := resp["result"].(type) {
		case map[string]interface{}:
			return mapstruct(data, v)
		default:
			_ = t
			return mapstruct(resp, v)
		}
	}
	// if response data does not contain data key, map entire response to v
	return mapstruct(resp, v)
}

func mapstruct(data interface{}, v interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)

	if err != nil {
		return err
	}

	err = decoder.Decode(data)

	return err
}
func (c *Client) GetUUID(body HostedPageRequest) (Response, error) {
	result, _ := c.Authorization.Login()
	body.Merchant_id = c.options.MerchantId //update body with merchant Id
	if str, ok := result["access"].(string); ok {
		c.Key = str
	} else {
		fmt.Println("not a string", result)
		panic("Authorization could not be completed")
	}
	fmt.Println(body, "body")
	return c.HostedPage.Create(body)
}
func newAPIError(httpResp *http.Response) error {
	return errors.New("There was an error")
}
