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

type Culqi struct {
	Conf *Config
	Http *http.Client
}

type CulqiError struct {
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

func (e *CulqiError) Error() string {
	return fmt.Sprintf("%v - %v", e.Code, e.MerchantMessage)
}

type Config struct {
	MerchantCode string
	APIKey       string
	APIVersion   string
}

func New(config *Config, http *http.Client) *Culqi {
	// set valores por defecto
	return &Culqi{
		Conf: config,
		Http: http,
	}
}

func DefaultWithCredentials(apiKey string) *Culqi {
	conf := &Config{
		APIKey: apiKey,
	}
	return &Culqi{
		Conf: conf,
		Http: http.DefaultClient,
	}
}

func (c *Culqi) WithCustomClient(http *http.Client) {
	c.Http = http
}

func extractError(resp *http.Response) *CulqiError {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	t := CulqiError{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil
	}

	t.HTTPError = resp.StatusCode

	return &t
}
