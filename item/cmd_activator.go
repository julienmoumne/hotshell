package item

import (
	"fmt"
	"github.com/julienmoumne/hotshell/formatter"
	"github.com/julienmoumne/hotshell/interpreter"
	"io"
	"os"
	"os/exec"
)

type CmdActivator struct {
	conf interpreter.Conf
	item *Item
	Out  io.Writer
}

func (c *CmdActivator) Activate(conf interpreter.Conf, item *Item) *Item {
	c.conf = conf
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
	return c.item.Parent
}

func (c *CmdActivator) printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c.Out, format, a...)
}

func (c *CmdActivator) print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(c.Out, a...)
}
