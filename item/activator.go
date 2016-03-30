package item

import "os"

type activator interface {
	Activate(item *Item) *Item
}

func Activate(item *Item) *Item {
	var activator activator
	if item.IsCmd() {
		activator = &CmdActivator{Out: os.Stdout}
	} else {
		activator = &MenuActivator{Out: os.Stdout}
	}
	return activator.Activate(item)
}
