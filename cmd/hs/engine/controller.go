package engine

import (
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/julienmoumne/hotshell/cmd/hs/settings"
	"github.com/julienmoumne/hotshell/cmd/term"
)

type Controller struct {
	activeMenu        *item.Item
	lastActivatedItem *item.Item
	term              term.Term
	keys              settings.Keys
	dispatcher        Dispatcher
}

type ValidKeyEvent struct {
	Key string
}
type CmdEvent struct {
	Item *item.Item
}
type MenuEvent struct {
	Item *item.Item
}

func (c *Controller) Start(keySettings settings.Keys, root *item.Item, dispatcher Dispatcher) (bool, error) {
	c.dispatcher = dispatcher
	c.keys = keySettings
	c.activeMenu = root
	c.lastActivatedItem = root
	return c.mainLoop()
}

func (c *Controller) mainLoop() (bool, error) {
	c.sendMenuEvent(c.activeMenu)
	for {
		key, err := c.dispatcher.ReadUserInput()
		if err != nil {
			return false, err
		}
		switch key {

		// todo test the fact actions are ordered (helps dealing multiple actions having the same key)
		case string(4):
			return false, nil
		case c.keys.Reload:
			c.sendKeyEvent(key)
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

func (c *Controller) triggerItem(key string, it *item.Item) {
	c.sendKeyEvent(key)
	c.lastActivatedItem = it
	if it.IsCmd() {
		c.sendCmdEvent(it)
	} else {
		c.activeMenu = it
	}
	c.sendMenuEvent(c.activeMenu)
}

func (c *Controller) sendCmdEvent(cmd *item.Item) {
	c.dispatcher.DispatchEvent(CmdEvent{Item: cmd})
}

func (c *Controller) sendMenuEvent(menu *item.Item) {
	c.dispatcher.DispatchEvent(MenuEvent{Item: menu})
}

func (c *Controller) sendKeyEvent(key string) {
	c.dispatcher.DispatchEvent(ValidKeyEvent{Key: key})
}
