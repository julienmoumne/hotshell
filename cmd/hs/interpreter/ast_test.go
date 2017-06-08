package interpreter_test

import (
	. "github.com/julienmoumne/hotshell/cmd/hs/interpreter"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	in  interface{}
	out []Ast
	err bool
}{
	// Various Otto types for integers
	{
		in:  []map[string]interface{}{{"desc": 1, "key": int64(1)}},
		out: []Ast{{Desc: "1", Key: "1"}},
	},
	{
		in:  []map[string]interface{}{{"desc": 1.0, "key": float64(1)}},
		out: []Ast{{Desc: "1", Key: "1"}},
	},
	// Empty values
	{
		in:  nil,
		out: nil,
	},
	{
		in:  []map[string]interface{}{},
		out: []Ast{},
	},
	{
		in:  []map[string]interface{}{{}},
		out: []Ast{{}},
	},
	{
		in:  []map[string]interface{}{{}, {}},
		out: []Ast{{}, {}},
	},
	{
		in:  []map[string]interface{}{{"ignoredKey": []string{"test"}}},
		out: []Ast{{}},
	},
	// Type errors
	{
		in:  []interface{}{"test", 2},
		err: true,
	},
	{
		in:  []map[string]interface{}{{"key": []string{}}},
		err: true,
	},
	{
		in:  []map[string]interface{}{{"desc": []string{}}},
		err: true,
	},
	{
		in:  []map[string]interface{}{{"cmd": []string{}}},
		err: true,
	},
	{
		in:  []map[string]interface{}{{"items": []string{"test"}}},
		err: true,
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
		out: []Ast{{Key: "t", Desc: "test",
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
	},
}

func TestAst(t *testing.T) {
	for _, test := range tests {
		a, err := NewAst(test.in)
		if test.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, test.out, a)
		}
	}
}
