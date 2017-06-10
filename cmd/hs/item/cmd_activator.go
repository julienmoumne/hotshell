package item

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/formatter"
	"os"
	"os/exec"
)

type CmdActivator struct {
	item *Item
}

func (c *CmdActivator) Activate(item *Item) {
	c.item = item
	command := exec.Command("bash", "-c", c.item.Cmd)

	c.print(formatter.ExecutedCmdFmt(" %s\n\n", formatter.FormatCommand(command)))

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Start(); err != nil {
		c.printf("\n%s\n", err)
	}

	if err := command.Wait(); err != nil {
		c.printf("\n%s\n", err)
	}

	c.print("\n")
}

func (c *CmdActivator) printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stdout, format, a...)
}

func (c *CmdActivator) print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(os.Stdout, a...)
}
