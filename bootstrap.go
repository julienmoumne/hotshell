package hotshell

import (
	"fmt"
	"github.com/julienmoumne/hotshell/demo"
	"github.com/julienmoumne/hotshell/interpreter"
	"github.com/julienmoumne/hotshell/item"
	"os"
	"path/filepath"
)

type Bootstrap struct {
	options *options
	ast     []interpreter.Ast
	item    *item.Item
	term    *term
	osCwd   string
}

func (b *Bootstrap) Boot() error {
	for reload, err := b.doBoot(); err != nil || reload; reload, err = b.doBoot() {
		if err != nil {
			return err
		}
		if err := b.restoreCwd(); err != nil {
			return err
		}
	}
	return nil
}

func (b *Bootstrap) doBoot() (bool, error) {
	if err := b.initOptions(); err != nil {
		return false, err
	}

	if err := b.interpretDSL(); err != nil {
		return false, err
	}

	if err := b.buildMenu(); err != nil {
		return false, err
	}

	if b.options.flags.GenerateDemo {
		return false, b.generateDemo()
	} else {
		return b.bootMenu()
	}
}

func (b *Bootstrap) generateDemo() error {
	gen := demo.Generator{Item: b.item, Filename: filepath.Base(b.options.filename)}
	return gen.Generate()
}

func (b *Bootstrap) bootMenu() (bool, error) {

	if err := b.initTerm(); err != nil {
		return false, err
	}

	defer func() {
		if err := b.term.close(); err != nil {
			fmt.Println(err)
		}
	}()

	return b.startController()
}

func (b *Bootstrap) initOptions() error {

	options, err := newOptions()
	if err != nil {
		return err
	}

	b.options = options

	if err := b.adjustWorkingDirectory(); err != nil {
		return err
	}

	return nil
}

func (b *Bootstrap) chdir() bool {
	return b.options.flags.Chdir && !b.options.defaultMenuLoaded
}

func (b *Bootstrap) adjustWorkingDirectory() error {
	if !b.chdir() {
		return nil
	}
	var err error
	b.osCwd, err = os.Getwd()
	if err != nil {
		return err
	}
	return os.Chdir(filepath.Dir(b.options.filename))
}

func (b *Bootstrap) restoreCwd() error {
	if !b.chdir() {
		return nil
	}
	return os.Chdir(b.osCwd)
}

func (b *Bootstrap) buildMenu() error {
	var err error
	b.item, err = (&item.Builder{}).Build(b.ast)
	return err
}

func (b *Bootstrap) startController() (bool, error) {

	ctrl := controller{root: b.item, term: b.term}

	return ctrl.start()
}

func (b *Bootstrap) initTerm() error {

	term, err := NewTerm()
	b.term = term
	return err
}

func (b *Bootstrap) interpretDSL() error {

	interpreter := interpreter.Interpreter{Filename: b.options.filename, Dsl: b.options.dsl}
	result, err := interpreter.Interpret()
	b.ast = result

	return err
}
