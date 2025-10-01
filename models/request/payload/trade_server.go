package payload

type PutLimit struct {
	ID     int            `json:"id"`
	Method string         `json:"method"`
	Params *[]interface{} `json:"params"`
}
