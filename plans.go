package culqi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	plansBase = "plans"
)

type Plan struct {
	Object        string            `json:"object"`
	ID            string            `json:"id"`
	Amount        int               `json:"amount"`
	CreationDate  int               `json:"creation_date"`
	CurrencyCode  string            `json:"currency_code"`
	Interval      string            `json:"interval"`
	IntervalCount int               `json:"interval_count"`
	Limit         int               `json:"limit"`
	Name          string            `json:"name"`
	Subscriptions []Subscription    `json:"subscriptions"`
	TrialDays     int               `json:"trial_days"`
	Metadata      map[string]string `json:"metadata"`
}

type PlansPaging struct {
	Data   []Plan `json:"data"`
	Paging Paging `json:"paging"`
}

func (c *Culqi) GetPlan(id string) (*Plan, error) {

	req, err := http.NewRequest("GET", defaultBaseURL+"v2/"+plansBase+"/"+id, nil)
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

	t := Plan{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *Culqi) AllPlans() (*PlansPaging, error) {
	req, err := http.NewRequest("GET", defaultBaseURL+"v2/"+plansBase, nil)
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

	t := PlansPaging{}

	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

//TODO: Create plans
