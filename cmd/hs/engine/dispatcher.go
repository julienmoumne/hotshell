package engine

// todo unit test
import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/formatter"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/julienmoumne/hotshell/cmd/hs/settings"
	"github.com/julienmoumne/hotshell/cmd/term"
	"os"
	"os/signal"
)

type Dispatcher interface {
	DispatchEvent(rawEvent interface{})
	ReadUserInput() (string, error)
	Cleanup()
}

func DefaultDispatcher(keys settings.Keys)(Dispatcher, error) {
	d := &sysDispatcher{keys: keys}
	d.initSignals()
	return d, d.initTerm()
}

type sysDispatcher struct {
	keys             settings.Keys
	activeSubprocess bool
	term             *term.Term
}

func (d *sysDispatcher) DispatchEvent(rawEvent interface{}) {
	switch e := rawEvent.(type) {
	case CmdEvent:
		d.triggerCmd(e)
	case MenuEvent:
		d.displayMenu(e)
	case ValidKeyEvent:
		d.printKey(e)
	}
}

func (d *sysDispatcher) printKey(e ValidKeyEvent) {
	fmt.Print(formatter.KeyActivatedFmt("%s\n\n", settings.KeyName(e.Key)))
}

func (d *sysDispatcher) triggerCmd(e CmdEvent) {
	d.activeSubprocess = true
	(&item.CmdActivator{}).Activate(e.Item)
	d.activeSubprocess = false
}

func (d *sysDispatcher) displayMenu(e MenuEvent) {
	(&item.MenuPrinter{Out: os.Stdout}).Print(e.Item, d.keys)
}

func (d *sysDispatcher) ReadUserInput() (string, error) {
	return d.term.ReadUserInput()
}

func (d *sysDispatcher) initTerm() (err error) {
	d.term, err = term.NewTerm()
	return
}

func (d *sysDispatcher) Cleanup() {
	defer d.closeTerm()
	defer signal.Reset()
}

func (d *sysDispatcher) closeTerm() {
	if d.term == nil {
		return
	}
	if err := d.term.Close(); err != nil {
		fmt.Println(err)
	}
}

func (d *sysDispatcher) initSignals() {
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
