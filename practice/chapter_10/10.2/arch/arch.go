package arch

import (
	"errors"
	"path/filepath"
)

type reader func(name string) error

var formats = make(map[string]reader)

func RegisterFormat(format string, fn reader)  {
	formats[format] = fn
}

func ReadArchive(name string) error {
	if fn, ok := formats[filepath.Ext(name)[1:]]; ok {
		fn(name)
		return nil
	}
	return errors.New("unsupported archive format")
}