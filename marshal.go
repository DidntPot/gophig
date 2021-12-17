package gophig

import (
	"io/fs"
	"os"
)

type Marshaller interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

func GetConf(marshaller Marshaller, filepath string, v interface{}) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	return marshaller.Unmarshal(data, &v)
}

func SetConf(marshaller Marshaller, filepath string, v interface{}, perm fs.FileMode) error {
	data, err := marshaller.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, data, perm)
}
