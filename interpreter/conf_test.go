package interpreter_test

import (
	"github.com/julienmoumne/hotshell/interpreter"
	. "gopkg.in/check.v1"
)

type ConfTestSuite struct{}

var _ = Suite(&ConfTestSuite{})

var confTests = []struct {
	in  map[string]interface{}
	out interpreter.Conf
}{
	// Various Otto types for loop indexes
	{
		map[string]interface{}{ interpreter.BREADCRUMB_TYPE_KEY: 1 },
		interpreter.Conf{ BreadcrumbType: interpreter.BREADCRUMB_TYPE_DEFAULT },
	},
	{
		map[string]interface{}{ interpreter.BREADCRUMB_TYPE_KEY: int64(1) },
		interpreter.Conf{ BreadcrumbType: interpreter.BREADCRUMB_TYPE_DEFAULT },
	},
	{
		map[string]interface{}{ interpreter.BREADCRUMB_TYPE_KEY: 1.0 },
		interpreter.Conf{ BreadcrumbType: interpreter.BREADCRUMB_TYPE_DEFAULT },
	},
	{
		map[string]interface{}{ interpreter.BREADCRUMB_TYPE_KEY: float64(1) },
		interpreter.Conf{ BreadcrumbType: interpreter.BREADCRUMB_TYPE_DEFAULT },
	},
	{
		map[string]interface{}{ interpreter.BREADCRUMB_TYPE_KEY: "1" },
		interpreter.Conf{ BreadcrumbType: interpreter.BREADCRUMB_TYPE_DEFAULT },
	},

	// Conf does not validate anything and is fail-safe
	{
		nil,
		interpreter.Conf{ BreadcrumbType: interpreter.BREADCRUMB_TYPE_DEFAULT },
	},
	{
		map[string]interface{}{},
		interpreter.Conf{ BreadcrumbType: interpreter.BREADCRUMB_TYPE_DEFAULT },
	},
	{
		map[string]interface{}{
			"test": map[string]interface{}{},
			"dummy": map[string]interface{}{},
		},
		interpreter.Conf{ BreadcrumbType: interpreter.BREADCRUMB_TYPE_DEFAULT },
	},
	{
		map[string]interface{}{ interpreter.BREADCRUMB_TYPE_KEY: map[string]interface{}{} },
		interpreter.Conf{ BreadcrumbType: interpreter.BREADCRUMB_TYPE_DEFAULT },
	},
	
	// Valid input tests
	
	{
		map[string]interface{}{ interpreter.BREADCRUMB_TYPE_KEY: "horizontal" },
		interpreter.Conf{ BreadcrumbType: "horizontal" },
	},
	{
		map[string]interface{}{ interpreter.BREADCRUMB_TYPE_KEY: "vertical" },
		interpreter.Conf{ BreadcrumbType: "vertical" },
	},
}

func (s *ConfTestSuite) TestConf(c *C) {
	for _, tt := range confTests {
		actualOut := interpreter.NewConf(tt.in)
		c.Check(actualOut, DeepEquals, tt.out)
	}
}
