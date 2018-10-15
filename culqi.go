// Package culqi implementa metodos para inicializar cliente
//
// Inicia cliente con datos de condifuración de comercio.
package culqi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	libraryVersion     = "0.1.0"
	defaultBaseURL     = "https://api.culqi.com/"
	userAgent          = "culqi-go/" + libraryVersion
	mediaType          = "application/json"
	format             = "json"
	headerCulqiTrackID = "x-culqi-tracking-id"
	headerEnvironment  = "x-culqi-environment"
	apiVersion         = "v2"
)

//Culqi is the complete server struct
type Culqi struct {
	Conf *Config
	HTTP *http.Client
}

//Error wraps errors returned by culqi's servers.
type Error struct {
	Object          string `json:"object"`
	Type            string `json:"type"`
	ChargeID        string `json:"charge_id"`
	Code            string `json:"code"`
	DeclineCode     string `json:"decline_code"`
	MerchantMessage string `json:"merchant_message"`
	UserMessage     string `json:"user_message"`
	Param           string `json:"param"`
	HTTPError       int
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v - %v", e.Code, e.MerchantMessage)
}

//Config takes the api keys from culqi's developer panel to authenticate with their servers.
type Config struct {
	MerchantCode string
	APIKey       string
	APIVersion   string
}

//New creates a new culqi server instance
func New(config *Config, http *http.Client) *Culqi {
	return &Culqi{
		Conf: config,
		HTTP: http,
	}
}

//DefaultWithCredentials with credentials creates a default configuration with an api key.
func DefaultWithCredentials(apiKey string) *Culqi {
	conf := &Config{
		APIKey: apiKey,
	}
	return &Culqi{
		Conf: conf,
		HTTP: http.DefaultClient,
	}
}

//WithCustomClient lets you change the client in case it's deployed with a different one, such as app engine's http client.
func (c *Culqi) WithCustomClient(http *http.Client) {
	c.HTTP = http
}

func extractError(resp *http.Response) *Error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	t := Error{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil
	}

	t.HTTPError = resp.StatusCode

	return &t
}
