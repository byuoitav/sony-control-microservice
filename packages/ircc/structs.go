package ircc

type Capability struct {
	ID     int           `json:"id"`
	Result []interface{} `json:"result"`
}
