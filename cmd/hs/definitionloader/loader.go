//go:generate go-bindata -nometadata -pkg definitionloader ../../../examples/default/default.hs.js
package definitionloader

import (
	"fmt"
	"github.com/blang/vfs"
)

type Loader struct {
	fs               vfs.Filesystem
	file             string
	defaultLocations []string
	definition       Definition
}

type Definition struct {
	Filename string
	Dsl      []byte
}

func (d *Loader) Load(fs vfs.Filesystem, defaultMenu bool, file string) (Definition, error) {
	d.fs = fs
	var err error
	d.definition = Definition{}
	d.file = file
	if defaultMenu {
		err = d.loadDefaultMenu()
	} else if len(d.file) > 0 {
		err = d.fetchFile(d.file)
	} else if !d.loadFileFromCwd() {
		err = d.loadDefaultMenu()
	}
	return d.definition, err
}

func (d *Loader) loadDefaultMenu() error {
	var err error
	d.definition.Filename = "default.hs.js"
	d.definition.Dsl, err = Asset(fmt.Sprintf("../../../examples/default/%s", d.definition.Filename))
	return err
}

func (d *Loader) loadFileFromCwd() bool {
	return d.fetchFile("./hs.js") == nil
}

func (d *Loader) fetchFile(path string) error {
	var err error
	d.definition.Filename = path
	d.definition.Dsl, err = vfs.ReadFile(d.fs, path)
	return err
}