package item_test

import (
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	. "gopkg.in/check.v1"
	"testing"
)

func TestItem(t *testing.T) { TestingT(t) }

type ItemTestSuite struct{}

var _ = Suite(&ItemTestSuite{})

func (s *ItemTestSuite) TestIsCmd(c *C) {
	it := item.NewItem("", "", "echo 'test'")
	checkIsCmd(c, it, true)
}

func (s *ItemTestSuite) TestMenuWithItemsIsNotCmd(c *C) {
	it := item.NewItem("", "", "")
	it.AddItem(it)
	checkIsCmd(c, it, false)
}

func (s *ItemTestSuite) TestMenuWithItemsWithCmdIsNotCmd(c *C) {
	it := item.NewItem("", "", "echo 'test'")
	it.AddItem(it)
	checkIsCmd(c, it, false)
}

func (s *ItemTestSuite) TestEmptyMenuIsNotCmd(c *C) {
	it := item.NewItem("", "", "")
	checkIsCmd(c, it, false)
}

func checkIsCmd(c *C, it *item.Item, isCmd bool) {
	c.Check(it.IsCmd(), Equals, isCmd)
}

var descTests = []struct {
	item       *item.Item
	desc       string
	inMenuDesc string
}{

	// Without children
	{singleItem("", "", ""), "missing-desc", "missing-desc"},
	{singleItem("k", "", ""), "missing-desc", "k missing-desc"},
	{singleItem("", "desc", ""), "desc", "desc"},
	{singleItem("", "", "cmd"), "cmd", "cmd"},
	{singleItem("k", "desc", ""), "desc", "k desc"},
	{singleItem("k", "", "cmd"), "cmd", "k cmd"},
	{singleItem("", "desc", "cmd"), "desc cmd", "desc cmd"},
	{singleItem("k", "desc", "cmd"), "desc cmd", "k desc cmd"},

	// With children
	{itemWithChild("", "", ""), "missing-desc", "missing-desc"},
	{itemWithChild("k", "", ""), "missing-desc", "k missing-desc"},
	{itemWithChild("", "desc", ""), "desc", "desc"},
	{itemWithChild("", "", "cmd"), "missing-desc", "missing-desc"},
	{itemWithChild("k", "desc", ""), "desc", "k desc"},
	{itemWithChild("k", "", "cmd"), "missing-desc", "k missing-desc"},
	{itemWithChild("", "desc", "cmd"), "desc", "desc"},
	{itemWithChild("k", "desc", "cmd"), "desc", "k desc"},

	// With parent and child
	{itemWithParentAndChild("", "", ""), "missing-desc", "missing-desc >"},
	{itemWithParentAndChild("k", "", ""), "missing-desc", "k missing-desc >"},
	{itemWithParentAndChild("", "desc", ""), "desc", "desc >"},
	{itemWithParentAndChild("", "", "cmd"), "missing-desc", "missing-desc >"},
	{itemWithParentAndChild("k", "desc", ""), "desc", "k desc >"},
	{itemWithParentAndChild("k", "", "cmd"), "missing-desc", "k missing-desc >"},
	{itemWithParentAndChild("", "desc", "cmd"), "desc", "desc >"},
	{itemWithParentAndChild("k", "desc", "cmd"), "desc", "k desc >"},
}

func (s *ItemTestSuite) TestDesc(c *C) {
	for _, t := range descTests {
		c.Check(t.item.GetDesc(), Equals, t.desc)
		c.Check(t.item.GetInMenuDesc(), Equals, t.inMenuDesc)
	}
}

func singleItem(key string, desc string, cmd string) *item.Item {
	return item.NewItem(key, desc, cmd)
}

func itemWithChild(key string, desc string, cmd string) *item.Item {
	it := item.NewItem(key, desc, cmd)
	it.AddItem(item.NewItem(key, desc, cmd))
	return it
}

func itemWithParentAndChild(key string, desc string, cmd string) *item.Item {
	it := item.NewItem(key, desc, cmd)
	child := item.NewItem(key, desc, cmd)
	child.AddItem(item.NewItem(key, desc, cmd))
	it.AddItem(child)
	return child
}
