package dslrunner_test

import (
	. "fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/dslrunner"
	. "github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var tests = []testCase{
	// empty values
	{
		in:  ``,
		err: errMsg("no items found"),
	},
	{
		in:  `require('hotshell').item({})`,
		out: &Item{Wd: "./"},
	},
	{
		in: `
		var item = require('hotshell').item
		item({})
		item({})
		`,
		err: errMsg("only one top level item is allowed, found 2"),
	},
	{
		in:  `require('hotshell').item({ignoredKey: ['test']})`,
		out: &Item{Wd: "./"},
	},
	// numerical values
	{
		in:  `require('hotshell').item({key: '1', desc: 1})`,
		out: &Item{Key: "1", Desc: "1", Wd: "./"},
	},
	{
		in:  `require('hotshell').item({desc: 1.0})`,
		out: &Item{Desc: "1", Wd: "./"},
	},
	// invalid root item
	{
		in:  `require('hotshell').item({cmd: "echo 'test'"})`,
		err: errMsg("top level item must not be a command"),
	},
	// missing desc
	{
		in: `
		var item = require('hotshell').item
		item({}, function() {
			item({key: 'k', cmd: 'missing-desc'})
		})`,
		out: &Item{Items: []*Item{{Key: "k", Cmd: "missing-desc", Wd: "./"}}, Wd: "./"},
	},
	// empty menu
	{
		in:  `require('hotshell').item({desc: 'top-level-no-items'})`,
		out: &Item{Desc: "top-level-no-items", Wd: "./"},
	},
	// invalid keys
	{
		in: `
		var item = require('hotshell').item
		item({desc: 'invalid-keys'}, function(){
			item({Cmd: 'command-without-key'})
			item({key: 'É'})
			item({}, function() {
				item({})
			})
			item({})
			item({key: 'É'})
			item({key: 'too-long'})
		})`,
		out: &Item{
			Desc: "invalid-keys",
			Items: []*Item{
				{Key: "key-not-provided", Cmd: "command-without-key", Wd: "./"},
				{Key: "duplicated-key:É", Wd: "./"},
				{Key: "key-not-provided", Items: []*Item{{Wd: "./"}}, Wd: "./"},
				{Wd: "./"},
				{Key: "É", Wd: "./"},
				{Key: "invalid-key:too-long", Wd: "./"},
			},
			Wd: "./",
		},
	},
	// JS runtime errors
	{
		in:  `item()`,
		err: errMsg("ReferenceError: 'item' is not defined\n    at <anonymous>:2:1\n    at <unknown>"),
	},
	{
		in: `
		var item = require('hotshell').item
		item({desc: 'ref err in closure'}, function () {
		    undefinedMethod()
		}) `,
		out: &Item{Desc: "ref err in closure [Exception caught, ReferenceError: 'undefinedMethod' is not defined]", Wd: "./"},
	},
	{
		in:  `invalidStatement{}`,
		err: errMsg("(anonymous): Line 2:17 Unexpected token { (and 2 more errors)"),
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
			{Desc: "nested closure [Exception caught, Error: Runtime Error]", Wd: "./"},
			{Key: "n", Desc: "not skipped", Cmd: "echo not skipped", Wd: "./"},
		}, Wd: "./"},
	},
	{
		in: `
		var item = require('hotshell').item
		var exec = require('hotshell').exec
		item({desc: exec('eco "1"')})
		`,
		err: errMsg("\"/bin/bash -c 'eco \"1\"'\" failed with exit status 127 \"bash: eco: command not found\"\n    at <unknown>\n    at <anonymous>:5:15\n    at <unknown>"),
	},
	{
		in: `
		var item = require('hotshell').item
		var exec = require('hotshell').exec
		item({desc: 'exec error'}, function() {
			exec('eco "1"')
		})
		`,
		out: &Item{Desc: "exec error [Exception caught, Error: \"/bin/bash -c 'eco \"1\"'\" failed with exit status 127 \"bash: eco: command not found\"]", Wd: "./"},
	},
	// type errors
	{
		in:  `require('hotshell').items = ['test', 2]`,
		err: errMsg("2 error(s) decoding:\n\n* '[0]' expected a map, got 'string'\n* '[1]' expected a map, got 'int64'"),
	},
	{
		in:  `require('hotshell').item({key: []})`,
		err: errMsg("1 error(s) decoding:\n\n* '[0].Key' expected type 'string', got unconvertible type '[]interface {}'"),
	},
	{
		in:  `require('hotshell').item({desc: []})`,
		err: errMsg("1 error(s) decoding:\n\n* '[0].Desc' expected type 'string', got unconvertible type '[]interface {}'"),
	},
	{
		in:  `require('hotshell').item({cmd: []})`,
		err: errMsg("1 error(s) decoding:\n\n* '[0].Cmd' expected type 'string', got unconvertible type '[]interface {}'"),
	},
	{
		in:  `require('hotshell').items = [{items: ["test"]}]`,
		err: errMsg("1 error(s) decoding:\n\n* '[0].Items[0]' expected a map, got 'string'"),
	},
	// doubly nested menu
	{
		in: `
		var hotshell = require('hotshell')
		var item = hotshell.item
		item({key: 't', desc: 'test'}, function() {
			item({key: 'f', desc: hotshell.current.desc + ' > first cmd', cmd: "echo 'first cmd'"})
			item({key: 's', desc: 'second cmd', cmd: "echo 'second cmd'"})
			item({key: 'm', desc: 'submenu'}, function() {
				item({key: 's', desc: hotshell.current.desc + ' > cmd', cmd: "echo 'submenu cmd'"})
			})
		})`,
		out: &Item{
			Key:  "t",
			Desc: "test",
			Wd:   "./",
			Items: []*Item{
				{Key: "f", Desc: "test > first cmd", Cmd: "echo 'first cmd'", Wd: "./"},
				{Key: "s", Desc: "second cmd", Cmd: "echo 'second cmd'", Wd: "./"},
				{Key: "m", Desc: "submenu", Wd: "./",
					Items: []*Item{
						{
							Key:  "s",
							Desc: "submenu > cmd",
							Cmd:  "echo 'submenu cmd'",
							Wd:   "./",
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
		out: &Item{Desc: "1", Wd: "./"},
	},
	{
		in: `
		var item = require('hotshell').item
		var exec = require('hotshell').exec
		item({key: exec('pwd')})
		`,
		out: &Item{Key: cwd(), Wd: "./"},
	},
	{
		in: `
		var item = require('hotshell').item
		eval("item({desc: 'evaled menu'})")
		`,
		out: &Item{Desc: "evaled menu", Wd: "./"},
	},
	// submodule
	{
		in: `
		var item = require('hotshell').item
		var submenu = require('./test/submodule_test.js')
		item({}, function () {
		    submenu()
		})
		`,
		out: &Item{Items: []*Item{{Key: "e", Cmd: Sprintf("%s", cwd()), Wd: "./"}}, Wd: "./"},
	},
	// working directory
	{
		in: `
		var item = require('hotshell').item
		var exec = require('hotshell').exec

		item({desc: exec('pwd')}, function() {
			item({desc: exec('pwd'), wd: './'})
			item({desc: exec('pwd'), wd: '.'})
			item({desc: exec('pwd'), wd: ''})
			item({desc: exec('pwd')})
			item({desc: exec('pwd'), wd: './test/'})
		})
		`,
		out: &Item{
			Desc: cwd(),
			Wd:   "./",
			Items: []*Item{
				{Desc: cwd(), Wd: "././/"},
				{Desc: cwd(), Wd: "././"},
				{Desc: cwd(), Wd: ".//"},
				{Desc: cwd(), Wd: "./"},
				{Desc: cwd(), Wd: "././test//"},
			},
		},
	},
	{
		in: `
		var item = require('hotshell').item
		item({wd: './test/'})
		`,
		out: &Item{Wd: "././test//"},
	},
	{
		in: `
		var item = require('hotshell').item
		item({wd: './unknown-directory/'}, function() {
		    item({desc: require('hotshell').exec('pwd')})
		})
		`,
		out: &Item{
			Desc: "undefined [Exception caught, Error: \"/bin/bash -c 'pwd'\" failed with chdir /home/ju/work/go/src/github.com/julienmoumne/hotshell/cmd/hs/dslrunner/././unknown-directory//: no such file or directory \"\"]",
			Wd:   "././unknown-directory//",
		},
	},
	{
		in: `
		var item = require('hotshell').item
		var exec = require('hotshell').exec
		var moduleItem = require('./test/submodule_test.js')

		item({desc: exec('pwd'), wd: './'}, function() {
			item({wd: './test/', key: 'k'}, function () {
			    item({desc: exec('pwd')})
			    moduleItem()
			    item({wd: './subdir/', key: 'k'}, function () {
			        item({desc: exec('pwd')})
			    })
			    item({key: 'q'}, function () {
			        item({desc: exec('pwd')})
			    })
			})
			item({desc: exec('pwd')})
		})
		`,
		out: &Item{
			Desc: cwd(),
			Wd:   "././/",
			Items: []*Item{
				{Key: "k",
					Wd: "././/./test//",
					Items: []*Item{
						{Wd: "././/./test//", Desc: Sprintf("%s/test", cwd())},
						{Wd: "././/./test//", Key: "e", Cmd: Sprintf("%s/test", cwd())},
						{Wd: "././/./test//./subdir//", Key: "k", Items: []*Item{
							{Wd: "././/./test//./subdir//", Desc: Sprintf("%s/test/subdir", cwd())},
						}},
						{Wd: "././/./test//", Key: "q", Items: []*Item{
							{Wd: "././/./test//", Desc: Sprintf("%s/test", cwd())},
						}},
					}},
				{Wd: "././/", Desc: cwd()},
			},
		},
	},
}

var a *assert.Assertions

type testCase struct {
	in  string
	out *Item
	err string
}

func errMsg(msg string) string {
	return Sprintf("Error while reading the menu definition\n%s", msg)
}

func cwd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

func TestDslRunner(t *testing.T) {
	a = assert.New(t)
	for _, tt := range tests {
		runTest(tt)
	}
}

func runTest(t testCase) {
	setupTest(t)
	validateTest(t)
}

func setupTest(t testCase) {
	adjustParentLinks(t.out, nil)
}

func validateTest(t testCase) {
	actualOut, err := (&dslrunner.DslRunner{}).Run(t.in)
	if t.err != "" {
		a.NotNil(err)
		a.Equal(t.err, err.Error())
	} else {
		a.Nil(err)
		a.Equal(t.out, actualOut)
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
