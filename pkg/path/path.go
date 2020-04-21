package path

import (
	"path/filepath"
)

type Path string

// Decode decodes a string path to a path
func (p *Path) Decode(value string) error {
	*p = Path(filepath.Clean(filepath.FromSlash(value)))
	return nil
}

// String returns string path
func (p *Path) String() string {
	return string(*p)
}
