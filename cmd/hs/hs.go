package main

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/engine"
	"github.com/julienmoumne/hotshell/cmd/hs/versioning"
	"github.com/julienmoumne/hotshell/cmd/options"
	"github.com/robertkrimen/otto"
	"os"
)

func main() {
	handleError((&hs{}).start())
}

type hs struct {
	options *options.Options
}

func (m *hs) start() error {
	if err := m.parseOptions(); err != nil {
		return err
	}
	if m.options.Version {
		return m.printVersion()
	}
	return m.startHotshell()
}

func (m *hs) startHotshell() error {
	return (&engine.Starter{Options: m.options}).Start()
}

func (m *hs) parseOptions() error {
	var err error
	m.options, err = (&options.OptionParser{}).Parse()
	return err
}

func (m *hs) printVersion() error {
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
