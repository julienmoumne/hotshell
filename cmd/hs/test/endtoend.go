package test

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/term"
	dmp "github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

const (
	TestCasesDir = "test/cases/"
	TestTmpDir   = "test/tmp/"
)

type EndToEnd struct {
	SpecDirectory string
	Testing       *testing.T
	Exit          *func(int)
	Main          func()
}

func (e *EndToEnd) path(file string) string {
	return fmt.Sprintf("%s%s/%s", TestCasesDir, e.SpecDirectory, file)
}

func (e *EndToEnd) Run() {
	e.executeScenario()
	e.generateDoc("md", "md")
	e.generateDoc("demo", "html")
}

func (e *EndToEnd) createMain(args []string) func() {
	return func() {
		os.Args = args
		// todo test exitCode
		var exitCode int
		(*e.Exit) = func(code int) {
			exitCode = code
		}
		e.Main()
	}
}

func (e *EndToEnd) generateDoc(cmd string, ext string) {
	e.executeTest([]string{"", fmt.Sprintf("--generate-%s", cmd)}, ext)
}

func (e *EndToEnd) executeScenario() {
	e.executeTest([]string{""}, "scenario")
}

func (e *EndToEnd) executeTest(args []string, stdout string) {
	d := term.TestDriver{
		Input: e.getStdinSpec(),
		Main:  e.createMain(args),
		Cwd:   e.path(""),
	}
	actualStdout, actualStderr := d.Run()
	e.diff(actualStdout, stdout)
	e.diff(actualStderr, "stderr")
}

func (e *EndToEnd) getStdinSpec() (inputBytes []byte) {
	input := e.readFile("stdin")
	if len(input) == 0 {
		return
	}

	var commentRegexp = regexp.MustCompile(`[\s\/].*$`) // everything on the right that starts with an empty space or a forward slash
	keys := strings.Split(input, "\n")
	for _, key := range keys {
		for _, byteChar := range strings.Split(commentRegexp.ReplaceAllLiteralString(key, ""), ",") {
			intValue, err := strconv.Atoi(byteChar)
			if err != nil {
				e.Testing.Fatal(err)
			}
			inputBytes = append(inputBytes, byte(intValue))
		}
	}
	return
}

func (e *EndToEnd) diff(actual string, postfix string) {
	expected := e.readFile("expected." + postfix)
	if expected == actual {
		return
	}
	e.reportFailedTest(actual, expected, postfix)
}

func (e *EndToEnd) reportFailedTest(actual string, expected string, postfix string) {
	testDir := fmt.Sprintf("%s%s", TestTmpDir, e.SpecDirectory)
	if err := os.MkdirAll(testDir, 0755); err != nil {
		e.Testing.Fatal(err)
	}

	diffFilename := fmt.Sprintf("%s/diff.%s.html", testDir, postfix)
	actualFilename := fmt.Sprintf("%s/actual.%s", testDir, postfix)

	dmp := dmp.New()
	diff := dmp.DiffMain(expected, actual, true)
	e.writeFile(diffFilename, dmp.DiffPrettyHtml(diff))
	e.writeFile(actualFilename, actual)
	e.Testing.Errorf("%s produced unexpected output, see %s for test outputs", postfix, testDir)
}

func (e *EndToEnd) readFile(file string) string {

	bytes, err := ioutil.ReadFile(e.path(file))
	if err != nil {
		e.Testing.Fatal(err)
	}
	return string(bytes)
}

func (e *EndToEnd) writeFile(path, data string) {
	if err := ioutil.WriteFile(path, []byte(data), 0644); err != nil {
		e.Testing.Fatal(err)
	}
}
