package culqi

//Event is used to receive webhook data. Unfinished.
type Event struct {
	Object       string                 `json:"object"`
	ID           string                 `json:"id"`
	Type         string                 `json:"type"`
	CreationDate int                    `json:"creation_date"`
	Data         map[string]interface{} `json:"data"`
}
