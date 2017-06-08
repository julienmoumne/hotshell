//go:generate go-bindata -nometadata -ignore \.go$ -pkg interpreter ./
package interpreter

import (
	"fmt"
	"github.com/ddliu/motto"
	_ "github.com/ddliu/motto/underscore"
	"github.com/julienmoumne/hotshell/cmd/hs/formatter"
	"github.com/robertkrimen/otto"
	"os/exec"
	"regexp"
	"strings"
)

type Interpreter struct {
	dsl []byte
	vm  *motto.Motto
}

func (i *Interpreter) Interpret(dsl []byte) (interface{}, error) {
	i.dsl = dsl
	i.vm = motto.New()
	if err := i.loadHotshellModule(); err != nil {
		return nil, err
	}
	if err := i.exec(); err != nil {
		return nil, err
	}
	return i.retrieveResult()
}

func (i *Interpreter) retrieveResult() (interface{}, error) {
	hsModule, err := i.vm.Require("hotshell", "")
	if err != nil {
		return nil, err
	}
	value, err := hsModule.Object().Get("items")
	if err != nil {
		return nil, err
	}
	return value.Export()
}

func (i *Interpreter) loadHotshellModule() error {
	js, err := Asset("dslrunner.js")
	if err != nil {
		return err
	}
	i.vm.AddModule("hotshell", func(vm *motto.Motto) (otto.Value, error) {
		module, err := motto.CreateLoaderFromSource(string(js), "")(i.vm)
		if err != nil {
			return otto.Value{}, err
		}
		if err := module.Object().Set("exec", nativeExec); err != nil {
			return otto.Value{}, err
		}
		return module, nil
	})
	return nil
}

func (i *Interpreter) exec() error {
	_, err := motto.CreateLoaderFromSource(string(i.dsl), "")(i.vm)
	return err
}

func (i *Interpreter) compileAndRun(filename string, content []byte) error {
	script, err := i.vm.Otto.Compile(filename, content)
	if err != nil {
		return err
	}
	_, err = i.vm.Otto.Run(script)
	return err
}

func nativeExec(call otto.FunctionCall) otto.Value {
	cmd := exec.Command("bash", "-c", call.Argument(0).String())
	outBytes, execErr := cmd.CombinedOutput()
	out := strings.TrimSpace(string(outBytes))
	handleErrorInNative(call.Otto, cmd, execErr, out)
	outOtto, errConv := call.Otto.ToValue(out)
	if errConv != nil {
		panic(errConv) // not sure what to do in this case
	}
	return outOtto
}

func handleErrorInNative(vm *otto.Otto, cmd *exec.Cmd, err error, out string) {
	if err == nil {
		return
	}
	exception, callErr := vm.Call(
		"new Error",
		nil,
		fmt.Sprintf(
			"\"%s\" failed with %s \"%s\"",
			formatter.FormatCommand(cmd),
			err,
			formatStderr(out),
		),
	)
	if callErr == nil {
		panic(exception) // https://github.com/robertkrimen/otto/issues/17
	}
	panic(callErr) // not sure what to do in this case
}

func formatStderr(stderr string) string {
	r, rexExpErr := regexp.Compile("\\s+")
	if rexExpErr != nil {
		panic(rexExpErr) // not sure what to do in this case
	}

	return string(r.ReplaceAll([]byte(strings.TrimSpace(stderr)), []byte(" ")))
}
