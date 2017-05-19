package interpreter_test

import (
	"github.com/julienmoumne/hotshell/cmd/hs/interpreter"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	in  []map[string]interface{}
	out []interpreter.Ast
}{
	// Various Otto types for loop indexes
	{
		[]map[string]interface{}{map[string]interface{}{interpreter.KeyPropName: 1}},
		[]interpreter.Ast{{Key: "1", Items: []interpreter.Ast{}}},
	},
	{
		[]map[string]interface{}{map[string]interface{}{interpreter.KeyPropName: int64(1)}},
		[]interpreter.Ast{{Key: "1", Items: []interpreter.Ast{}}},
	},
	{
		[]map[string]interface{}{map[string]interface{}{interpreter.KeyPropName: 1.0}},
		[]interpreter.Ast{{Key: "1", Items: []interpreter.Ast{}}},
	},
	{
		[]map[string]interface{}{map[string]interface{}{interpreter.KeyPropName: float64(1)}},
		[]interpreter.Ast{{Key: "1", Items: []interpreter.Ast{}}},
	},
	{
		[]map[string]interface{}{map[string]interface{}{interpreter.KeyPropName: "1"}},
		[]interpreter.Ast{{Key: "1", Items: []interpreter.Ast{}}},
	},

	// Ast does not validate anything and is fail-safe
	{
		nil,
		[]interpreter.Ast{},
	},
	{
		[]map[string]interface{}{},
		[]interpreter.Ast{},
	},
	{
		[]map[string]interface{}{map[string]interface{}{}},
		[]interpreter.Ast{{Items: []interpreter.Ast{}}},
	},
	{
		[]map[string]interface{}{
			map[string]interface{}{},
			map[string]interface{}{},
		},
		[]interpreter.Ast{
			{Items: []interpreter.Ast{}},
			{Items: []interpreter.Ast{}},
		},
	},

	// Doubly nested menu
	{
		[]map[string]interface{}{map[string]interface{}{
			interpreter.KeyPropName:  "t",
			interpreter.DescPropName: "test",
			interpreter.ItemsPropName: []map[string]interface{}{
				map[string]interface{}{
					interpreter.KeyPropName:  "f",
					interpreter.DescPropName: "first cmd",
					interpreter.CmdPropName:  "echo 'first cmd'",
				},
				map[string]interface{}{
					interpreter.KeyPropName:  "s",
					interpreter.DescPropName: "second cmd",
					interpreter.CmdPropName:  "echo 'second cmd'",
				},
				map[string]interface{}{
					interpreter.KeyPropName:  "m",
					interpreter.DescPropName: "submenu",
					interpreter.ItemsPropName: []map[string]interface{}{
						map[string]interface{}{
							interpreter.KeyPropName:  "s",
							interpreter.DescPropName: "submenu cmd",
							interpreter.CmdPropName:  "echo 'submenu cmd'",
						},
					},
				},
			},
		}},
		[]interpreter.Ast{{Key: "t", Desc: "test",
			Items: []interpreter.Ast{
				{Key: "f", Desc: "first cmd", Cmd: "echo 'first cmd'", Items: []interpreter.Ast{}},
				{Key: "s", Desc: "second cmd", Cmd: "echo 'second cmd'", Items: []interpreter.Ast{}},
				{Key: "m", Desc: "submenu",
					Items: []interpreter.Ast{
						{Key: "s", Desc: "submenu cmd", Cmd: "echo 'submenu cmd'", Items: []interpreter.Ast{}},
					},
				},
			},
		}},
	},
}

func TestAst(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.out, interpreter.NewAst(test.in))
	}
}
