package item

import (
	"errors"
	"fmt"
	"github.com/julienmoumne/hotshell/interpreter"
	_ "github.com/robertkrimen/otto/underscore"
)

type Builder struct{}

func (b *Builder) Build(ast []interpreter.Ast) (*Item, error) {
	astLength := len(ast)

	if len(ast) == 0 {
		return nil, errors.New("no items found")
	}
	if len(ast) > 1 {
		return nil, errors.New(fmt.Sprintf("only one top level item is allowed, found %d", astLength))
	}

	item := b.recursiveBuild(ast[0], nil)
	if item.IsCmd() {
		return nil, errors.New("top level item must not be a command")
	}

	return item, nil
}

func (b *Builder) recursiveBuild(config interpreter.Ast, parent *Item) *Item {
	item := b.buildItem(config, parent)
	if parent != nil {
		parent.AddItem(item)
	}
	return item
}

func getKey(config interpreter.Ast, parent *Item) string {
	if parent == nil {
		return ""
	}

	key := config.Key
	if key == "" {
		if config.Cmd != "" || len(config.Items) > 0 {
			return "key-not-provided"
		}
		return ""
	}

	if len(key) > 1 {
		return fmt.Sprintf("invalid-key %v", key)
	}

	if _, duplicated := parent.GetItem(MakeKey(key)); duplicated {
		return fmt.Sprintf("duplicated-key:%v", key)
	}
	return key
}

func (b *Builder) buildItem(config interpreter.Ast, parent *Item) *Item {
	item := NewItem(getKey(config, parent), config.Desc, config.Cmd)

	if item.IsCmd() {
		return item
	}

	for _, subConfig := range config.Items {
		b.recursiveBuild(subConfig, item)
	}

	return item
}
