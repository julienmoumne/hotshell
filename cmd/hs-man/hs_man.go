package main

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/optionparser"
)

func main() {
	fmt.Printf((&optionparser.OptionParser{}).CreateManPage())
}
