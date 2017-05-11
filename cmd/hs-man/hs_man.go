package main

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/options"
)

func main() {
	fmt.Printf((&options.OptionParser{}).CreateManPage())
}
