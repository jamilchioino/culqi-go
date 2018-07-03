package culqi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	subscriptionsBase = "subscriptions"
)

type Subscription struct {
	Object          string      `json:"object"`
	ID              string      `json:"id"`
	CreationDate    int         `json:"creation_date"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	Address         string      `json:"address"`
	AddressCity     string      `json:"address_city"`
	CountryCode     string      `json:"country_code"`
	Token           Token       `json:"token"`
	Email           string      `json:"email"`
	Phone           int         `json:"phone"`
	PlanID          string      `json:"plan_id"`
	Charges         []Charge    `json:"charges"`
	Current         int         `json:"current"`
	NextBillingDate int         `json:"next_billing_date"`
	metadata        interface{} `json:"metadata"`
}

type SubscriptionParams struct {
	CardID string `json:"card_id"`
	PlanID string `json:"plan_id"`
}

func (c *Culqui) GetSubscription(id string) (*Subscription, error) {

	req, err := http.NewRequest("GET", defaultBaseURL+"v2/"+subscriptionsBase+id, nil)
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

	t := Subscription{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *Culqui) CreateSubscription(params *ChargeParams) (*ChargeResponse, error) {

	if params == nil {
		return nil, fmt.Errorf("no se envió parametros")
	}

	reqJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", defaultBaseURL+"v2/"+subscriptionsBase, bytes.NewBuffer(reqJSON))
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
