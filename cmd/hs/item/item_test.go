package item_test

import (
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsCmd(t *testing.T) {
	it := &item.Item{Cmd: "echo 'test'"}
	checkIsCmd(t, it, true)
}

func TestMenuWithItemsIsNotCmd(t *testing.T) {
	it := &item.Item{}
	it.AddItem(it)
	checkIsCmd(t, it, false)
}

func TestMenuWithItemsWithCmdIsNotCmd(t *testing.T) {
	it := &item.Item{Cmd: "echo 'test'"}
	it.AddItem(it)
	checkIsCmd(t, it, false)
}

func TestEmptyMenuIsNotCmd(t *testing.T) {
	it := &item.Item{}
	checkIsCmd(t, it, false)
}

func checkIsCmd(t *testing.T, it *item.Item, isCmd bool) {
	assert.Equal(t, isCmd, it.IsCmd())
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

func TestDesc(t *testing.T) {
	for _, test := range descTests {
		a := assert.New(t)
		a.Equal(test.desc, test.item.GetDesc())
		a.Equal(test.inMenuDesc, test.item.GetInMenuDesc())
	}
}

func singleItem(key string, desc string, cmd string) *item.Item {
	return &item.Item{Key: key, Desc: desc, Cmd: cmd}
}

func itemWithChild(key string, desc string, cmd string) *item.Item {
	it := &item.Item{Key: key, Desc: desc, Cmd: cmd}
	it.AddItem(&item.Item{Key: key, Desc: desc, Cmd: cmd})
	return it
}

func itemWithParentAndChild(key string, desc string, cmd string) *item.Item {
	it := &item.Item{Key: key, Desc: desc, Cmd: cmd}
	child := &item.Item{Key: key, Desc: desc, Cmd: cmd}
	child.AddItem(&item.Item{Key: key, Desc: desc, Cmd: cmd})
	it.AddItem(child)
	return child
}
