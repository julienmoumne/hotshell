package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const TEST_CASES_DIR = "testcases/"

func TestEndToEnd(t *testing.T) {

	err := os.RemoveAll(TEST_TMP_DIR)
	if err != nil {
		t.Fatal(err)
	}

	testCases := listTestCases(t)

	for _, testName := range testCases {
		runTest(t, testName)
	}
}

func listTestCases(t *testing.T) []string {

	testCases := make([]string, 0)

	directories, err := ioutil.ReadDir(TEST_CASES_DIR)
	if err != nil {
		t.Fatal(err)
	}

	for _, directory := range directories {

		if !directory.IsDir() {
			continue
		}

		dirName := directory.Name()
		files, err := ioutil.ReadDir(TEST_CASES_DIR + dirName)
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
	endToEnd := endToEnd{
		SpecDirectory: testName,
		Testing:       t,
	}

	if err := endToEnd.run(); err != nil {
		t.Fatal(err)
	}
}
