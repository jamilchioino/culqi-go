package culqi

import (
	"gopkg.in/resty.v1"
)

const (
	chargesBase = "charges"
)

type ChargeParams struct {
	TokenID            string `json:"source_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	Address            string `json:"address"`
	AddressCity        string `json:"address_city"`
	PhoneNumber        int    `json:"phone_number"`
	CountryCode        string `json:"country_code"`
	CurrencyCode       string `json:"currency_code"`
	Amount             int    `json:"amount"`
	Installments       int    `json:"installments"`
	ProductDescription string `json:"product_description"`
}

func (c *Culqui) Charge(params *ChargeParams) (*resty.Response, error) {

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.conf.APIKey).
		SetHeader("User-Agent", userAgent).
		SetBody(params).
		Post(defaultBaseURL + "v2/" + chargesBase)
	return resp, err

}
