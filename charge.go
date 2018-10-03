package culqi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	chargesBase = "charges"
)

type Charge struct {
	Object             string `json:"object"`
	ID                 string `json:"id"`
	Amount             int    `json:"amount"`
	AmountRefunded     int    `json:"amount_refunded"`
	CurrentAmount      int    `json:"current_amount"`
	Installments       int    `json:"installments"`
	InstallmentsAmount int    `json:"installments_amount"`
	Currency           string `json:"currency"`
	Email              string `json:"email"`
	Description        string `json:"description"`
	Source             Source `json:"source"`
}

type ChargeParams struct {
	Amount           string           `json:"amount"`
	CurrencyCode     string           `json:"currency_code"`
	Email            string           `json:"email"`
	AntifraudDetails AntifraudDetails `json:"antifraud_details"`
	SourceID         string           `json:"source_id"`
}

type ChargeResponse struct {
	Duplicated         bool              `json:"duplicated"`
	Object             string            `json:"object"`
	ID                 string            `json:"id"`
	Amount             int               `json:"amount"`
	AmountRefunded     int               `json:"amount_refunded"`
	CurrentAmount      int               `json:"current_amount"`
	Installments       int               `json:"installments"`
	InstallmentsAmount int               `json:"installments_amount"`
	Currency           string            `json:"currency"`
	Email              string            `json:"email"`
	Description        string            `json:"description"`
	Source             Source            `json:"source"`
	FraudScore         float64               `json:"fraud_score"`
	AntifraudDetails   AntifraudDetails  `json:"antifraud_details"`
	Date               int               `json:"date"`
	ReferenceCode      string            `json:"reference_code"`
	Fee                int               `json:"fee"`
	FeeDetails         FeeDetails        `json:"fee_details"`
	NetAmount          int               `json:"net_amount"`
	ResponseCode       string            `json:"response_code"`
	MerchantMessage    string            `json:"merchant_message"`
	UserMessage        string            `json:"user_message"`
	DeviceIP           string            `json:"device_ip"`
	DeviceCountry      string            `json:"device_country"`
	CountryIP          string            `json:"country_ip"`
	Product            string            `json:"product"`
	State              string            `json:"state"`
	Metadata           map[string]string `json:"metadata"`
}

type FeeDetails struct {
	Type         string  `json:"type"`
	Amount       float64 `json:"amount"`
	TotalAmount  float64 `json:"total_amount"`
	CurrencyCode string  `json:"currency_code"`
	Object       string  `json:"object"`
}

type Client struct {
	IP                string `json:"ip"`
	IPCountry         string `json:"ip_country"`
	IPCountryCode     string `json:"ip_country_code"`
	Browser           string `json:"browser"`
	DeviceFingerprint string `json:"device_fingerprint"`
	DeviceType        string `json:"device_type"`
}

type AntifraudDetails struct {
	Address     string `json:"address"`
	AddressCity string `json:"address_city"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	CountryCode string `json:"country_code"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Object      string `json:"object"`
}

func (c *Culqi) GetCharge(id string) (*Charge, error) {

	req, err := http.NewRequest("GET", defaultBaseURL+"v2/"+chargesBase+id, nil)
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

	t := Charge{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *Culqi) CreateCharge(params *ChargeParams) (*ChargeResponse, error) {

	if params == nil {
		return nil, fmt.Errorf("no se envió parametros")
	}

	reqJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", defaultBaseURL+"v2/"+chargesBase, bytes.NewBuffer(reqJSON))
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

	t := ChargeResponse{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil

}
