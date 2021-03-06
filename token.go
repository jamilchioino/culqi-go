package culqi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Token is used to tokenize a card. It can only be done through culqi's servers due to card storage laws.
type Token struct {
	Object       string            `json:"object"`
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	Email        string            `json:"email"`
	CreationDate int               `json:"creation_date"`
	CardNumber   string            `json:"card_number"`
	LastFour     string            `json:"last_four"`
	Active       bool              `json:"active"`
	IIN          IIN               `json:"iin"`
	Client       Client            `json:"client"`
	Metadata     map[string]string `json:"metadata"`
}

//GetToken retrieves card tokens from the culqi server.
func (c *Culqi) GetToken(id string) (*Charge, error) {

	req, err := http.NewRequest("GET", defaultBaseURL+"v2/"+chargesBase+id, nil)
	req.Header.Set("Authorization", "Bearer "+c.Conf.APIKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.HTTP.Do(req)
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
