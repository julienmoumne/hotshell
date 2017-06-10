package main

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/term"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMan(t *testing.T) {
	driver := term.TestDriver{Main: main}
	actualStdout, _ := driver.Run()
	a := assert.New(t)
	a.Equal(expectedMan, actualStdout)
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
\fB\fB\-f\fR, \fB\-\-file\fR \fIFILE\fR\fP
Specify an alternate definition file (default: ./hs.js)
.TP
\fB\fB\-\-generate-demo\fR\fP
Generate an interactive HTML demo to stdout
.TP
\fB\fB\-\-generate-md\fR\fP
Generate a markdown documentation of the menu to stdout
.TP
\fB\fB\-v\fR, \fB\-\-version\fR\fP
Print version information and quit
`, time.Now().Format("2 January 2006"))
