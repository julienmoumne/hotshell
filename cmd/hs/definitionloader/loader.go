//go:generate go-bindata -nometadata -pkg definitionloader ../../../examples/default/default.hs.js
//go:generate mockery -name UserGetter -inpkg -case underscore
package definitionloader

import (
	"fmt"
	"github.com/blang/vfs"
	"os/user"
)

const defaultFilename = "hs.js"

type UserGetter interface {
	Get() (*user.User, error)
}

type Loader struct {
	fs               vfs.Filesystem
	userGetter       UserGetter
	file             string
	defaultLocations []string
	definition       Definition
}

type Definition struct {
	Filename string
	Dsl      []byte
}

func (d *Loader) Load(fs vfs.Filesystem, userGetter UserGetter, defaultMenu bool, file string) (Definition, error) {
	d.fs = fs
	d.userGetter = userGetter
	var err error
	d.definition = Definition{}
	d.file = file
	if defaultMenu {
		err = d.loadDefaultMenu()
	} else if len(d.file) > 0 {
		err = d.fetchFile(d.file)
	} else if !d.loadFileFromDefaultLocations() {
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

func (d *Loader) loadFileFromDefaultLocations() bool {
	d.initDefaultLocations()
	for _, loc := range d.defaultLocations {
		if err := d.fetchFile(loc); err == nil {
			return true
		}
	}
	return false
}

func (d *Loader) initDefaultLocations() {
	d.defaultLocations = make([]string, 1)
	d.defaultLocations[0] = fmt.Sprintf("./%s", defaultFilename)

	usr, err := d.userGetter.Get()
	if err != nil {
		return
	}

	hsInHomeDir := fmt.Sprintf("%s/.hs/%s", usr.HomeDir, defaultFilename)
	d.defaultLocations = append(d.defaultLocations, hsInHomeDir)
}

func (d *Loader) fetchFile(path string) error {
	var err error
	d.definition.Filename = path
	d.definition.Dsl, err = vfs.ReadFile(d.fs, path)
	return err
}