package main

import (
	"github.com/julienmoumne/hotshell/cmd/testutil"
	. "gopkg.in/check.v1"
	"testing"
	"fmt"
	"time"
)

func TestBuilder(t *testing.T) { TestingT(t) }

type TestHsMan struct{}

var _ = Suite(&TestHsMan{})

func (s *TestHsMan) TestMan(c *C) {
	driver := testutil.Driver{Main: main}
	actualStdout, _, err := driver.Run()
	c.Check(err, IsNil)
	c.Check(actualStdout, Equals, expectedMan)
}

var expectedMan = fmt.Sprintf(`.TH hs 1 "%s"
.SH NAME
hs \- Interactive single keystroke menus for the shell
.SH SYNOPSIS
\fBhs\fP [-f <arg>...] [Options]
.SH DESCRIPTION
Hotshell is a command-line application to efficiently recall and share commands.
.SH OPTIONS
.TP
\fB\fB\-\-default\fR\fP
Load the default menu to get to know Hotshell
.TP
\fB\fB\-f\fR, \fB\-\-file\fR \fIFILE|DIR\fR\fP
Specify an alternate definition file (defaults: ./hs.js, ~/.hs/hs.js)
.TP
\fB\fB\-\-chdir\fR\fP
Set the working directory to the location of the menu definition
.TP
\fB\fB\-\-generate-demo\fR\fP
Generate an interactive HTML demo to stdout
.TP
\fB\fB\-\-generate-doc\fR\fP
Generate a markdown documentation of the menu to stdout
.TP
\fB\fB\-v\fR, \fB\-\-version\fR\fP
Print version information and quit
`, time.Now().Format("02 January 2006"))
