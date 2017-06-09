package main

// todo find a way to test all examples

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/test"
	"github.com/julienmoumne/hotshell/cmd/hs/versioning"
	"github.com/julienmoumne/hotshell/cmd/term"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestVersion(t *testing.T) {
	a := assert.New(t)
	driver := term.TestDriver{Main: func() {
		os.Args = []string{"", "--version"}
		main()
	}}
	actualStdout, _ := driver.Run()
	var version []byte
	version, err := versioning.GetVersion()
	a.Nil(err)
	a.Equal(fmt.Sprintf("Hotshell version %s\n", version), actualStdout)
}

func TestEndToEnd(t *testing.T) {
	err := os.RemoveAll(test.TestTmpDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, testName := range listTestCases(t) {
		runTest(t, testName)
	}
}

func listTestCases(t *testing.T) (cases []string) {
	directories, err := ioutil.ReadDir(test.TestCasesDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, directory := range directories {
		cases = append(cases, directory.Name())
	}
	return
}

func runTest(t *testing.T, testName string) {
	(&test.EndToEnd{
		SpecDirectory: testName,
		Testing:       t,
		Exit:          &exit,
		Main:          main,
	}).Run()
}
