package item

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/formatter"
	"io"
	"os"
	"os/exec"
)

type CmdActivator struct {
	Out   io.Writer
	cmd   *exec.Cmd
	osCwd string
	item  *Item
}

func (c *CmdActivator) Activate(item *Item) {
	c.item = item
	c.cmd = exec.Command("bash", "-c", item.Cmd)
	c.print(formatter.ExecutedCmdFmt(" %s\n\n", formatter.FormatCommand(c.cmd)))
	var err error
	c.osCwd, err = os.Getwd()
	if !c.displayErrIfNotNil(err) {
		c.runCmd()
	}
	c.print("\n")
}

func (c *CmdActivator) runCmd() {
	c.cmd.Stdin = os.Stdin
	c.cmd.Stdout = c.Out
	c.cmd.Stderr = os.Stderr
	c.cmd.Dir = fmt.Sprintf("%s/%s", c.osCwd, c.item.Wd)
	c.displayErrIfNotNil(c.cmd.Start())
	c.displayErrIfNotNil(c.cmd.Wait())
}

func (c *CmdActivator) displayErrIfNotNil(err error) bool {
	if err == nil {
		return false
	}
	c.printf("\n%s\n", err)
	return true
}

func (c *CmdActivator) printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c.Out, format, a...)
}

func (c *CmdActivator) print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(c.Out, a...)
}
