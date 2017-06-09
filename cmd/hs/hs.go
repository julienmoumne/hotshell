package main

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/engine"
	"github.com/julienmoumne/hotshell/cmd/hs/versioning"
	"github.com/julienmoumne/hotshell/cmd/options"
	"os"
)

func main() {
	if err := (&hs{}).start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		exit(1)
	}
}

type hs struct {
	options options.Options
}

func (h *hs) start() error {
	if err := h.parseOptions(); err != nil {
		return err
	}
	if h.options.Version {
		return h.printVersion()
	}
	return h.startHotshell()
}

func (h *hs) startHotshell() error {
	return (&engine.Starter{}).Start(h.options)
}

func (h *hs) parseOptions() error {
	var err error
	h.options, err = (&options.OptionParser{}).Parse()
	return err
}

func (h *hs) printVersion() error {
	version, err := versioning.GetVersion()
	if err != nil {
		return err
	}
	fmt.Printf("Hotshell version %s\n", version)
	return nil
}

// var is required to change the definition during tests
var exit = func(code int) {
	os.Exit(code)
}
