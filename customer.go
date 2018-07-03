package culqi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	customerBase = "customers"
)

type Customer struct {
	Object           string           `json:"object"`
	ID               string           `json:"id"`
	CreationDate     string           `json:"creation_date"`
	Email            string           `json:"email"`
	AntifraudDetails AntifraudDetails `json:"antifraud_details"`
}

type CustomerParams struct {
	FirstName   string           `json:"first_name"`
	LastName    string           `json:"last_name"`
	Email       string           `json:"email"`
	Address     AntifraudDetails `json:"address"`
	AddressCity string           `json:"source_id"`
	CountryCode string           `json:"country_code"`
	PhoneNumber string           `json:"phone_number"`
}

func (c *Culqui) GetCustomer(id string) (*Customer, error) {

	req, err := http.NewRequest("GET", defaultBaseURL+"v2/"+customerBase+"/"+id, nil)
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

	t := Customer{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *Culqui) CreateCustomer(params *CustomerParams) (*Customer, error) {

	if params == nil {
		return nil, fmt.Errorf("no se envió parametros")
	}

	reqJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", defaultBaseURL+"v2/"+customerBase, bytes.NewBuffer(reqJSON))
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

	t := Customer{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil
}
