package item

import (
	"fmt"
	"github.com/julienmoumne/hotshell/formatter"
	"github.com/julienmoumne/hotshell/interpreter"
	"io"
)

type MenuActivator struct {
	conf interpreter.Conf
	item *Item
	Out  io.Writer
}

func (m *MenuActivator) Activate(conf interpreter.Conf, item *Item) *Item {
	m.conf = conf
	m.item = item
	m.printBreadcrumb()
	m.printItems()
	m.printHelp()
	return m.item
}

func (m *MenuActivator) printHelp() {
	m.printf(
		" %v back, %v bash, %v repeat, %v reload, %v quit",
		formatter.HelpFmt(PREVIOUS_MENU_KEY.String()),
		formatter.HelpFmt(BASH_KEY.String()),
		formatter.HelpFmt(REPEAT_KEY.String()),
		formatter.HelpFmt(RELOAD_KEY.String()),
		formatter.HelpFmt("^d or ^c"),
	)
	m.print("\n")
	m.print("\n")
}

func (m *MenuActivator) printItems() {
	if len(m.item.Items) == 0 {
		m.printf(" %s\n", "no items found")
	}
	for _, item := range m.item.Items {
		m.printf(" %s\n", item.GetInMenuDesc())
	}
	m.print("\n")
}

func (m *MenuActivator) printBreadcrumb() {
	var bc string
	for curMenu := m.item.Parent; curMenu != nil; curMenu = curMenu.Parent {
		bc = fmt.Sprintf(" %s\n%s", curMenu.GetDesc(), bc)
	}
	m.print(formatter.ParentMenuFmt("%s", bc))

	m.printf(" %s", formatter.ActiveMenuFmt("%s", m.item.GetDesc()))
	m.print("\n")
	m.print("\n")
}

func (m *MenuActivator) printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(m.Out, format, a...)
}

func (m *MenuActivator) print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(m.Out, a...)
}
