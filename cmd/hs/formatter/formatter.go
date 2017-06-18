package formatter

import (
	"fmt"
	"github.com/fatih/color"
	"os/exec"
	"strings"
)

func newColorSprinter(attributes ...color.Attribute) func(string, ...interface{}) string {
	return color.New(attributes...).SprintfFunc()
}

var ExecutedCmdFmt = newColorSprinter(color.FgBlue)
var CmdDefFmt = newColorSprinter(color.FgGreen)
var KeyHintFmt = newColorSprinter(color.FgGreen)
var KeyActivatedFmt = newColorSprinter(color.FgYellow)
var ParentMenuFmt = newColorSprinter(color.FgYellow)
var ActiveMenuFmt = newColorSprinter(color.FgYellow, color.Underline)
var HelpFmt = newColorSprinter(color.FgGreen)
var WdFmt = newColorSprinter(color.FgYellow)

func FormatCommand(cmd *exec.Cmd) string {
	return fmt.Sprintf(
		"%s %s '%s'",
		cmd.Path,
		cmd.Args[1],
		strings.Join(cmd.Args[2:], " "),
	)
}
