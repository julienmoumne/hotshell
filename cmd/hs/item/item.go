package item

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/formatter"
	"strings"
)

var BashCmd = NewItem("", "bash", "bash -l")

type Item struct {
	Key    string
	Desc   string
	Items  []*Item
	Cmd    string
	Parent *Item
}

func NewItem(key string, desc string, cmd string) *Item {
	item := Item{}
	item.Key = key
	item.Desc = desc
	item.Cmd = strings.TrimSpace(cmd)
	item.Items = make([]*Item, 0)
	return &item
}

func (i *Item) IsCmd() bool {
	return i.Cmd != "" && len(i.Items) == 0
}

func (i *Item) AddItem(item *Item) {
	i.Items = append(i.Items, item)
	item.Parent = i
}

func (i *Item) GetItem(key Key) (*Item, bool) {
	for _, item := range i.Items {
		if MakeKey(item.Key) == key {
			return item, true
		}
	}
	return nil, false
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
