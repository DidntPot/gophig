package gophig

import (
	"io/fs"
)

type Gophig struct {
	Marshaller Marshaller
	FilePath   string
	Perm       fs.FileMode
}

func NewGophig(marshaller Marshaller, filepath string, perm fs.FileMode) *Gophig {
	return &Gophig{
		Marshaller: marshaller,
		FilePath:   filepath,
		Perm:       perm,
	}
}

func (gophig *Gophig) SetConf(v interface{}) error {
	return SetConf(gophig.Marshaller, gophig.FilePath, v, gophig.Perm)
}

func (gophig *Gophig) GetConf(v interface{}) error {
	return GetConf(gophig.Marshaller, gophig.FilePath, v)
}
