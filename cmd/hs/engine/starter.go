package engine

// todo unit test me
import (
	"github.com/blang/vfs"
	"github.com/julienmoumne/hotshell/cmd/hs/definitionloader"
	"github.com/julienmoumne/hotshell/cmd/hs/generator"
	"github.com/julienmoumne/hotshell/cmd/hs/interpreter"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/julienmoumne/hotshell/cmd/options"
	"path/filepath"
)

type Starter struct {
	options    options.Options
	ast        []interpreter.Ast
	item       *item.Item
	osCwd      string
	definition definitionloader.Definition
	bootSeq    []func() error
}

func (s *Starter) Start(options options.Options) error {
	s.options = options
	s.initBootSeq()
	for reload, err := s.doStart(); err != nil || reload; reload, err = s.doStart() {
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Starter) initBootSeq() {
	s.bootSeq = []func() error{
		s.loadDefinitionFile,
		s.interpretDSL,
		s.buildMenu,
	}
}

func (s *Starter) doStart() (bool, error) {
	if err := s.executeBootSeq(); err != nil {
		return false, err
	}
	return s.activateAction()
}

func (s *Starter) executeBootSeq() error {
	for _, step := range s.bootSeq {
		if err := step(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Starter) activateAction() (bool, error) {
	if s.options.GenerateDemo {
		return false, s.generateDemo()
	} else if s.options.GenerateMd {
		return false, s.generateMd()
	} else {
		return s.startController()
	}
}

func (s *Starter) generateMd() error {
	return (&generator.Md{}).Generate(s.item, filepath.Base(s.definition.Filename))
}

func (s *Starter) generateDemo() error {
	return (&generator.Demo{}).Generate(s.item, filepath.Base(s.definition.Filename))
}

func (s *Starter) loadDefinitionFile() error {
	var err error
	s.definition, err = (&definitionloader.Loader{}).Load(
		vfs.ReadOnly(vfs.OS()),
		s.options.Default, s.options.File,
	)
	return err
}

func (s *Starter) buildMenu() error {
	var err error
	s.item, err = (&item.Builder{}).Build(s.ast)
	return err
}

func (s *Starter) startController() (bool, error) {
	return (&controller{}).Start(s.item)
}

func (s *Starter) interpretDSL() error {
	var err error
	s.ast, err = (&interpreter.Interpreter{}).Interpret(s.definition.Dsl)
	return err
}
