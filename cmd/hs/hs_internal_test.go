package main

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/testutil"
	"github.com/julienmoumne/hotshell/versioning"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"os"
	"testing"
)

func TestBuilder(t *testing.T) { TestingT(t) }

type TestHs struct{}

var _ = Suite(&TestHs{})

const TEST_CASES_DIR = "testcases/"

func (s *TestHs) TestVersion(c *C) {
	driver := testutil.Driver{Main: func() {
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

	err := os.RemoveAll(TEST_TMP_DIR)
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

	directories, err := ioutil.ReadDir(TEST_CASES_DIR)
	if err != nil {
		c.Fatal(err)
	}

	for _, directory := range directories {

		if !directory.IsDir() {
			continue
		}

		dirName := directory.Name()
		files, err := ioutil.ReadDir(TEST_CASES_DIR + dirName)
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
	endToEnd := endToEnd{
		SpecDirectory: testName,
		Testing:       c,
	}

	if err := endToEnd.run(); err != nil {
		c.Fatal(err)
	}
}
