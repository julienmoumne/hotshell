package item_test

import (
	"errors"
	"github.com/julienmoumne/hotshell/cmd/hs/interpreter"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	. "gopkg.in/check.v1"
	"testing"
)

func TestBuilder(t *testing.T) { TestingT(t) }

type BuilderTestSuite struct{}

var _ = Suite(&BuilderTestSuite{})

var builderTests = []struct {
	in  []interpreter.Ast
	out *item.Item
	err error
}{
	// Failures
	{
		[]interpreter.Ast{},
		nil,
		errors.New("no items found"),
	},
	{
		[]interpreter.Ast{{}, {}},
		nil,
		errors.New("only one top level item is allowed, found 2"),
	},
	{
		[]interpreter.Ast{{Cmd: "echo 'test'"}},
		nil,
		errors.New("top level item must not be a command"),
	},

	// Missing desc, nothing special happens
	{
		[]interpreter.Ast{{Items: []interpreter.Ast{{Key: "k", Cmd: "missing-desc"}}}},
		&item.Item{
			Items: []*item.Item{
				&item.Item{
					Key:   "k",
					Cmd:   "missing-desc",
					Items: []*item.Item{},
				},
			},
		},
		nil,
	},

	// Empty menu, nothing special happens
	{
		[]interpreter.Ast{{Desc: "top-level-no-items"}},
		&item.Item{
			Desc:  "top-level-no-items",
			Items: []*item.Item{},
		},
		nil,
	},

	// Invalid keys
	{
		[]interpreter.Ast{{Desc: "invalid-keys",
			Items: []interpreter.Ast{
				interpreter.Ast{Key: "", Desc: "key-not-provided-cmd", Cmd: "key-not-provided-cmd"},
				interpreter.Ast{Key: "d", Desc: "duplicated-key", Cmd: "duplicated-key"},
				interpreter.Ast{Key: "", Desc: "key-not-provided-menu", Items: []interpreter.Ast{
					interpreter.Ast{Desc: "key-not-provided-empty-menu"},
				}},
				interpreter.Ast{Key: "d", Desc: "duplicated-key", Cmd: "duplicated-key"},
				interpreter.Ast{Key: "too-long", Desc: "too-long", Cmd: "too-long"},
			},
		}},
		&item.Item{
			Desc: "invalid-keys",
			Items: []*item.Item{
				&item.Item{Key: "key-not-provided", Desc: "key-not-provided-cmd", Cmd: "key-not-provided-cmd",
					Items: []*item.Item{},
				},
				&item.Item{Key: "d", Desc: "duplicated-key", Cmd: "duplicated-key",
					Items: []*item.Item{},
				},
				&item.Item{Key: "key-not-provided", Desc: "key-not-provided-menu",
					Items: []*item.Item{&item.Item{
						Desc:  "key-not-provided-empty-menu",
						Items: []*item.Item{},
					}},
				},
				&item.Item{Key: "duplicated-key:d", Desc: "duplicated-key", Cmd: "duplicated-key",
					Items: []*item.Item{},
				},
				&item.Item{Key: "invalid-key too-long", Desc: "too-long", Cmd: "too-long",
					Items: []*item.Item{},
				},
			},
		},
		nil,
	},

	// Doubly nested menu
	{
		[]interpreter.Ast{{Key: "t", Desc: "test",
			Items: []interpreter.Ast{
				interpreter.Ast{Key: "f", Desc: "first cmd", Cmd: "echo 'first cmd'"},
				interpreter.Ast{Key: "s", Desc: "second cmd", Cmd: "echo 'second cmd'"},
				interpreter.Ast{Key: "m", Desc: "submenu",
					Items: []interpreter.Ast{
						interpreter.Ast{Key: "s", Desc: "submenu cmd", Cmd: "echo 'submenu cmd'"},
					},
				},
			},
		}},
		&item.Item{
			Desc: "test",
			Items: []*item.Item{
				&item.Item{Key: "f", Desc: "first cmd", Cmd: "echo 'first cmd'",
					Items: []*item.Item{},
				},
				&item.Item{Key: "s", Desc: "second cmd", Cmd: "echo 'second cmd'",
					Items: []*item.Item{},
				},
				&item.Item{Key: "m", Desc: "submenu",
					Items: []*item.Item{
						&item.Item{
							Key:   "s",
							Desc:  "submenu cmd",
							Cmd:   "echo 'submenu cmd'",
							Items: []*item.Item{},
						},
					},
				},
			},
		},
		nil,
	},
}

func (s *BuilderTestSuite) TestBuilder(c *C) {
	for _, tt := range builderTests {
		if tt.out != nil {
			adjustParentLinks(tt.out, nil)
		}
		actualOut, err := (&item.Builder{}).Build(tt.in)
		c.Check(actualOut, DeepEquals, tt.out)
		c.Check(err, DeepEquals, tt.err)
	}
}

func adjustParentLinks(item *item.Item, parent *item.Item) {
	item.Parent = parent
	for _, it := range item.Items {
		adjustParentLinks(it, item)
	}
}
