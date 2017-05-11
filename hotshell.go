package hotshell

// todo unit test me
import (
	"fmt"
	"github.com/julienmoumne/hotshell/definitionloader"
	"github.com/julienmoumne/hotshell/demo"
	"github.com/julienmoumne/hotshell/doc"
	"github.com/julienmoumne/hotshell/interpreter"
	"github.com/julienmoumne/hotshell/item"
	"os"
	"path/filepath"
)

type Hotshell struct {
	Options    *Options
	ast        []interpreter.Ast
	item       *item.Item
	term       *term
	osCwd      string
	definition *definitionloader.Definition
	bootSeq    []func() error
}

// todo probably factor-out reloading logic
func (h *Hotshell) Start() error {
	h.initBootSeq()
	for reload, err := h.doStart(); err != nil || reload; reload, err = h.doStart() {
		if err != nil {
			return err
		}
		if err := h.restoreCwd(); err != nil {
			return err
		}
	}
	return nil
}

func (h *Hotshell) initBootSeq() {
	h.bootSeq = []func() error{
		h.loadDefinitionFile,
		h.adjustWorkingDirectory,
		h.interpretDSL,
		h.buildMenu,
	}
}

func (h *Hotshell) doStart() (bool, error) {
	if err := h.executeBootSeq(); err != nil {
		return false, err
	}
	return h.activateAction()
}

func (h *Hotshell) executeBootSeq() error {
	for _, step := range h.bootSeq {
		if err := step(); err != nil {
			return err
		}
	}
	return nil
}

func (h *Hotshell) activateAction() (bool, error) {
	if h.Options.GenerateDemo {
		return false, h.generateDemo()
	} else if h.Options.GenerateDoc {
		return false, h.generateDoc()
	} else {
		return h.startMenu()
	}
}

func (h *Hotshell) generateDoc() error {
	gen := doc.Generator{Item: h.item, Filename: filepath.Base(h.definition.Filename)}
	return gen.Generate()
}

func (h *Hotshell) generateDemo() error {
	gen := demo.Generator{Item: h.item, Filename: filepath.Base(h.definition.Filename)}
	return gen.Generate()
}

func (h *Hotshell) startMenu() (bool, error) {
	if err := h.initTerm(); err != nil {
		return false, err
	}

	defer func() {
		if err := h.term.close(); err != nil {
			fmt.Println(err)
		}
	}()

	return h.startController()
}

func (h *Hotshell) loadDefinitionFile() error {
	var loader = definitionloader.DefinitionLoader{Default: h.Options.Default, File: h.Options.File}
	var err error
	h.definition, err = loader.LoadDefinition()
	return err
}

// todo factor out cwd logic
func (h *Hotshell) chdir() bool {
	return h.Options.Chdir && !h.definition.DefaultMenuLoaded
}

func (h *Hotshell) adjustWorkingDirectory() error {
	if !h.chdir() {
		return nil
	}
	var err error
	h.osCwd, err = os.Getwd()
	if err != nil {
		return err
	}
	return os.Chdir(filepath.Dir(h.definition.Filename))
}

func (h *Hotshell) restoreCwd() error {
	if !h.chdir() {
		return nil
	}
	return os.Chdir(h.osCwd)
}

func (h *Hotshell) buildMenu() error {
	var err error
	h.item, err = (&item.Builder{}).Build(h.ast)
	return err
}

func (h *Hotshell) startController() (bool, error) {
	return (&controller{root: h.item, term: h.term}).start()
}

func (h *Hotshell) initTerm() error {
	var err error
	h.term, err = NewTerm()
	return err
}

func (h *Hotshell) interpretDSL() error {
	inter := interpreter.Interpreter{Filename: h.definition.Filename, Dsl: h.definition.Dsl}
	var err error
	h.ast, err = inter.Interpret()
	return err
}
