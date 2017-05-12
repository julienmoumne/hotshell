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
	Options    *options.Options
	ast        []interpreter.Ast
	item       *item.Item
	term       *term.Term
	osCwd      string
	definition *definitionloader.Definition
	bootSeq    []func() error
}

// todo probably factor-out reloading logic
func (b *Starter) Start() error {
	b.initBootSeq()
	for reload, err := b.doStart(); err != nil || reload; reload, err = b.doStart() {
		if err != nil {
			return err
		}
		if err := b.restoreCwd(); err != nil {
			return err
		}
	}
	return nil
}

func (b *Starter) initBootSeq() {
	b.bootSeq = []func() error{
		b.loadDefinitionFile,
		b.adjustWorkingDirectory,
		b.interpretDSL,
		b.buildMenu,
	}
}

func (b *Starter) doStart() (bool, error) {
	if err := b.executeBootSeq(); err != nil {
		return false, err
	}
	return b.activateAction()
}

func (b *Starter) executeBootSeq() error {
	for _, step := range b.bootSeq {
		if err := step(); err != nil {
			return err
		}
	}
	return nil
}

func (b *Starter) activateAction() (bool, error) {
	if b.Options.GenerateDemo {
		return false, b.generateDemo()
	} else if b.Options.GenerateMd {
		return false, b.generateMd()
	} else {
		return b.startMenu()
	}
}

func (b *Starter) generateMd() error {
	gen := generator.Md{Item: b.item, Filename: filepath.Base(b.definition.Filename)}
	return gen.Generate()
}

func (b *Starter) generateDemo() error {
	gen := generator.Demo{Item: b.item, Filename: filepath.Base(b.definition.Filename)}
	return gen.Generate()
}

func (b *Starter) startMenu() (bool, error) {
	if err := b.initTerm(); err != nil {
		return false, err
	}

	defer func() {
		if err := b.term.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	return b.startController()
}

func (b *Starter) loadDefinitionFile() error {
	var loader = definitionloader.DefinitionLoader{Default: b.Options.Default, File: b.Options.File}
	var err error
	b.definition, err = loader.LoadDefinition()
	return err
}

// todo factor out cwd logic
func (b *Starter) chdir() bool {
	return b.Options.Chdir && !b.definition.DefaultMenuLoaded
}

func (b *Starter) adjustWorkingDirectory() error {
	if !b.chdir() {
		return nil
	}
	var err error
	b.osCwd, err = os.Getwd()
	if err != nil {
		return err
	}
	return os.Chdir(filepath.Dir(b.definition.Filename))
}

func (b *Starter) restoreCwd() error {
	if !b.chdir() {
		return nil
	}
	return os.Chdir(b.osCwd)
}

func (b *Starter) buildMenu() error {
	var err error
	b.item, err = (&item.Builder{}).Build(b.ast)
	return err
}

func (b *Starter) startController() (bool, error) {
	return (&controller{root: b.item, term: b.term}).start()
}

func (b *Starter) initTerm() error {
	var err error
	b.term, err = term.NewTerm()
	return err
}

func (b *Starter) interpretDSL() error {
	inter := interpreter.Interpreter{Filename: b.definition.Filename, Dsl: b.definition.Dsl}
	var err error
	b.ast, err = inter.Interpret()
	return err
}
