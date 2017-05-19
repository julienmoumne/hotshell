package item_test

import (
	"bytes"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	a         *assert.Assertions
	buf       *bytes.Buffer
	activator item.MenuActivator
	tests     []testCase
)

func init() {
	emptyDescInNestedMenu := item.NewItem("", "notice", "")
	emptyDescInNestedMenu.AddItem(item.NewItem("", "notice", ""))
	cmdWithoutDesc := item.NewItem("", "cmd-without-desc", "")
	cmdWithoutDesc.AddItem(item.NewItem("k", "", "cmd-without-desc"))

	tests = []testCase{
		{
			in{item.NewItem("", "empty-menu", "")},
			out{" empty-menu\n" +
				"\n" +
				" no items found\n" +
				"\n" +
				" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n" +
				"\n"},
		},
		{
			in{item.NewItem("", "", "")},
			out{" missing-desc\n" +
				"\n" +
				" no items found\n" +
				"\n" +
				" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n" +
				"\n"},
		},
		{
			in{emptyDescInNestedMenu},
			out{" notice\n" +
				"\n" +
				" notice\n" +
				"\n" +
				" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n" +
				"\n"},
		},
		{
			in{cmdWithoutDesc},
			out{" cmd-without-desc\n" +
				"\n" +
				" k cmd-without-desc\n" +
				"\n" +
				" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n" +
				"\n"},
		},
	}
}

type (
	in struct {
		item *item.Item
	}
	out struct {
		str string
	}
	testCase struct {
		in
		out
	}
)

func TestMenuActivator(t *testing.T) {
	a = assert.New(t)
	for _, test := range tests {
		runTest(test)
	}
}

func runTest(t testCase) {
	setupTest()
	validateTest(t)
}

func setupTest() {
	buf = &bytes.Buffer{}
	activator = item.MenuActivator{Out: buf}
}

func validateTest(t testCase) {
	activator.Activate(t.in.item)
	a.Equal(t.out.str, buf.String())
}
