//go:generate go-bindata -nometadata -pkg hs examples/default/default.hs.js VERSION
package hs

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/julienmoumne/hs/filefetcher"
	"os"
	"os/user"
)

const DEF_HS_FILENAME string = "hs.js"

type options struct {
	filename          string
	dsl               []byte
	defaultLocations  []string
	defaultMenuLoaded bool
	flags             struct {
		Default      bool   `long:"default" description:"Load the default menu to get to know Hotshell"`
		File         string `short:"f" long:"file" description:"Specify an alternate definition file (defaults: ./hs.js, ~/.hs/hs.js)" value-name:"FILE|DIR"`
		Chdir        bool   `long:"chdir" description:"Set the working directory to the location of the menu definition"`
		GenerateDemo bool   `long:"generate-demo" description:"Generate an interactive HTML demo to stdout"`
		Version      bool   `short:"v" long:"version" description:"Print version information and quit"`
	}
}

func CreateOptionsParser() (*flags.Parser, *options) {
	options := options{}
	parser := flags.NewParser(&options.flags, flags.HelpFlag)

	parser.Name = "hs"
	parser.Usage = "[-f <arg>...] [options]"
	parser.ShortDescription = "Command-line productivity enhancer"
	parser.LongDescription = "Hotshell is a command-line application to efficiently recall and share commands."

	return parser, &options
}

func newOptions() (*options, error) {

	var parser *flags.Parser
	parser, options := CreateOptionsParser()

	if _, err := parser.Parse(); err != nil {
		return nil, err
	}

	if options.flags.Version {
		v, err := Asset("VERSION")
		if err != nil {
			return nil, err
		}
		fmt.Printf("Hotshell version %s\n", v)
		os.Exit(0)
	}

	err := options.loadFile()

	return options, err
}

func (o *options) loadFile() error {
	if o.flags.Default {
		return o.loadDefaultMenu()
	}
	if len(o.flags.File) > 0 {
		if err := o.loadUserProvidedFile(); err != nil {
			return err
		}
		return nil
	}
	if o.loadFileFromDefaultLocations() {
		return nil
	}
	if err := o.loadDefaultMenu(); err != nil {
		return err
	}
	return nil
}

func (o *options) loadDefaultMenu() error {
	o.filename = "examples/default/default.hs.js"
	dsl, err := Asset(o.filename)
	if err != nil {
		return err
	}
	o.dsl = dsl
	o.defaultMenuLoaded = true
	return nil
}

func (o *options) loadFileFromDefaultLocations() bool {
	o.initDefaultLocations()

	for _, dir := range o.defaultLocations {
		if err := o.fetchFile(dir); err == nil {
			return true
		}
	}

	return false
}

func (o *options) initDefaultLocations() {
	o.defaultLocations = make([]string, 1)
	o.defaultLocations[0] = fmt.Sprintf("./%s", DEF_HS_FILENAME)

	usr, err := user.Current()
	if err != nil {
		return
	}

	hsInHomeDir := fmt.Sprintf("%s/.hs/%s", usr.HomeDir, DEF_HS_FILENAME)
	o.defaultLocations = append(o.defaultLocations, hsInHomeDir)
}

func (o *options) loadUserProvidedFile() error {
	return o.fetchFile(o.flags.File)
}

func (o *options) fetchFile(path string) error {
	ff := filefetcher.Filefetcher{
		Fs:                      filefetcher.NativeFS{},
		WebClient:               filefetcher.NewWebClient(),
		DefaultHSFilename: DEF_HS_FILENAME,
	}
	content, filename, err := ff.FetchFile(path)
	o.dsl = content
	o.filename = filename
	return err
}
