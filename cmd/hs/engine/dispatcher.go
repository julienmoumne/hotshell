package engine

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/formatter"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/julienmoumne/hotshell/cmd/hs/settings"
	"github.com/julienmoumne/hotshell/cmd/term"
	"os"
	"os/signal"
)

type dispatcher struct {
	keys             settings.Keys
	activeSubprocess bool
	term             *term.Term
}

func newDispatcher(keys settings.Keys) (*dispatcher, error) {
	d := &dispatcher{keys: keys}
	d.initSignals()
	return d, d.initTerm()
}

func (d *dispatcher) dispatchEvent(rawEvent interface{}) {
	switch e := rawEvent.(type) {
	case cmdEvent:
		d.triggerCmd(e)
	case menuEvent:
		d.displayMenu(e)
	case validKeyEvent:
		d.printKey(e)
	}
}

func (d *dispatcher) printKey(e validKeyEvent) {
	fmt.Print(formatter.KeyActivatedFmt("%s\n\n", settings.KeyName(e.key)))
}

func (d *dispatcher) triggerCmd(e cmdEvent) {
	d.activeSubprocess = true
	(&item.CmdActivator{}).Activate(e.item)
	d.activeSubprocess = false
}

func (d *dispatcher) displayMenu(e menuEvent) {
	(&item.MenuPrinter{Out: os.Stdout}).Print(e.item, d.keys)
}

func (d *dispatcher) readUserInput() (string, error) {
	return d.term.ReadUserInput()
}

func (d *dispatcher) initTerm() (err error) {
	d.term, err = term.NewTerm()
	return
}

func (d *dispatcher) cleanup() {
	defer d.closeTerm()
	defer signal.Reset()
}

func (d *dispatcher) closeTerm() {
	if d.term == nil {
		return
	}
	if err := d.term.Close(); err != nil {
		fmt.Println(err)
	}
}

func (d *dispatcher) initSignals() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	go func() {
		for _ = range channel {
			if d.activeSubprocess {
				continue
			}
			d.term.Restore()
			fmt.Print("\n")
			os.Exit(0)
		}
	}()
}
