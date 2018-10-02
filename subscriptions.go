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
	Object             string                 `json:"object"`
	ID                 string                 `json:"id"`
	CreationDate       int                    `json:"creation_date"`
	Status             string                 `json:"status"`
	CurrentPeriod      int                    `json:"current_period"`
	TotalPeriods       int                    `json:"total_periods"`
	CurrentPeriodStart int                    `json:"current_period_start"`
	CurrentPeriodEnd   int                    `json:"current_period_end"`
	CancelAtPeriodEnd  bool                   `json:"cancel_at_period_end"`
	CanceledAt         int                    `json:"canceled_at"`
	EndedAt            int                    `json:"ended_at"`
	NextBillingDate    int                    `json:"next_billing_date"`
	TrialStart         int                    `json:"trial_start"`
	TrialEnd           int                    `json:"trial_end"`
	Charges            []Charge               `json:"charges"`
	Plan               Plan                   `json:"plan"`
	Card               Card                   `json:"card"`
	Metadata           map[string]interface{} `json:"metadata"`
}

type SubscriptionParams struct {
	CardID string `json:"card_id"`
	PlanID string `json:"plan_id"`
}

func (c *Culqi) GetSubscription(id string) (*Subscription, error) {

	req, err := http.NewRequest("GET", defaultBaseURL+"v2/"+subscriptionsBase+"/"+id, nil)
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

	t := Subscription{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *Culqi) CreateSubscription(params *SubscriptionParams) (*Subscription, error) {

	if params == nil {
		return nil, fmt.Errorf("no se envió parametros")
	}

	reqJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", defaultBaseURL+"v2/"+subscriptionsBase, bytes.NewBuffer(reqJSON))
	req.Header.Set("Content-Type", "application/json")
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

	t := Subscription{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *Culqi) DeleteSubscription(id string) error {
	req, err := http.NewRequest("DELETE", defaultBaseURL+"v2/"+subscriptionsBase+"/"+id, nil)
	req.Header.Set("Authorization", "Bearer "+c.Conf.APIKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.Http.Do(req)
	if resp.StatusCode >= 400 {
		return extractError(resp)
	}

	if err != nil {
		return err
	}

	return nil
}
