package culqi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	tokensBase = "tokens"
)

type TokenParams struct {
	Email      string `json:"email"`
	CardNumber int    `json:"card_number"`
	Cvv        int    `json:"cvv"`
	ExpMonth   int    `json:"expiration_month"`
	ExpYear    int    `json:"expiration_year"`
}

type Issuer struct {
	Name        string `json:"name"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Website     string `json:"website"`
	PhoneNumber string `json:"phone_number"`
}

type Client struct {
	IP                string `json:"ip"`
	IPCountry         string `json:"ip_country"`
	IPCountryCode     string `json:"ip_country_code"`
	Browser           string `json:"browser"`
	DeviceFingerprint string `json:"device_fingerprint"`
	DeviceType        string `json:"device_type"`
}

type IIN struct {
	Object              string `json:"object"`
	Bin                 string `json:"bin"`
	CardBrand           string `json:"card_brand"`
	CardCategory        string `json:"card_category"`
	Issuer              Issuer `json:"issuer"`
	InstallmentsAllowed []int  `json:"installments_allowed"`
}

type TokenResponse struct {
	Object       string      `json:"object"`
	ID           string      `json:"id"`
	Type         string      `json:"type"`
	Email        string      `json:"email"`
	CreationDate int         `json:"creation_date"`
	CardNumber   string      `json:"card_number"`
	LastFour     string      `json:"last_four"`
	Action       bool        `json:"action"`
	IIN          IIN         `json:"iin"`
	Client       Client      `json:"client"`
	Metadata     interface{} `json:"metadata"`
}

type CardHolder struct {
	// Object: "cardholder"
	Object    string `json:"object"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	LastName  string `json:"last_name"`
}

func (c *Culqui) Token(params *TokenParams) (*TokenResponse, error) {

	if params == nil {
		return nil, fmt.Errorf("no se envió parametros")
	}

	reqJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", defaultBaseURL+"v2/"+tokensBase, bytes.NewBuffer(reqJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.conf.APIKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	t := TokenResponse{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil

}
