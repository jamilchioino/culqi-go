package culqi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	cardsBase = "cards"
)

type Cards struct {
	Object     string      `json:"object"`
	ID         string      `json:"id"`
	Date       int         `json:"date"`
	CustomerID string      `json:"customer_id"`
	Source     Source      `json:"source"`
	Metadata   interface{} `json:"metadata"`
}

type Issuer struct {
	Name        string `json:"name"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Website     string `json:"website"`
	PhoneNumber string `json:"phone_number"`
}

type IIN struct {
	Object              string `json:"object"`
	Bin                 string `json:"bin"`
	CardBrand           string `json:"card_brand"`
	CardCategory        string `json:"card_category"`
	Issuer              Issuer `json:"issuer"`
	InstallmentsAllowed []int  `json:"installments_allowed"`
}

type CardsParams struct {
	Amount           string           `json:"amount"`
	CurrencyCode     string           `json:"currency_code"`
	Email            string           `json:"email"`
	AntifraudDetails AntifraudDetails `json:"antifraud_details"`
	SourceID         string           `json:"source_id"`
}

func (c *Culqui) GetCard(id string) (*Charge, error) {

	req, err := http.NewRequest("GET", defaultBaseURL+"v2/"+cardsBase+id, nil)
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

	t := Charge{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *Culqui) CreateCard(params *CardsParams) (*ChargeResponse, error) {

	if params == nil {
		return nil, fmt.Errorf("no se envió parametros")
	}

	reqJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", defaultBaseURL+"v2/"+cardsBase, bytes.NewBuffer(reqJSON))
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

	t := ChargeResponse{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil

}
