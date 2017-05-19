//go:generate go-bindata -nometadata -pkg definitionloader ../../../examples/default/default.hs.js
package definitionloader

import (
	"fmt"
	"github.com/blang/vfs"
	"github.com/julienmoumne/hotshell/cmd/hs/fileloader"
	"os/user"
)

const defaultFilename = "hs.js"

type DefinitionLoader struct {
	FileLoader       fileloader.FileLoader
	Fs               vfs.Filesystem
	file             string
	defaultLocations []string
	definition       Definition
}

type Definition struct {
	Filename          string
	DefaultMenuLoaded bool
	Dsl               []byte
}

var Default = DefinitionLoader{
	FileLoader: fileloader.Default,
	Fs:         vfs.ReadOnly(vfs.OS()),
}

func (d *DefinitionLoader) Load(defaultMenu bool, file string) (Definition, error) {
	var err error
	d.definition = Definition{}
	d.file = file
	if defaultMenu {
		err = d.loadDefaultMenu()
	} else if len(d.file) > 0 {
		err = d.loadUserProvidedFile()
	} else if !d.loadFileFromDefaultLocations() {
		err = d.loadDefaultMenu()
	}
	return d.definition, err
}

func (d *DefinitionLoader) loadDefaultMenu() error {
	var err error
	d.definition.DefaultMenuLoaded = true
	d.definition.Filename = "default.hs.js"
	d.definition.Dsl, err = Asset(fmt.Sprintf("../../../examples/default/%s", d.definition.Filename))
	return err
}

func (d *DefinitionLoader) loadFileFromDefaultLocations() bool {
	d.initDefaultLocations()
	for _, loc := range d.defaultLocations {
		if err := d.fetchFile(loc); err == nil {
			return true
		}
	}
	return false
}

func (d *DefinitionLoader) initDefaultLocations() {
	d.defaultLocations = make([]string, 1)
	d.defaultLocations[0] = fmt.Sprintf("./%s", defaultFilename)

	usr, err := user.Current()
	if err != nil {
		return
	}

	hsInHomeDir := fmt.Sprintf("%s/.hs/%s", usr.HomeDir, defaultFilename)
	d.defaultLocations = append(d.defaultLocations, hsInHomeDir)
}

func (d *DefinitionLoader) loadUserProvidedFile() error {
	isDir, err := d.userProvidedFileIsDir()
	if err != nil {
		return err
	}
	if isDir {
		d.file += fmt.Sprintf("/%s", defaultFilename)
	}
	return d.fetchFile(d.file)
}

func (d *DefinitionLoader) userProvidedFileIsDir() (bool, error) {
	info, err := d.Fs.Stat(d.file)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

func (d *DefinitionLoader) fetchFile(path string) error {
	var err error
	d.definition.Filename = path
	d.definition.Dsl, err = d.FileLoader.Load(path)
	return err
}
