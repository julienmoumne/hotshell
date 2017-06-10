package main

// todo find a way to test all examples

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/test"
	"github.com/julienmoumne/hotshell/cmd/hs/versioning"
	"github.com/julienmoumne/hotshell/cmd/term"
	"github.com/mitchellh/go-homedir"
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
	// todo find a better way to shadow the settings file
	if renameSettingsfile() {
		defer restoreSettingsfile()
	}
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

func settingsFilePaths() (string, string) {
	path, err := homedir.Expand("~/.hsrc.js")
	if err != nil {
		return "", ""
	}
	return path, fmt.Sprintf("%s.backup", path)
}

func renameSettingsfile() bool {
	path, tmp := settingsFilePaths()
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	if err := os.Rename(path, tmp); err != nil {
		panic(err)
	}
	return true
}

func restoreSettingsfile() {
	path, tmp := settingsFilePaths()
	if err := os.Rename(tmp, path); err != nil {
		panic(err)
	}
}
