package item

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/formatter"
	"github.com/julienmoumne/hotshell/cmd/hs/settings"
	"io"
)

type MenuPrinter struct {
	item *Item
	Out  io.Writer
	keys settings.Keys
}

func (m *MenuPrinter) Print(item *Item, keys settings.Keys) {
	m.item = item
	m.keys = keys
	m.printBreadcrumb()
	m.printItems()
	m.printPrompt()
}

func (m *MenuPrinter) printPrompt() {
	m.printf(
		" %v back, %v bash, %v repeat, %v reload, %v quit",
		formatter.HelpFmt(settings.KeyName(m.keys.Back)),
		formatter.HelpFmt(settings.KeyName(m.keys.Bash)),
		formatter.HelpFmt(settings.KeyName(m.keys.Repeat)),
		formatter.HelpFmt(settings.KeyName(m.keys.Reload)),
		formatter.HelpFmt("^d or ^c"),
	)
	m.print("\n")
	m.print("\n")
	m.print(formatter.KeyActivatedFmt(" ? "))
}

func (m *MenuPrinter) printItems() {
	if len(m.item.Items) == 0 {
		m.printf(" %s\n", "no items found")
	}
	for _, item := range m.item.Items {
		m.printf(" %s\n", item.GetInMenuDesc())
	}
	m.print("\n")
}

func (m *MenuPrinter) printBreadcrumb() {
	var bc string
	for curMenu := m.item.Parent; curMenu != nil; curMenu = curMenu.Parent {
		bc = fmt.Sprintf(" %s\n%s", curMenu.GetDesc(), bc)
	}
	m.print(formatter.ParentMenuFmt("%s", bc))

	m.printf(" %s", formatter.ActiveMenuFmt("%s", m.item.GetDesc()))
	m.print("\n")
	m.print("\n")
}

func (m *MenuPrinter) printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(m.Out, format, a...)
}

func (m *MenuPrinter) print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(m.Out, a...)
}
