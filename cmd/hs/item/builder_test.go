package item_test

import (
	"errors"
	. "github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"testing"
)

var builderTests = []struct {
	in  interface{}
	out *Item
	err error
}{
	// Empty values
	{
		in:  nil,
		err: errors.New("no items found"),
	},
	{
		in:  []map[string]interface{}{},
		err: errors.New("no items found"),
	},
	{
		in:  []map[string]interface{}{{}},
		out: &Item{},
	},
	{
		in:  []map[string]interface{}{{}, {}},
		err: errors.New("only one top level item is allowed, found 2"),
	},
	{
		in:  []map[string]interface{}{{"ignoredKey": []string{"test"}}},
		out: &Item{},
	},
	// Various Otto types for integers
	{
		in:  []map[string]interface{}{{"desc": 1, "key": int64(1)}},
		out: &Item{Key: "1", Desc: "1"},
	},
	{
		in:  []map[string]interface{}{{"desc": 1.0, "key": float64(1)}},
		out: &Item{Key: "1", Desc: "1"},
	},
	// Invalid root item
	{
		in:  []map[string]interface{}{{"cmd": "echo 'test'"}},
		err: errors.New("top level item must not be a command"),
	},
	// Missing desc, nothing special happens
	{
		in:  []map[string]interface{}{{"items": []map[string]interface{}{{"key": "k", "cmd": "missing-desc"}}}},
		out: &Item{Items: []*Item{{Key: "k", Cmd: "missing-desc"}}},
	},
	// Empty menu, nothing special happens
	{
		in:  []map[string]interface{}{{"desc": "top-level-no-items"}},
		out: &Item{Desc: "top-level-no-items"},
	},
	// Invalid keys
	{
		in: []map[string]interface{}{{
			"desc": "invalid-keys",
			"items": []map[string]interface{}{
				{"desc": "key-not-provided-cmd", "cmd": "key-not-provided-cmd"},
				{"key": "d", "desc": "duplicated-key", "cmd": "duplicated-key"},
				{"desc": "key-not-provided-menu", "items": []map[string]interface{}{{
					"desc": "key-not-provided-empty-menu",
				}}},
				{"key": "d", "desc": "duplicated-key", "cmd": "duplicated-key"},
				{"key": "too-long", "desc": "too-long", "cmd": "too-long"},
			}}},
		out: &Item{
			Desc: "invalid-keys",
			Items: []*Item{
				{Key: "key-not-provided", Desc: "key-not-provided-cmd", Cmd: "key-not-provided-cmd"},
				{Key: "duplicated-key:d", Desc: "duplicated-key", Cmd: "duplicated-key"},
				{Key: "key-not-provided", Desc: "key-not-provided-menu",
					Items: []*Item{{Desc: "key-not-provided-empty-menu"}},
				},
				{Key: "duplicated-key:d", Desc: "duplicated-key", Cmd: "duplicated-key"},
				{Key: "invalid-key too-long", Desc: "too-long", Cmd: "too-long"},
			},
		},
	},
	// Type errors
	{
		in:  []interface{}{"test", 2},
		err: &mapstructure.Error{Errors: []string{"'[0]' expected a map, got 'string'", "'[1]' expected a map, got 'int'"}},
	},
	{
		in:  []map[string]interface{}{{"key": []string{}}},
		err: &mapstructure.Error{Errors: []string{"'[0].Key' expected type 'string', got unconvertible type '[]string'"}},
	},
	{
		in:  []map[string]interface{}{{"desc": []string{}}},
		err: &mapstructure.Error{Errors: []string{"'[0].Desc' expected type 'string', got unconvertible type '[]string'"}},
	},
	{
		in:  []map[string]interface{}{{"cmd": []string{}}},
		err: &mapstructure.Error{Errors: []string{"'[0].Cmd' expected type 'string', got unconvertible type '[]string'"}},
	},
	{
		in:  []map[string]interface{}{{"items": []string{"test"}}},
		err: &mapstructure.Error{Errors: []string{"'[0].Items[0]' expected a map, got 'string'"}},
	},
	// Doubly nested menu
	{
		in: []map[string]interface{}{{
			"key":  "t",
			"desc": "test",
			"items": []map[string]interface{}{
				{
					"key":  "f",
					"desc": "first cmd",
					"cmd":  "echo 'first cmd'",
				},
				{
					"key":  "s",
					"desc": "second cmd",
					"cmd":  "echo 'second cmd'",
				},
				{
					"key":  "m",
					"desc": "submenu",
					"items": []map[string]interface{}{
						{
							"key":  "s",
							"desc": "submenu cmd",
							"cmd":  "echo 'submenu cmd'",
						},
					},
				},
			},
		}},
		out: &Item{
			Key:  "t",
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
