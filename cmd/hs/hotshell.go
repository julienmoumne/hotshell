package main

import (
	"fmt"
	"github.com/julienmoumne/hotshell"
	"github.com/robertkrimen/otto"
	"os"
)

func main() {

	err := (&hotshell.Bootstrap{}).Boot()

	handleError(err)
}

func handleError(err error) {

	if err == nil {
		return
	}

	switch err := err.(type) {
	case *otto.Error:
		fmt.Fprint(os.Stderr, err.String())
	default:
		fmt.Fprintln(os.Stderr, err)
	}

	exit(1)
}

// var is required to change the definition during tests
var exit = func(code int) {
	os.Exit(code)
}
