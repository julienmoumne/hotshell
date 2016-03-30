package main

import (
	"fmt"
	dmp "github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

const TEST_TMP_DIR = "tmp/failed-cases/"

type endToEnd struct {
	SpecDirectory string
	Testing       *testing.T
}

func (e *endToEnd) path(file string) string {
	return fmt.Sprintf("%s%s/%s", TEST_CASES_DIR, e.SpecDirectory, file)
}

func (e *endToEnd) run() error {

	input, err := e.getStdinSpec()
	if err != nil {
		return err
	}

	driver := driver{
		menuDefinition: e.path("hs.js"),
		input:          input,
	}

	actualStdout, actualStderr, err := driver.run()
	if err != nil {
		return err
	}

	e.diff(actualStdout, "stdout")
	e.diff(actualStderr, "stderr")

	return nil
}

func (e *endToEnd) getStdinSpec() ([]byte, error) {

	input, err := e.readFile("stdin")
	if err != nil {
		return nil, err
	}

	inputBytes := make([]byte, 0)

	if len(input) == 0 {
		return inputBytes, nil
	}

	var commentRegexp = regexp.MustCompile(`[\s\/].*$`) // everything on the right that starts with an empty space or a forward slash

	keys := strings.Split(input, "\n")
	for _, key := range keys {

		intValue, err := strconv.Atoi(commentRegexp.ReplaceAllLiteralString(key, ""))
		if err != nil {
			return nil, err
		}
		inputBytes = append(inputBytes, byte(intValue))
	}

	return inputBytes, nil
}

func (e *endToEnd) diff(actual string, postfix string) error {

	expected, err := e.readFile("expected." + postfix)
	if err != nil {
		return err
	}

	if expected == actual {
		return nil
	}

	return e.reportFailedTest(actual, expected, postfix)
}

func (e *endToEnd) reportFailedTest(actual string, expected string, postfix string) error {

	testDir := fmt.Sprintf("%s%s", TEST_TMP_DIR, e.SpecDirectory)
	if err := os.MkdirAll(testDir, 0755); err != nil {
		return err
	}

	diffFilename := fmt.Sprintf("%s/diff.%s.html", testDir, postfix)
	actualFilename := fmt.Sprintf("%s/actual.%s", testDir, postfix)

	dmp := dmp.New()
	diff := dmp.DiffMain(expected, actual, true)
	if err := writeFile(diffFilename, dmp.DiffPrettyHtml(diff)); err != nil {
		return err
	}

	if err := writeFile(actualFilename, actual); err != nil {
		return err
	}

	e.Testing.Errorf("%s produced unexpected output, see %s for test outputs", postfix, testDir)

	return nil
}

func (e *endToEnd) readFile(file string) (string, error) {

	bytes, err := ioutil.ReadFile(e.path(file))
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func writeFile(path, data string) error {

	if err := ioutil.WriteFile(path, []byte(data), 0644); err != nil {
		return err
	}
	return nil
}
