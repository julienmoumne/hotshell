package main

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
	actualStdout, _, err := driver.Run()
	a.Nil(err)
	var version []byte
	version, err = versioning.GetVersion()
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

func listTestCases(t *testing.T) []string {

	testCases := make([]string, 0)

	directories, err := ioutil.ReadDir(test.TestCasesDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, directory := range directories {

		if !directory.IsDir() {
			continue
		}

		dirName := directory.Name()
		files, err := ioutil.ReadDir(test.TestCasesDir + dirName)
		if err != nil {
			t.Fatal(err)
		}

		for _, file := range files {

			testCases = append(
				testCases,
				fmt.Sprintf("%s/%s", dirName, file.Name()),
			)
		}
	}

	return testCases
}

func runTest(t *testing.T, testName string) {
	endToEnd := test.EndToEnd{
		SpecDirectory: testName,
		Testing:       t,
		Exit:          &exit,
		Main:          main,
	}

	if err := endToEnd.Run(); err != nil {
		t.Fatal(err)
	}
}
