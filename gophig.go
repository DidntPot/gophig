package gophig

import (
	"io/fs"
)

type Gophig struct {
	Name, Extension string
	Perm            fs.FileMode
}

func NewGophig(name, extension string, perm fs.FileMode) *Gophig {
	return &Gophig{
		Name:      name,
		Extension: extension,
		Perm:      perm,
	}
}

func (gophig *Gophig) SetConf(v interface{}) error {
	marshaler, err := extMarshaler(gophig.Extension)
	if err != nil {
		return err
	}
	return SetConfComplex(gophig.Name+"."+gophig.Extension, marshaler, v, gophig.Perm)
}

func (gophig *Gophig) GetConf(v interface{}) error {
	marshaler, err := extMarshaler(gophig.Extension)
	if err != nil {
		return err
	}
	return GetConfComplex(gophig.Name+"."+gophig.Extension, marshaler, v)
}
