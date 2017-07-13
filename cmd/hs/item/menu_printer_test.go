package item_test

import (
	"bytes"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/julienmoumne/hotshell/cmd/hs/settings"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	menuPrinterAssert *assert.Assertions
	menuPrinterBuf    *bytes.Buffer
	printer           item.MenuPrinter
	menuPrinterTests  []menuPrinterTestCase
)

func init() {
	twoLevelMenu := &item.Item{Desc: "lvl1"}
	subMenu := &item.Item{Desc: "lvl2"}
	twoLevelMenu.AddItem(subMenu)
	cmdWithoutDesc := &item.Item{Desc: "cmd-without-desc"}
	cmdWithoutDesc.AddItem(&item.Item{Key: "k", Cmd: "cmd-without-desc"})

	menuPrinterTests = []menuPrinterTestCase{
		// empty menu
		{
			in{&item.Item{}},
			out{" missing-desc\n" +
				"\n" +
				" no items found\n" +
				"\n" +
				" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n" +
				"\n ? "},
		},
		// empty menu with description
		{
			in{&item.Item{Desc: "empty-menu"}},
			out{" empty-menu\n" +
				"\n" +
				" no items found\n" +
				"\n" +
				" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n" +
				"\n ? "},
		},
		// two level menu - lvl 1
		{
			in{twoLevelMenu},
			out{" lvl1\n" +
				"\n" +
				" lvl2\n" +
				"\n" +
				" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n" +
				"\n ? "},
		},
		// two level menu - lvl 2
		{
			in{subMenu},
			out{" lvl1 > lvl2\n" +
				"\n" +
				" no items found\n" +
				"\n" +
				" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n" +
				"\n ? "},
		},
		// menu with desc-less command
		{
			in{cmdWithoutDesc},
			out{" cmd-without-desc\n" +
				"\n" +
				" k cmd-without-desc\n" +
				"\n" +
				" spacebar back, tabulation bash, return repeat, backspace reload, ^d or ^c quit\n" +
				"\n ? "},
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
	menuPrinterTestCase struct {
		in
		out
	}
)

func TestMenuPrinter(t *testing.T) {
	menuPrinterAssert = assert.New(t)
	for _, test := range menuPrinterTests {
		runMenuPrinterTest(test)
	}
}

func runMenuPrinterTest(t menuPrinterTestCase) {
	setupMenuPrinterTest()
	validateMenuPrinterTest(t)
}

func setupMenuPrinterTest() {
	menuPrinterBuf = &bytes.Buffer{}
	printer = item.MenuPrinter{Out: menuPrinterBuf}
}

func validateMenuPrinterTest(t menuPrinterTestCase) {
	printer.Print(t.in.item, settings.Defaults().Keys)
	menuPrinterAssert.Equal(t.out.str, menuPrinterBuf.String())
}
