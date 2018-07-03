package culqi

type TokenParams struct {
	Email      string `json:"email"`
	CardNumber int    `json:"card_number"`
	Cvv        int    `json:"cvv"`
	ExpMonth   int    `json:"expiration_month"`
	ExpYear    int    `json:"expiration_year"`
}

type Token struct {
	Object       string `json:"object"`
	ID           string `json:"id"`
	Type         string `json:"type"`
	Email        string `json:"email"`
	CreationDate int    `json:"creation_date"`
	CardNumber   string `json:"card_number"`
	LastFour     string `json:"last_four"`
	Active       bool   `json:"active"`
	IIN          IIN    `json:"iin"`
	Client       Client `json:"client"`
}

type CardHolder struct {
	// Object: "cardholder"
	Object    string `json:"object"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	LastName  string `json:"last_name"`
}
