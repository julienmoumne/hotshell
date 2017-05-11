//go:generate go-bindata -nometadata -pkg definitionloader ../../../examples/default/default.hs.js
package definitionloader

// todo unit test me
import (
	"fmt"
	"os/user"
	"github.com/julienmoumne/hotshell/cmd/hs/filefetcher"
)

const DEF_HS_FILENAME string = "hs.js"

type DefinitionLoader struct {
	Default          bool
	File             string
	defaultLocations []string
	definition       *Definition
}

type Definition struct {
	Filename          string
	DefaultMenuLoaded bool
	Dsl               []byte
}

func (d *DefinitionLoader) LoadDefinition() (*Definition, error) {
	var err error
	d.definition = &Definition{}
	if d.Default {
		err = d.loadDefaultMenu()
	} else if len(d.File) > 0 {
		err = d.loadUserProvidedFile()
	} else if !d.loadFileFromDefaultLocations() {
		err = d.loadDefaultMenu()
	}
	return d.definition, err
}

func (d *DefinitionLoader) loadDefaultMenu() error {
	var err error
	d.definition.DefaultMenuLoaded = true
	d.definition.Filename = "../../../examples/default/default.hs.js"
	d.definition.Dsl, err = Asset(d.definition.Filename)
	return err
}

func (d *DefinitionLoader) loadFileFromDefaultLocations() bool {
	d.initDefaultLocations()
	for _, dir := range d.defaultLocations {
		if err := d.fetchFile(dir); err == nil {
			return true
		}
	}
	return false
}

func (d *DefinitionLoader) initDefaultLocations() {
	d.defaultLocations = make([]string, 1)
	d.defaultLocations[0] = fmt.Sprintf("./%s", DEF_HS_FILENAME)

	usr, err := user.Current()
	if err != nil {
		return
	}

	hsInHomeDir := fmt.Sprintf("%s/.hs/%s", usr.HomeDir, DEF_HS_FILENAME)
	d.defaultLocations = append(d.defaultLocations, hsInHomeDir)
}

func (d *DefinitionLoader) loadUserProvidedFile() error {
	return d.fetchFile(d.File)
}

func (d *DefinitionLoader) fetchFile(path string) error {
	var err error
	d.definition.Dsl, d.definition.Filename, err = (&filefetcher.Filefetcher{
		Fs:                filefetcher.NativeFS{},
		WebClient:         filefetcher.NewWebClient(),
		DefaultHSFilename: DEF_HS_FILENAME,
	}).FetchFile(path)
	return err
}
