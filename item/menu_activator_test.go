package item_test

import (
	"bytes"
	"github.com/julienmoumne/hotshell/item"
	. "gopkg.in/check.v1"
	"testing"
)

func TestMenuActivator(t *testing.T) { TestingT(t) }

type MenuActivatorTestSuite struct {
	buf       *bytes.Buffer
	activator item.MenuActivator
}

var _ = Suite(&MenuActivatorTestSuite{})

func (s *MenuActivatorTestSuite) SetUpTest(c *C) {
	s.buf = &bytes.Buffer{}
	s.activator = item.MenuActivator{Out: s.buf}
}

func (s *MenuActivatorTestSuite) TestEmptyMenu(c *C) {
	it := item.NewItem("", "empty-menu", "")
	s.validateOut(c, it,
		" empty-menu\n"+
			"\n"+
			" no items found\n"+
			"\n"+
			" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n"+
			"\n",
	)
}

func (s *MenuActivatorTestSuite) TestEmptyDescEmptyMenu(c *C) {
	it := item.NewItem("", "", "")
	s.validateOut(c, it,
		" missing-desc\n"+
			"\n"+
			" no items found\n"+
			"\n"+
			" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n"+
			"\n",
	)
}

func (s *MenuActivatorTestSuite) TestEmptyDescInNestedMenu(c *C) {
	it := item.NewItem("", "notice", "")
	it.AddItem(item.NewItem("", "notice", ""))
	s.validateOut(c, it,
		" notice\n"+
			"\n"+
			" notice\n"+
			"\n"+
			" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n"+
			"\n",
	)
}

func (s *MenuActivatorTestSuite) TestCmdWithoutDesc(c *C) {
	it := item.NewItem("", "cmd-without-desc", "")
	it.AddItem(item.NewItem("k", "", "cmd-without-desc"))
	s.validateOut(c, it,
		" cmd-without-desc\n"+
			"\n"+
			" k cmd-without-desc\n"+
			"\n"+
			" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n"+
			"\n",
	)
}

func (s *MenuActivatorTestSuite) validateOut(c *C, it *item.Item, out string) {
	s.activator.Activate(it)
	c.Check(s.buf.String(), Equals, out)
}
