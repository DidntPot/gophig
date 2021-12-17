package gophig

import (
	"github.com/pelletier/go-toml"
)

type TOMLMarshaler struct{}

func (TOMLMarshaler) Marshal(v interface{}) ([]byte, error) {
	return toml.Marshal(v)
}

func (TOMLMarshaler) Unmarshal(data []byte, v interface{}) error {
	return toml.Unmarshal(data, &v)
}
