package main

import (
	"fmt"
	"github.com/julienmoumne/hotshell"
	"github.com/julienmoumne/hotshell/cmd/optionparser"
	"github.com/julienmoumne/hotshell/versioning"
	"github.com/robertkrimen/otto"
	"os"
)

func main() {
	handleError((&Main{}).boot())
}

type Main struct {
	options *hotshell.Options
}

func (m *Main) boot() error {
	if err := m.parseOptions(); err != nil {
		return err
	}
	if m.options.Version {
		return m.printVersion()
	}
	return m.startHotshell()
}

func (m *Main) startHotshell() error {
	return (&hotshell.Hotshell{Options: m.options}).Start()
}

func (m *Main) parseOptions() error {
	var err error
	m.options, err = (&optionparser.OptionParser{}).Parse()
	return err
}

func (m *Main) printVersion() error {
	version, err := versioning.GetVersion()
	if err != nil {
		return err
	}
	fmt.Printf("Hotshell version %s\n", version)
	return nil
}

func handleError(err error) {
	if err == nil {
		return
	}
	switch err := err.(type) {
	case *otto.Error:
		fmt.Fprint(os.Stderr, err.String())
	default:
		fmt.Fprintln(os.Stderr, err)
	}
	exit(1)
}

// var is required to change the definition during tests
var exit = func(code int) {
	os.Exit(code)
}
