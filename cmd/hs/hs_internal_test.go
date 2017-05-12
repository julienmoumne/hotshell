package main

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/test"
	"github.com/julienmoumne/hotshell/cmd/hs/versioning"
	"github.com/julienmoumne/hotshell/cmd/term"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"os"
	"testing"
)

func TestBuilder(t *testing.T) { TestingT(t) }

type TestHs struct{}

var _ = Suite(&TestHs{})

func (s *TestHs) TestVersion(c *C) {
	driver := term.TestDriver{Main: func() {
		os.Args = []string{"", "--version"}
		main()
	}}
	actualStdout, _, err := driver.Run()
	c.Check(err, IsNil)
	var version []byte
	version, err = versioning.GetVersion()
	c.Check(err, IsNil)
	c.Check(actualStdout, Equals, fmt.Sprintf("Hotshell version %s\n", version))
}

func (s *TestHs) TestEndToEnd(c *C) {

	err := os.RemoveAll(test.TEST_TMP_DIR)
	if err != nil {
		c.Fatal(err)
	}

	testCases := listTestCases(c)

	for _, testName := range testCases {
		runTest(c, testName)
	}
}

func listTestCases(c *C) []string {

	testCases := make([]string, 0)

	directories, err := ioutil.ReadDir(test.TEST_CASES_DIR)
	if err != nil {
		c.Fatal(err)
	}

	for _, directory := range directories {

		if !directory.IsDir() {
			continue
		}

		dirName := directory.Name()
		files, err := ioutil.ReadDir(test.TEST_CASES_DIR + dirName)
		if err != nil {
			c.Fatal(err)
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

func runTest(c *C, testName string) {
	endToEnd := test.EndToEnd{
		SpecDirectory: testName,
		Testing:       c,
		Exit:          &exit,
		Main:          main,
	}

	if err := endToEnd.Run(); err != nil {
		c.Fatal(err)
	}
}
