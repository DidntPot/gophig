package gophig

import "encoding/json"

type JSONMarshaler struct{}

func (JSONMarshaler) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (JSONMarshaler) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
