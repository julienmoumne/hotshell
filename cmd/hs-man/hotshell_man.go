package main

import (
	"bytes"
	"fmt"
	"github.com/julienmoumne/hs"
)

func main() {

	parser, _ := hs.CreateOptionsParser()
	var buf bytes.Buffer
	parser.WriteManPage(&buf)
	fmt.Printf(buf.String())
}
