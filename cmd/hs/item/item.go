package item

import (
	"errors"
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/formatter"
)

var BashCmd = &Item{Key: "", Desc: "bash", Cmd: "bash -l"}

type Item struct {
	Key    string
	Desc   string
	Items  []*Item
	Cmd    string
	Parent *Item
}

func (i *Item) IsCmd() bool {
	return i.Cmd != "" && len(i.Items) == 0
}

func (i *Item) AddItem(item *Item) {
	i.Items = append(i.Items, item)
	item.Parent = i
}

func (i *Item) GetItem(key string) (*Item, error) {
	var found []*Item
	for _, item := range i.Items {
		if item.Key == key {
			found = append(found, item)
		}
	}
	if len(found) != 1 {
		return nil, errors.New(fmt.Sprintf("could not find item for key '%v'", key))
	}
	return found[0], nil
}

func (i *Item) GetDesc() string {
	var desc = i.Desc
	if i.IsCmd() {
		if desc != "" {
			desc += " "
		}
		postfix := formatter.CmdDefFmt(i.Cmd)
		if postfix != "" {
			desc += postfix
		}
	} else {
		if desc == "" {
			desc = "missing-desc"
		}
	}
	return desc
}

func (i *Item) GetInMenuDesc() string {
	var postfix string
	if !i.IsCmd() && i.Parent != nil && (i.Key != "" || len(i.Items) > 0) {
		postfix = formatter.CmdDefFmt(" >")
	}

	if i.Key == "" {
		return fmt.Sprintf("%s%s", i.GetDesc(), postfix)
	}
	return fmt.Sprintf("%s %s%s", formatter.KeyHintFmt("%s", i.Key), i.GetDesc(), postfix)
}
