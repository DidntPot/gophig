package gophig

import (
	"io/fs"
	"os"
)

func GetConf(name, extension string, v interface{}) error {
	data, err := os.ReadFile(name + "." + extension)
	if err != nil {
		return err
	}
	marshaler, err := extMarshaler(extension)
	if err != nil {
		return err
	}
	return marshaler.Unmarshal(data, &v)
}

func SetConf(name, extension string, v interface{}, perm fs.FileMode) error {
	marshaler, err := extMarshaler(extension)
	if err != nil {
		return err
	}
	data, err := marshaler.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(name+"."+extension, data, perm)
}
