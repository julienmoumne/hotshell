package engine

// todo unit test me
import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/definitionloader"
	"github.com/julienmoumne/hotshell/cmd/hs/generator"
	"github.com/julienmoumne/hotshell/cmd/hs/interpreter"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/julienmoumne/hotshell/cmd/options"
	"github.com/julienmoumne/hotshell/cmd/term"
	"os"
	"path/filepath"
)

type Starter struct {
	Options    options.Options
	ast        []interpreter.Ast
	item       *item.Item
	term       term.Term
	osCwd      string
	definition definitionloader.Definition
	bootSeq    []func() error
}

// todo probably factor-out reloading logic
func (s *Starter) Start() error {
	s.initBootSeq()
	for reload, err := s.doStart(); err != nil || reload; reload, err = s.doStart() {
		if err != nil {
			return err
		}
		if err := s.restoreCwd(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Starter) initBootSeq() {
	s.bootSeq = []func() error{
		s.loadDefinitionFile,
		s.adjustWorkingDirectory,
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
	if s.Options.GenerateDemo {
		return false, s.generateDemo()
	} else if s.Options.GenerateMd {
		return false, s.generateMd()
	} else {
		return s.startMenu()
	}
}

func (s *Starter) generateMd() error {
	gen := generator.Md{Item: s.item, Filename: filepath.Base(s.definition.Filename)}
	return gen.Generate()
}

func (s *Starter) generateDemo() error {
	gen := generator.Demo{Item: s.item, Filename: filepath.Base(s.definition.Filename)}
	return gen.Generate()
}

func (s *Starter) startMenu() (bool, error) {
	if err := s.initTerm(); err != nil {
		return false, err
	}

	defer func() {
		if err := s.term.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	return s.startController()
}

func (s *Starter) loadDefinitionFile() error {
	var err error
	s.definition, err = definitionloader.Default.Load(s.Options.Default, s.Options.File)
	return err
}

// todo factor out cwd logic
func (s *Starter) chdir() bool {
	return s.Options.Chdir && !s.definition.DefaultMenuLoaded
}

func (s *Starter) adjustWorkingDirectory() error {
	if !s.chdir() {
		return nil
	}
	var err error
	s.osCwd, err = os.Getwd()
	if err != nil {
		return err
	}
	return os.Chdir(filepath.Dir(s.definition.Filename))
}

func (s *Starter) restoreCwd() error {
	if !s.chdir() {
		return nil
	}
	return os.Chdir(s.osCwd)
}

func (s *Starter) buildMenu() error {
	var err error
	s.item, err = (&item.Builder{}).Build(s.ast)
	return err
}

func (s *Starter) startController() (bool, error) {
	return (&controller{root: s.item, term: s.term}).start()
}

func (s *Starter) initTerm() error {
	var err error
	s.term, err = term.NewTerm()
	return err
}

func (s *Starter) interpretDSL() error {
	inter := interpreter.Interpreter{Filename: s.definition.Filename, Dsl: s.definition.Dsl}
	var err error
	s.ast, err = inter.Interpret()
	return err
}
