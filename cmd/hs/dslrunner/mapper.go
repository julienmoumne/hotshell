package dslrunner

import (
	"errors"
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/mitchellh/mapstructure"
)

type mapper struct{}

func (b *mapper) mapp(ast interface{}) (*item.Item, error) {
	var items []item.Item
	if err := mapstructure.WeakDecode(ast, &items); err != nil {
		return nil, err
	}
	rootItemCount := len(items)
	if rootItemCount == 0 {
		return nil, errors.New("no items found")
	}
	if rootItemCount > 1 {
		return nil, errors.New(fmt.Sprintf("only one top level item is allowed, found %d", rootItemCount))
	}
	it := &items[0]
	if it.IsCmd() {
		return nil, errors.New("top level item must not be a command")
	}
	b.recursiveSetup(it)
	return it, nil
}

func (b *mapper) recursiveSetup(it *item.Item) {
	b.adjustKey(it)
	for _, child := range it.Items {
		child.Parent = it
		b.recursiveSetup(child)
	}
}

func (b *mapper) adjustKey(it *item.Item) {
	if it.Parent == nil {
		return
	}
	if it.Key == "" {
		if it.Cmd != "" || len(it.Items) > 0 {
			it.Key = "key-not-provided"
		}
		return
	}
	if len(it.Key) > 1 {
		it.Key = fmt.Sprintf("invalid-key:%v", it.Key)
		return
	}
	if _, err := it.Parent.GetItem(it.Key[0]); err != nil {
		it.Key = fmt.Sprintf("duplicated-key:%v", it.Key)
	}
}
