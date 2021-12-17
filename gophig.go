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
	return SetConf(gophig.Name, gophig.Extension, v, gophig.Perm)
}

func (gophig *Gophig) GetConf(v interface{}) error {
	return GetConf(gophig.Name, gophig.Extension, v)
}
