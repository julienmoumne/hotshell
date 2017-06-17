package engine_test

import (
	"errors"
	"github.com/julienmoumne/hotshell/cmd/hs/dslrunner"
	. "github.com/julienmoumne/hotshell/cmd/hs/engine"
	. "github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/julienmoumne/hotshell/cmd/hs/settings"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []testCase{
	// input error straight after startup
	{
		keys: []string{
			errorWhileReadingUserInput,
		},
		err: errorWhileReadingUserInput,
		events: []interface{}{
			MenuEvent{Item: root},
		},
	},
	// input error in submenu
	{
		keys: []string{
			"m",
			errorWhileReadingUserInput,
		},
		err: errorWhileReadingUserInput,
		events: []interface{}{
			MenuEvent{Item: root},
			ValidKeyEvent{Key: "m"},
			MenuEvent{Item: submenu},
		},
	},
	// eot straight after startup
	{
		keys: []string{
			eot,
		},
		events: []interface{}{
			MenuEvent{Item: root},
		},
	},
	// back issued in root menu
	{
		keys: []string{
			keys.Back,
			eot,
		},
		events: []interface{}{
			MenuEvent{Item: root},
			ValidKeyEvent{Key: keys.Back},
			MenuEvent{Item: root},
		},
	},
	// repeat issued in root menu
	{
		keys: []string{
			keys.Repeat,
			eot,
		},
		events: []interface{}{
			MenuEvent{Item: root},
			ValidKeyEvent{Key: keys.Repeat},
			MenuEvent{Item: root},
		},
	},
	// repeat cmd
	{
		keys: []string{
			"s",
			keys.Repeat,
			eot,
		},
		events: []interface{}{
			MenuEvent{Item: root},
			ValidKeyEvent{Key: "s"},
			CmdEvent{Item: secondCmd},
			MenuEvent{Item: root},
			ValidKeyEvent{Key: keys.Repeat},
			CmdEvent{Item: secondCmd},
			MenuEvent{Item: root},
		},
	},
	// up down up
	{
		keys: []string{
			"m",
			"s",
			keys.Back,
			"s",
			eot,
		},
		events: []interface{}{
			MenuEvent{Item: root},
			ValidKeyEvent{Key: "m"},
			MenuEvent{Item: submenu},
			ValidKeyEvent{Key: "s"},
			CmdEvent{Item: submenuCmd},
			MenuEvent{Item: submenu},
			ValidKeyEvent{Key: keys.Back},
			MenuEvent{Item: root},
			ValidKeyEvent{Key: "s"},
			CmdEvent{Item: secondCmd},
			MenuEvent{Item: root},
		},
	},
	// nonexistent key
	{
		keys: []string{
			"a",
			eot,
		},
		events: []interface{}{
			MenuEvent{Item: root},
		},
	},
	// reload straight after startup
	{
		keys: []string{
			keys.Reload,
			eot,
		},
		events: []interface{}{
			MenuEvent{Item: root},
			ValidKeyEvent{Key: keys.Reload},
		},
		reload: true,
	},
	// bash then repeat
	{
		keys: []string{
			keys.Bash,
			keys.Repeat,
			eot,
		},
		events: []interface{}{
			MenuEvent{Item: root},
			ValidKeyEvent{Key: keys.Bash},
			CmdEvent{Item: BashCmd(keys.Bash)},
			MenuEvent{Item: root},
			ValidKeyEvent{Key: keys.Repeat},
			CmdEvent{Item: BashCmd(keys.Bash)},
			MenuEvent{Item: root},
		},
		reload: true,
	},
}

const (
	errorWhileReadingUserInput = "[ERROR_WHEN_READING_USER_INPUT]"
	eot                        = string(4)
)

var (
	a          *assert.Assertions
	dispatcher *dispatcherStub
	controller Controller
	keys       = settings.Defaults().Keys
	root, _    = (&dslrunner.DslRunner{}).Run(`
		var item = require('hotshell').item
		item({key: 't', desc: 'test'}, function() {
			item({key: 'f', desc: 'first cmd', cmd: "echo 'first cmd'"})
			item({key: 's', desc: 'second cmd', cmd: "echo 'second cmd'"})
			item({key: 'm', desc: 'submenu'}, function() {
				item({key: 's', desc: 'submenu cmd', cmd: "echo 'submenu cmd'"})
			})
		})`)
	secondCmd, _  = root.GetItem("s")
	submenu, _    = root.GetItem("m")
	submenuCmd, _ = submenu.GetItem("s")
)

type testCase struct {
	keys   []string
	err    string
	reload bool
	events []interface{}
}

func TestController(t *testing.T) {
	a = assert.New(t)
	for _, t := range tests {
		runTest(t)
	}
}

func runTest(t testCase) {
	setupTest(t)
	validateTest(t)
}

func setupTest(t testCase) {
	controller = Controller{}
	dispatcher = &dispatcherStub{keys: t.keys}
}

func validateTest(t testCase) {
	reload, err := controller.Start(keys, root, dispatcher)
	a.Equal(reload, reload)
	if t.err != "" {
		a.NotNil(err)
		a.Equal(t.err, err.Error())
	} else {
		a.Nil(err)
	}
	a.Equal(t.events, dispatcher.recordedEvents)
}

type dispatcherStub struct {
	recordedEvents []interface{}
	keys           []string
	nextKey        int
}

func (d *dispatcherStub) DispatchEvent(rawEvent interface{}) {
	d.recordedEvents = append(d.recordedEvents, rawEvent)
}
func (d *dispatcherStub) ReadUserInput() (string, error) {
	key := d.keys[d.nextKey]
	d.nextKey++
	if key == errorWhileReadingUserInput {
		return "", errors.New(errorWhileReadingUserInput)
	}
	return key, nil
}
func (d *dispatcherStub) Cleanup() {}