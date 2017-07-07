package item_test

// todo test Stdin & Stderr

import (
	"bytes"
	"fmt"
	. "github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var (
	cmdActivatorAssert *assert.Assertions
	cmdActivatorBuf    *bytes.Buffer
	activator          CmdActivator
	cmdActivatorTests  = []cmdActivatorTestCase{
		{
			cmd:    "pwd",
			stdout: fmt.Sprintf(" /bin/bash -c 'pwd'\n\n%s\n\n", cwd()),
		},
		{
			cmd:    "pwd",
			wd:     "./unknown-directory",
			stdout: fmt.Sprintf(" /bin/bash -c 'pwd'\n\n\nchdir %s/unknown-directory: no such file or directory\n\nexec: not started\n\n", cwd()),
		},
		{
			cmd:    "pwd",
			wd:     "././/./..///",
			stdout: fmt.Sprintf(" /bin/bash -c 'pwd'\n\n%s\n\n", strings.Replace(cwd(), "/item", "", 1)),
		},
	}
)

type (
	cmdActivatorTestCase struct {
		cmd    string
		wd     string
		stdout string
	}
)

func TestCmdActivator(t *testing.T) {
	cmdActivatorAssert = assert.New(t)
	for _, test := range cmdActivatorTests {
		runCmdActivatorTest(test)
	}
}

func runCmdActivatorTest(t cmdActivatorTestCase) {
	setupCmdActivatorTest()
	validateCmdActivatorTest(t)
}

func setupCmdActivatorTest() {
	cmdActivatorBuf = &bytes.Buffer{}
	activator = CmdActivator{Out: cmdActivatorBuf}
}

func validateCmdActivatorTest(t cmdActivatorTestCase) {
	activator.Activate(&(Item{Cmd: t.cmd, Wd: t.wd}))
	cmdActivatorAssert.Equal(t.stdout, cmdActivatorBuf.String())
}

func cwd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
