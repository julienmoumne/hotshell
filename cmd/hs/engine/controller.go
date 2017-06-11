package engine

// todo unit test me
import (
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/julienmoumne/hotshell/cmd/hs/settings"
	"github.com/julienmoumne/hotshell/cmd/term"
)

type controller struct {
	activeMenu        *item.Item
	lastActivatedItem *item.Item
	term              term.Term
	keys              settings.Keys
	dispatcher        *dispatcher
}

type validKeyEvent struct {
	key string
}
type itemEvent struct {
	item *item.Item
}
type cmdEvent struct {
	itemEvent
}
type menuEvent struct {
	itemEvent
}

func (c *controller) Start(keySettings settings.Keys, root *item.Item, dispatcher *dispatcher) (bool, error) {
	c.dispatcher = dispatcher
	c.keys = keySettings
	c.activeMenu = root
	c.lastActivatedItem = root
	return c.mainLoop()
}

func (c *controller) mainLoop() (bool, error) {
	c.sendMenuEvent(c.activeMenu)
	for {
		key, err := c.dispatcher.readUserInput()
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

func (c *controller) triggerItem(key string, it *item.Item) {
	c.sendKeyEvent(key)
	c.lastActivatedItem = it
	if it.IsCmd() {
		c.sendCmdEvent(it)
	} else {
		c.activeMenu = it
	}
	c.sendMenuEvent(c.activeMenu)
}

func (c *controller) sendCmdEvent(cmd *item.Item) {
	c.dispatcher.dispatchEvent(cmdEvent{itemEvent{item: cmd}})
}

func (c *controller) sendMenuEvent(menu *item.Item) {
	c.dispatcher.dispatchEvent(menuEvent{itemEvent{item: menu}})
}

func (c *controller) sendKeyEvent(key string) {
	c.dispatcher.dispatchEvent(validKeyEvent{key: key})
}
