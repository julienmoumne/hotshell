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
	activeItem       *item.Item
	lastActivatedCmd string
	term             term.Term
	activeSubprocess bool
	keys             settings.Keys
}

func (c *controller) Start(keySettings settings.Keys, root *item.Item) (bool, error) {
	c.keys = keySettings
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

		// todo test the fact actions are ordered (helps dealing multiple actions having the same key)
		case string(4):
			fmt.Print("\n")
			return false, nil
		case c.keys.Reload:
			c.printKey(c.keys.Reload)
			return true, nil
		case c.keys.Bash:
			c.triggerItem(key, item.BashCmd)
		case c.keys.Back:
			if c.activeItem.Parent != nil {
				c.triggerItem(key, c.activeItem.Parent)
			}
		case c.keys.Repeat:
			c.triggerLastCmd()
		default:
			if selectedItem, err := c.activeItem.GetItem(key); err == nil {
				c.triggerItem(key, selectedItem)
			}
		}
	}
}

func (c *controller) printPrompt() {
	fmt.Printf(
		" %v back, %v bash, %v repeat, %v reload, %v quit",
		formatter.HelpFmt(item.KeyName(c.keys.Back)),
		formatter.HelpFmt(item.KeyName(c.keys.Bash)),
		formatter.HelpFmt(item.KeyName(c.keys.Repeat)),
		formatter.HelpFmt(item.KeyName(c.keys.Reload)),
		formatter.HelpFmt("^d or ^c"),
	)
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print(formatter.KeyActivatedFmt(" ? "))
}

func (c *controller) triggerLastCmd() {
	if c.lastActivatedCmd == "" {
		return
	}

	var it *item.Item
	if c.lastActivatedCmd == c.keys.Bash {
		it = item.BashCmd
	} else {
		it, _ = c.activeItem.GetItem(c.lastActivatedCmd)
	}

	c.triggerItem(c.lastActivatedCmd, it)
}

func (c *controller) printKey(key string) {
	fmt.Print(formatter.KeyActivatedFmt("%s\n\n", item.KeyName(key)))
}

func (c *controller) triggerItem(key string, it *item.Item) {
	c.activeSubprocess = true
	defer func() {
		c.activeSubprocess = false
	}()

	c.printKey(key)

	nextMenu := item.Activate(it)

	menulessCmdActivated := nextMenu == nil
	cmdActivated := menulessCmdActivated || nextMenu == c.activeItem
	if cmdActivated {
		c.lastActivatedCmd = key
		item.Activate(c.activeItem)
	} else {
		c.lastActivatedCmd = ""
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
