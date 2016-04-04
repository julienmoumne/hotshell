package main

import (
	"bytes"
	"fmt"
	"github.com/julienmoumne/hotshell"
)

func main() {

	parser, _ := hotshell.CreateOptionsParser()
	var buf bytes.Buffer
	parser.WriteManPage(&buf)
	fmt.Printf(buf.String())
}
