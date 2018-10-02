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

type Card struct {
	Object     string            `json:"object"`
	ID         string            `json:"id"`
	Date       int               `json:"date"`
	CustomerID string            `json:"customer_id"`
	Source     Source            `json:"source"`
	Metadata   map[string]string `json:"metadata"`
}

type Source struct {
	Object       string            `json:"object"`
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	CreationDate int               `json:"creation_date"`
	CardNumber   string            `json:"card_number"`
	LastFour     string            `json:"last_four"`
	Active       bool              `json:"active"`
	Email        string            `json:"email"`
	IIN          IIN               `json:"iin"`
	Client       Client            `json:"client"`
	Metadata     map[string]string `json:"metadata"`
	Duplicated   bool              `json:"duplicated"`
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
	CustomerID string `json:"customer_id"`
	TokenID    string `json:"token_id"`
}

func (c *Culqi) GetCard(id string) (*Charge, error) {

	req, err := http.NewRequest("GET", defaultBaseURL+"v2/"+cardsBase+id, nil)
	req.Header.Set("Authorization", "Bearer "+c.Conf.APIKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.Http.Do(req)
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

func (c *Culqi) DeleteCard(id string) error {
	if id == "" {
		return fmt.Errorf("no se envió id")
	}

	req, err := http.NewRequest("DELETE", defaultBaseURL+"v2/"+cardsBase+id, nil)
	req.Header.Set("Authorization", "Bearer "+c.Conf.APIKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.Http.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return extractError(resp)
	}

	return nil

}

func (c *Culqi) CreateCard(params *CardsParams) (*Card, error) {

	if params == nil {
		return nil, fmt.Errorf("no se envió parametros")
	}

	reqJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", defaultBaseURL+"v2/"+cardsBase, bytes.NewBuffer(reqJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Conf.APIKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, extractError(resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	t := Card{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil

}
