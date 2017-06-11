package engine

// todo unit test me
import (
	"fmt"
	"github.com/blang/vfs"
	"github.com/julienmoumne/hotshell/cmd/hs/definitionloader"
	"github.com/julienmoumne/hotshell/cmd/hs/documentor"
	"github.com/julienmoumne/hotshell/cmd/hs/dslrunner"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/julienmoumne/hotshell/cmd/hs/settings"
	"github.com/julienmoumne/hotshell/cmd/options"
	"github.com/julienmoumne/hotshell/cmd/term"
	"path/filepath"
)

type Starter struct {
	options    options.Options
	item       *item.Item
	osCwd      string
	definition definitionloader.Definition
	bootSeq    []func() error
	settings   settings.Settings
}

func (s *Starter) Start(options options.Options) error {
	s.options = options
	s.initBootSeq()
	if s.options.GenerateDemo {
		return s.bootAndStart(s.generateDemo)
	} else if s.options.GenerateMd {
		return s.bootAndStart(s.generateMd)
	}
	return s.startControllerWithAutoReload()
}

func (s *Starter) startControllerWithAutoReload() error {
	var reload bool
	if err := s.bootAndStart(func() (err error) {
		reload, err = s.startController()
		return
	}); err != nil {
		return err
	}
	if reload {
		return s.startControllerWithAutoReload()
	}
	return nil
}

func (s *Starter) initBootSeq() {
	s.bootSeq = []func() error{
		s.loadSettings,
		s.loadDefinitionFile,
		s.interpretDSL,
	}
}

func (s *Starter) loadSettings() (err error) {
	s.settings, err = (&settings.Loader{}).Load(vfs.ReadOnly(vfs.OS()))
	return
}

func (s *Starter) bootAndStart(f func() error) error {
	if err := s.executeBootSeq(); err != nil {
		return err
	}
	return f()
}

func (s *Starter) executeBootSeq() error {
	for _, step := range s.bootSeq {
		if err := step(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Starter) generateMd() error {
	return (&documentor.Md{}).Generate(s.item, filepath.Base(s.definition.Filename))
}

func (s *Starter) generateDemo() error {

	return (&documentor.Demo{}).Generate(s.item, filepath.Base(s.definition.Filename))
}

func (s *Starter) loadDefinitionFile() (err error) {
	s.definition, err = (&definitionloader.Loader{}).Load(
		vfs.ReadOnly(vfs.OS()),
		s.options.Default, s.options.File,
	)
	return
}

func (s *Starter) startController() (bool, error) {
	t, err := term.NewTerm()
	if err != nil {
		return false, err
	}
	defer func() {
		if err := t.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Print("\n")
	return (&controller{}).Start(s.settings.Keys, s.item, t)

}

func (s *Starter) interpretDSL() (err error) {
	s.item, err = (&dslrunner.DslRunner{}).Run(string(s.definition.Dsl))
	return
}
