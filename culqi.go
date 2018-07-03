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
	conf *Config
	http *http.Client
}

type CulqiError struct {
	Object          string
	Type            string
	ChargeID        string
	Code            string
	DeclineCode     string
	MerchantMessage string
	UserMessage     string
	Param           string
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
		conf: config,
		http: http,
	}
}

func DefaultWithCredentials(apiKey string) *Culqi {
	conf := &Config{
		APIKey: apiKey,
	}
	return &Culqi{
		conf: conf,
		http: http.DefaultClient,
	}
}

func (c *Culqi) WithCustomClient(http *http.Client) {
	c.http = http
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

	return &t
}
