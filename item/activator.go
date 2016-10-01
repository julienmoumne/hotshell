package item

import (
  "github.com/julienmoumne/hotshell/interpreter"
  "os"
)

type activator interface {
	Activate(conf interpreter.Conf, item *Item) *Item
}

func Activate(conf interpreter.Conf, item *Item) *Item {
	var activator activator
	if item.IsCmd() {
		activator = &CmdActivator{Out: os.Stdout}
	} else {
		activator = &MenuActivator{Out: os.Stdout}
	}
	return activator.Activate( conf, item )
}
