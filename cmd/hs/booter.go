package main

// todo unit test me
import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/definitionloader"
	"github.com/julienmoumne/hotshell/cmd/hs/demo"
	"github.com/julienmoumne/hotshell/cmd/hs/doc"
	"github.com/julienmoumne/hotshell/cmd/hs/interpreter"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"os"
	"path/filepath"
	"github.com/julienmoumne/hotshell/cmd/options"
	"github.com/julienmoumne/hotshell/cmd/term"
)

type booter struct {
	options    *options.Options
	ast        []interpreter.Ast
	item       *item.Item
	term       *term.Term
	osCwd      string
	definition *definitionloader.Definition
	bootSeq    []func() error
}

// todo probably factor-out reloading logic
func (b *booter) start() error {
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

func (b *booter) initBootSeq() {
	b.bootSeq = []func() error{
		b.loadDefinitionFile,
		b.adjustWorkingDirectory,
		b.interpretDSL,
		b.buildMenu,
	}
}

func (b *booter) doStart() (bool, error) {
	if err := b.executeBootSeq(); err != nil {
		return false, err
	}
	return b.activateAction()
}

func (b *booter) executeBootSeq() error {
	for _, step := range b.bootSeq {
		if err := step(); err != nil {
			return err
		}
	}
	return nil
}

func (b *booter) activateAction() (bool, error) {
	if b.options.GenerateDemo {
		return false, b.generateDemo()
	} else if b.options.GenerateDoc {
		return false, b.generateDoc()
	} else {
		return b.startMenu()
	}
}

func (b *booter) generateDoc() error {
	gen := doc.Generator{Item: b.item, Filename: filepath.Base(b.definition.Filename)}
	return gen.Generate()
}

func (b *booter) generateDemo() error {
	gen := demo.Generator{Item: b.item, Filename: filepath.Base(b.definition.Filename)}
	return gen.Generate()
}

func (b *booter) startMenu() (bool, error) {
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

func (b *booter) loadDefinitionFile() error {
	var loader = definitionloader.DefinitionLoader{Default: b.options.Default, File: b.options.File}
	var err error
	b.definition, err = loader.LoadDefinition()
	return err
}

// todo factor out cwd logic
func (b *booter) chdir() bool {
	return b.options.Chdir && !b.definition.DefaultMenuLoaded
}

func (b *booter) adjustWorkingDirectory() error {
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

func (b *booter) restoreCwd() error {
	if !b.chdir() {
		return nil
	}
	return os.Chdir(b.osCwd)
}

func (b *booter) buildMenu() error {
	var err error
	b.item, err = (&item.Builder{}).Build(b.ast)
	return err
}

func (b *booter) startController() (bool, error) {
	return (&controller{root: b.item, term: b.term}).start()
}

func (b *booter) initTerm() error {
	var err error
	b.term, err = term.NewTerm()
	return err
}

func (b *booter) interpretDSL() error {
	inter := interpreter.Interpreter{Filename: b.definition.Filename, Dsl: b.definition.Dsl}
	var err error
	b.ast, err = inter.Interpret()
	return err
}
