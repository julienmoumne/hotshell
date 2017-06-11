package engine

// todo unit test me
import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/formatter"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/julienmoumne/hotshell/cmd/hs/settings"
	"github.com/julienmoumne/hotshell/cmd/term"
	"os"
	"os/signal"
)

type controller struct {
	activeMenu        *item.Item
	lastActivatedItem *item.Item
	term              term.Term
	activeSubprocess  bool
	keys              settings.Keys
}

func (c *controller) Start(keySettings settings.Keys, root *item.Item, term term.Term) (bool, error) {
	c.term = term
	c.keys = keySettings
	c.activeMenu = root
	c.lastActivatedItem = root
	c.initSignals()
	defer c.resetSignals()
	return c.mainLoop()
}

func (c *controller) mainLoop() (bool, error) {
	c.displayMenu(c.activeMenu)
	for {
		key, err := c.term.ReadUserChoice()
		if err != nil {
			return false, err
		}
		switch key {

		// todo test the fact actions are ordered (helps dealing multiple actions having the same key)
		case string(4):
			fmt.Print("\n")
			return false, nil
		case c.keys.Reload:
			c.printKey(key)
			return true, nil
		case c.keys.Bash:
			c.triggerItem(key, item.BashCmd(c.keys.Bash))
		case c.keys.Back:
			if c.activeMenu.Parent != nil {
				c.triggerItem(key, c.activeMenu.Parent)
			} else {
				c.triggerItem(key, c.activeMenu)
			}
		case c.keys.Repeat:
			c.triggerItem(key, c.lastActivatedItem)
		default:
			if it, err := c.activeMenu.GetItem(key); err == nil {
				c.triggerItem(key, it)
			}
		}
	}
}

func (c *controller) triggerItem(key string, it *item.Item) {
	c.printKey(key)
	c.lastActivatedItem = it
	if it.IsCmd() {
		c.triggerCmd(it)
	} else {
		c.activeMenu = it
	}
	c.displayMenu(c.activeMenu)
}

func (c *controller) triggerCmd(cmd *item.Item) {
	c.activeSubprocess = true
	(&item.CmdActivator{}).Activate(cmd)
	c.activeSubprocess = false
}

func (c *controller) displayMenu(menu *item.Item) {
	(&item.MenuPrinter{Out: os.Stdout}).Print(c.activeMenu, c.keys)
}

func (c *controller) printKey(key string) {
	fmt.Print(formatter.KeyActivatedFmt("%s\n\n", settings.KeyName(key)))
}

func (c *controller) initSignals() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	go func() {
		for _ = range channel {
			if c.activeSubprocess {
				continue
			}
			c.term.Restore()
			fmt.Print("\n")
			os.Exit(0)
		}
	}()
}

func (c *controller) resetSignals() {
	signal.Reset()
}
