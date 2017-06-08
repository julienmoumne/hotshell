package item_test

import (
	"errors"
	. "github.com/julienmoumne/hotshell/cmd/hs/interpreter"
	. "github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/stretchr/testify/assert"
	"testing"
)

var builderTests = []struct {
	in  []Ast
	out *Item
	err error
}{
	// Failures
	{
		[]Ast{},
		nil,
		errors.New("no items found"),
	},
	{
		[]Ast{{}, {}},
		nil,
		errors.New("only one top level item is allowed, found 2"),
	},
	{
		[]Ast{{Cmd: "echo 'test'"}},
		nil,
		errors.New("top level item must not be a command"),
	},

	// Missing desc, nothing special happens
	{
		[]Ast{{Items: []Ast{{Key: "k", Cmd: "missing-desc"}}}},
		&Item{Items: []*Item{{Key: "k", Cmd: "missing-desc"}}},
		nil,
	},

	// Empty menu, nothing special happens
	{
		[]Ast{{Desc: "top-level-no-items"}},
		&Item{Desc: "top-level-no-items"},
		nil,
	},

	// Invalid keys
	{
		[]Ast{{Desc: "invalid-keys",
			Items: []Ast{
				{Key: "", Desc: "key-not-provided-cmd", Cmd: "key-not-provided-cmd"},
				{Key: "d", Desc: "duplicated-key", Cmd: "duplicated-key"},
				{Key: "", Desc: "key-not-provided-menu", Items: []Ast{
					{Desc: "key-not-provided-empty-menu"},
				}},
				{Key: "d", Desc: "duplicated-key", Cmd: "duplicated-key"},
				{Key: "too-long", Desc: "too-long", Cmd: "too-long"},
			},
		}},
		&Item{
			Desc: "invalid-keys",
			Items: []*Item{
				{Key: "key-not-provided", Desc: "key-not-provided-cmd", Cmd: "key-not-provided-cmd"},
				{Key: "d", Desc: "duplicated-key", Cmd: "duplicated-key"},
				{Key: "key-not-provided", Desc: "key-not-provided-menu",
					Items: []*Item{{Desc: "key-not-provided-empty-menu"}},
				},
				{Key: "duplicated-key:d", Desc: "duplicated-key", Cmd: "duplicated-key"},
				{Key: "invalid-key too-long", Desc: "too-long", Cmd: "too-long"},
			},
		},
		nil,
	},

	// Doubly nested menu
	{
		[]Ast{{Key: "t", Desc: "test",
			Items: []Ast{
				{Key: "f", Desc: "first cmd", Cmd: "echo 'first cmd'"},
				{Key: "s", Desc: "second cmd", Cmd: "echo 'second cmd'"},
				{Key: "m", Desc: "submenu",
					Items: []Ast{
						{Key: "s", Desc: "submenu cmd", Cmd: "echo 'submenu cmd'"},
					},
				},
			},
		}},
		&Item{
			Desc: "test",
			Items: []*Item{
				{Key: "f", Desc: "first cmd", Cmd: "echo 'first cmd'"},
				{Key: "s", Desc: "second cmd", Cmd: "echo 'second cmd'"},
				{Key: "m", Desc: "submenu",
					Items: []*Item{
						{
							Key:  "s",
							Desc: "submenu cmd",
							Cmd:  "echo 'submenu cmd'",
						},
					},
				},
			},
		},
		nil,
	},
}

func TestBuilder(t *testing.T) {
	for _, tt := range builderTests {
		if tt.out != nil {
			adjustParentLinks(tt.out, nil)
		}
		actualOut, err := (&Builder{}).Build(tt.in)
		a := assert.New(t)
		a.Equal(tt.out, actualOut)
		a.Equal(tt.err, err)
	}
}

func adjustParentLinks(item *Item, parent *Item) {
	item.Parent = parent
	for _, it := range item.Items {
		adjustParentLinks(it, item)
	}
}
