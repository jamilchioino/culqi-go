// Package culqi implementa metodos para inicializar cliente
//
// Inicia cliente con datos de condifuración de comercio.
package culqi

import "net/http"

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

type Culqui struct {
	conf *Config
	http *http.Client
}

type Config struct {
	MerchantCode string
	APIKey       string
	APIVersion   string
}

func New(config *Config, http *http.Client) *Culqui {
	// set valores por defecto
	return &Culqui{
		conf: config,
		http: http,
	}
}

func DefaultWithCredentials(apiKey string) *Culqui {
	conf := &Config{
		APIKey: apiKey,
	}
	return &Culqui{
		conf: conf,
		http: http.DefaultClient,
	}
}

func (c *Culqui) WithCustomClient(http *http.Client) {
	c.http = http
}
