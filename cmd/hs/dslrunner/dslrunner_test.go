package dslrunner_test

import (
	"github.com/julienmoumne/hotshell/cmd/hs/dslrunner"
	. "github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	in  string
	out *Item
	err string
}{
	// empty values
	{
		in:  ``,
		err: "no items found",
	},
	{
		in:  `require('hotshell').item({})`,
		out: &Item{},
	},
	{
		in: `
		var item = require('hotshell').item
		item({})
		item({})
		`,
		err: "only one top level item is allowed, found 2",
	},
	{
		in:  `require('hotshell').item({ignoredKey: ['test']})`,
		out: &Item{},
	},
	// numerical values
	{
		in:  `require('hotshell').item({key: '1', desc: 1})`,
		out: &Item{Key: "1", Desc: "1"},
	},
	{
		in:  `require('hotshell').item({desc: 1.0})`,
		out: &Item{Desc: "1"},
	},
	// invalid root item
	{
		in:  `require('hotshell').item({cmd: "echo 'test'"})`,
		err: "top level item must not be a command",
	},
	// missing desc
	{
		in: `
		var item = require('hotshell').item
		item({}, function() {
			item({key: 'k', cmd: 'missing-desc'})
		})`,
		out: &Item{Items: []*Item{{Key: "k", Cmd: "missing-desc"}}},
	},
	// empty menu
	{
		in:  `require('hotshell').item({desc: 'top-level-no-items'})`,
		out: &Item{Desc: "top-level-no-items"},
	},
	// invalid keys
	{
		in: `
		var item = require('hotshell').item
		item({desc: 'invalid-keys'}, function(){
			item({desc: 'key-not-provided-cmd', cmd: 'key-not-provided-cmd'})
			item({key: 'd', desc: 'duplicated-key', cmd: 'duplicated-key'})
			item({desc: 'key-not-provided-menu'}, function() {
				item({desc: 'key-not-provided-empty-menu'})
			})
			item({desc: 'key-not-provided-empty-menu'})
			item({key: 'd', desc: 'duplicated-key', cmd: 'duplicated-key'})
			item({key: 'too-long', desc: 'too-long', cmd: 'too-long'})
		})`,
		out: &Item{
			Desc: "invalid-keys",
			Items: []*Item{
				{Key: "key-not-provided", Desc: "key-not-provided-cmd", Cmd: "key-not-provided-cmd"},
				{Key: "duplicated-key:d", Desc: "duplicated-key", Cmd: "duplicated-key"},
				{Key: "key-not-provided", Desc: "key-not-provided-menu",
					Items: []*Item{{Desc: "key-not-provided-empty-menu"}},
				},
				{Desc: "key-not-provided-empty-menu"},
				{Key: "duplicated-key:d", Desc: "duplicated-key", Cmd: "duplicated-key"},
				{Key: "invalid-key too-long", Desc: "too-long", Cmd: "too-long"},
			},
		},
	},
	// JS runtime errors
	{
		in:  `item()`,
		err: "ReferenceError: 'item' is not defined\n    at <anonymous>:2:1\n    at <unknown>",
	},
	{
		in: `
		var item = require('hotshell').item
		item({desc: 'ref err in closure'}, function () {
		    undefinedMethod()
		}) `,
		out: &Item{Desc: "ref err in closure [Exception caught, ReferenceError: 'undefinedMethod' is not defined]"},
	},
	{
		in:  `invalidStatement{}`,
		err: "(anonymous): Line 2:17 Unexpected token { (and 2 more errors)",
	},
	{
		in: `
		var item = require('hotshell').item
		item({}, function () {
		    item({desc: 'nested closure'}, function () {
			item({key: 's', desc: 'skipped'})
			throw new Error('Runtime Error')
		    })
		    item({key: 'n', desc: 'not skipped', cmd: 'echo not skipped'})
		})
		`,
		out: &Item{Items: []*Item{
			{Desc: "nested closure [Exception caught, Error: Runtime Error]"},
			{Key: "n", Desc: "not skipped", Cmd: "echo not skipped"},
		}},
	},
	{
		in: `
		var item = require('hotshell').item
		var exec = require('hotshell').exec
		item({desc: exec('eco "1"')})
		`,
		err: "\"/bin/bash -c 'eco \"1\"'\" failed with exit status 127 \"bash: eco: command not found\"\n    at <unknown>\n    at <anonymous>:5:15\n    at <unknown>",
	},
	{
		in: `
		var item = require('hotshell').item
		var exec = require('hotshell').exec
		item({desc: 'exec error'}, function() {
			exec('eco "1"')
		})
		`,
		out: &Item{Desc: "exec error [Exception caught, Error: \"/bin/bash -c 'eco \"1\"'\" failed with exit status 127 \"bash: eco: command not found\"]"},
	},
	// type errors
	{
		in:  `require('hotshell').items = ['test', 2]`,
		err: "2 error(s) decoding:\n\n* '[0]' expected a map, got 'string'\n* '[1]' expected a map, got 'int64'",
	},
	{
		in:  `require('hotshell').item({key: []})`,
		err: "1 error(s) decoding:\n\n* '[0].Key' expected type 'string', got unconvertible type '[]interface {}'",
	},
	{
		in:  `require('hotshell').item({desc: []})`,
		err: "1 error(s) decoding:\n\n* '[0].Desc' expected type 'string', got unconvertible type '[]interface {}'",
	},
	{
		in:  `require('hotshell').item({cmd: []})`,
		err: "1 error(s) decoding:\n\n* '[0].Cmd' expected type 'string', got unconvertible type '[]interface {}'",
	},
	{
		in:  `require('hotshell').items = [{items: ["test"]}]`,
		err: "1 error(s) decoding:\n\n* '[0].Items[0]' expected a map, got 'string'",
	},
	// doubly nested menu
	{
		in: `
		var item = require('hotshell').item
		item({key: 't', desc: 'test'}, function() {
			item({key: 'f', desc: 'first cmd', cmd: "echo 'first cmd'"})
			item({key: 's', desc: 'second cmd', cmd: "echo 'second cmd'"})
			item({key: 'm', desc: 'submenu'}, function() {
				item({key: 's', desc: 'submenu cmd', cmd: "echo 'submenu cmd'"})
			})
		})`,
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
	// utility functions
	{
		in: `
		var item = require('hotshell').item
		var _ = require('underscore')
		item({desc: _.min([2, 1])})
		`,
		out: &Item{Desc: "1"},
	},
	{
		in: `
		var item = require('hotshell').item
		var exec = require('hotshell').exec
		item({key: exec('echo "1"')})
		`,
		out: &Item{Key: "1"},
	},
	{
		in: `
		var item = require('hotshell').item
		eval("item({desc: 'evaled menu'})")
		`,
		out: &Item{Desc: "evaled menu"},
	},
	// submodule
	{
		in: `
		var item = require('hotshell').item
		var submenu = require('./submodule_test.js')
		item({}, function () {
		    submenu()
		})
		`,
		out: &Item{Items: []*Item{{Key: "e", Cmd: "echo submodule"}}},
	},
}

func TestDslRunner(t *testing.T) {
	for _, tt := range tests {
		adjustParentLinks(tt.out, nil)
		actualOut, err := (&dslrunner.DslRunner{}).Run(tt.in)
		a := assert.New(t)
		if tt.err != "" {
			a.Equal(tt.err, err.Error())
		} else {
			a.Nil(err)
			a.Equal(tt.out, actualOut)
		}
	}
}

func adjustParentLinks(item *Item, parent *Item) {
	if item == nil {
		return
	}
	item.Parent = parent
	for _, it := range item.Items {
		adjustParentLinks(it, item)
	}
}
