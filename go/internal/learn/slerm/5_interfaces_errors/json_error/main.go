package homework

import "encoding/json"

type Example struct {
	A int    `json:"a"`
	B string `json:"B,omitempty"`
}

func jsonToStruct(s []byte) (*Example, error) {

	d := &Example{}

	err := json.Unmarshal(s, d)
	if err != nil {
		return nil, err
	}

	return d, nil
}
