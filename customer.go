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
	CreationDate     int              `json:"creation_date"`
	Email            string           `json:"email"`
	AntifraudDetails AntifraudDetails `json:"antifraud_details"`
	Cards            []Card           `json:"cards"`
}

type CustomerParams struct {
	FirstName   *string                `json:"first_name,omitempty"`
	LastName    *string                `json:"last_name,omitempty"`
	Email       *string                `json:"email,omitempty"`
	Address     *string                `json:"address,omitempty"`
	AddressCity *string                `json:"address_city,omitempty"`
	CountryCode *string                `json:"country_code,omitempty"`
	PhoneNumber *string                `json:"phone_number,omitempty"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type DeletedCustomer struct {
	ID              string `json:"id"`
	Deleted         bool   `json:"deleted"`
	MerchantMessage string `json:"merchant_message"`
}

type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type Paging struct {
	Previous string  `json:"previous"`
	Next     string  `json:"next"`
	Cursors  Cursors `json:"cursors"`
}

type CustomerPaging struct {
	Data   []Customer `json:"data"`
	Paging Paging     `json:"paging"`
}

func (c *Culqi) GetCustomer(id string) (*Customer, error) {

	req, err := http.NewRequest("GET", defaultBaseURL+"v2/"+customerBase+"/"+id, nil)
	req.Header.Set("Authorization", "Bearer "+c.Conf.APIKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.Http.Do(req)
	if resp.StatusCode >= 400 {
		return nil, extractError(resp)
	}

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

func (c *Culqi) CreateCustomer(params *CustomerParams) (*Customer, error) {

	if params == nil {
		return nil, fmt.Errorf("params are empty")
	}

	reqJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", defaultBaseURL+"v2/"+customerBase, bytes.NewBuffer(reqJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Conf.APIKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.Http.Do(req)

	if resp.StatusCode >= 400 {
		return nil, extractError(resp)
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

func (c *Culqi) AllCustomers() (*CustomerPaging, error) {
	req, err := http.NewRequest("GET", defaultBaseURL+"v2/"+customerBase, nil)
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

	t := CustomerPaging{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *Culqi) DeleteCustomer(id string) (*DeletedCustomer, error) {
	req, err := http.NewRequest("DELETE", defaultBaseURL+"v2/"+customerBase+"/"+id, nil)
	req.Header.Set("Authorization", "Bearer "+c.Conf.APIKey)
	req.Header.Set("User-Agent", userAgent)

	if id == "" {
		return nil, fmt.Errorf("id not specified")
	}

	resp, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	t := DeletedCustomer{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *Culqi) UpdateCustomer(id string, params *CustomerParams) (*Customer, error) {

	if id == "" {
		return nil, fmt.Errorf("id not specified")
	}

	if params == nil {
		return nil, fmt.Errorf("params are empty")
	}

	reqJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", defaultBaseURL+"v2/"+customerBase+"/"+id, bytes.NewBuffer(reqJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Conf.APIKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.Http.Do(req)

	if resp.StatusCode >= 400 {
		return nil, extractError(resp)
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
