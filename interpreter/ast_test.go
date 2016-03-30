package interpreter_test

import (
	"github.com/julienmoumne/hs/interpreter"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type AstTestSuite struct{}

var _ = Suite(&AstTestSuite{})

var tests = []struct {
	in  []map[string]interface{}
	out []interpreter.Ast
}{
	// Various Otto types for loop indexes
	{
		[]map[string]interface{}{map[string]interface{}{interpreter.KEY_KEY: 1}},
		[]interpreter.Ast{{Key: "1", Items: []interpreter.Ast{}}},
	},
	{
		[]map[string]interface{}{map[string]interface{}{interpreter.KEY_KEY: int64(1)}},
		[]interpreter.Ast{{Key: "1", Items: []interpreter.Ast{}}},
	},
	{
		[]map[string]interface{}{map[string]interface{}{interpreter.KEY_KEY: 1.0}},
		[]interpreter.Ast{{Key: "1", Items: []interpreter.Ast{}}},
	},
	{
		[]map[string]interface{}{map[string]interface{}{interpreter.KEY_KEY: float64(1)}},
		[]interpreter.Ast{{Key: "1", Items: []interpreter.Ast{}}},
	},
	{
		[]map[string]interface{}{map[string]interface{}{interpreter.KEY_KEY: "1"}},
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
			interpreter.KEY_KEY:  "t",
			interpreter.DESC_KEY: "test",
			interpreter.ITEMS_KEY: []map[string]interface{}{
				map[string]interface{}{
					interpreter.KEY_KEY:  "f",
					interpreter.DESC_KEY: "first cmd",
					interpreter.CMD_KEY:  "echo 'first cmd'",
				},
				map[string]interface{}{
					interpreter.KEY_KEY:  "s",
					interpreter.DESC_KEY: "second cmd",
					interpreter.CMD_KEY:  "echo 'second cmd'",
				},
				map[string]interface{}{
					interpreter.KEY_KEY:  "m",
					interpreter.DESC_KEY: "submenu",
					interpreter.ITEMS_KEY: []map[string]interface{}{
						map[string]interface{}{
							interpreter.KEY_KEY:  "s",
							interpreter.DESC_KEY: "submenu cmd",
							interpreter.CMD_KEY:  "echo 'submenu cmd'",
						},
					},
				},
			},
		}},
		[]interpreter.Ast{{Key: "t", Desc: "test",
			Items: []interpreter.Ast{
				interpreter.Ast{Key: "f", Desc: "first cmd", Cmd: "echo 'first cmd'", Items: []interpreter.Ast{}},
				interpreter.Ast{Key: "s", Desc: "second cmd", Cmd: "echo 'second cmd'", Items: []interpreter.Ast{}},
				interpreter.Ast{Key: "m", Desc: "submenu",
					Items: []interpreter.Ast{
						interpreter.Ast{Key: "s", Desc: "submenu cmd", Cmd: "echo 'submenu cmd'", Items: []interpreter.Ast{}},
					},
				},
			},
		}},
	},
}

func (s *AstTestSuite) TestAst(c *C) {
	for _, tt := range tests {
		actualOut := interpreter.NewAst(tt.in)
		c.Check(actualOut, DeepEquals, tt.out)
	}
}
