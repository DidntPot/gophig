package gophig

import (
	"io/fs"
	"os"
)

func GetConfComplex(name string, marshaler Marshaler, v interface{}) error {
	data, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	return marshaler.Unmarshal(data, v)
}

func SetConfComplex(name string, marshaler Marshaler, v interface{}, perm fs.FileMode) error {
	data, err := marshaler.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(name, data, perm)
}
