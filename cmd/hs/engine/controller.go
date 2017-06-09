package engine

// todo unit test me
import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/formatter"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/julienmoumne/hotshell/cmd/term"
	"os"
	"os/signal"
)

type controller struct {
	activeItem       *item.Item
	lastActivatedCmd item.Key
	term             term.Term
	activeSubprocess bool
}

func (c *controller) Start(root *item.Item) (bool, error) {
	if err := c.initTerm(); err != nil {
		return false, err
	}
	defer func() {
		if err := c.term.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	c.initSignals()
	defer c.resetSignals()

	fmt.Print("\n")
	c.activeItem = item.Activate(root)
	return c.mainLoop()
}

func (c *controller) initTerm() (err error) {
	c.term, err = term.NewTerm()
	return
}

func (c *controller) mainLoop() (bool, error) {
	c.printPrompt()
	for {
		key, err := c.term.ReadUserChoice()
		if err != nil {
			return false, err
		}

		switch key {

		case item.EotKey:
			fmt.Print("\n")
			return false, nil
		case item.PreviousMenuKey:
			if c.activeItem.Parent != nil {
				c.triggerItem(key, c.activeItem.Parent)
			}
		case item.BashKey:
			c.triggerItem(key, item.BashCmd)
		case item.RepeatKey:
			c.triggerLastCmd()
		case item.ReloadKey:
			c.printKey(item.ReloadKey)
			return true, nil
		default:
			if selectedItem, err := c.activeItem.GetItem(key); err == nil {
				c.triggerItem(key, selectedItem)
			}
		}
	}
}

func (c *controller) printPrompt() {
	fmt.Print(formatter.KeyActivatedFmt(" ? "))
}

func (c *controller) triggerLastCmd() {
	if c.lastActivatedCmd == (item.Key{}) {
		return
	}

	var it *item.Item
	if c.lastActivatedCmd == item.BashKey {
		it = item.BashCmd
	} else {
		it, _ = c.activeItem.GetItem(c.lastActivatedCmd)
	}

	c.triggerItem(c.lastActivatedCmd, it)
}

func (c *controller) printKey(key item.Key) {
	fmt.Print(formatter.KeyActivatedFmt("%v\n\n", key.String()))
}

func (c *controller) triggerItem(key item.Key, it *item.Item) {
	c.activeSubprocess = true
	defer func() {
		c.activeSubprocess = false
	}()

	c.printKey(key)

	nextMenu := item.Activate(it)

	bashModeActivated := nextMenu == nil
	cmdActivated := nextMenu == c.activeItem
	if bashModeActivated || cmdActivated {
		c.lastActivatedCmd = key
		item.Activate(c.activeItem)
	} else {
		c.lastActivatedCmd = item.Key{}
		c.activeItem = nextMenu
	}

	c.printPrompt()
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
