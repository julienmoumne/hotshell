//go:generate go-bindata -nometadata -ignore \.go$ -pkg dslrunner ./
package dslrunner

import (
	"errors"
	"fmt"
	"github.com/ddliu/motto"
	_ "github.com/ddliu/motto/underscore"
	"github.com/julienmoumne/hotshell/cmd/hs/formatter"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"github.com/robertkrimen/otto"
	"os/exec"
	"regexp"
	"strings"
)

type DslRunner struct {
	menuDef   string
	vm        *motto.Motto
	rawResult interface{}
	item      *item.Item
}

func (i *DslRunner) Run(dsl string) (it *item.Item, err error) {
	defer func() {
		it = i.item
		err = convertOttoErrorToError(err)
	}()
	i.menuDef = dsl
	i.vm = motto.New()
	if err = i.loadHotshellModule(); err != nil {
		return
	}
	if err = i.runMenuDef(); err != nil {
		return
	}
	if err = i.retrieveResult(); err != nil {
		return
	}
	if err = i.mapToItems(); err != nil {
		return
	}
	return
}

func convertOttoErrorToError(err error) error {
	switch err := err.(type) {
	case *otto.Error:
		return errors.New(strings.TrimSpace(err.String()))
	default:
		return err
	}
}

func (i *DslRunner) mapToItems() (err error) {
	i.item, err = (&mapper{}).mapp(i.rawResult)
	return
}

func (i *DslRunner) retrieveResult() error {
	hsModule, err := i.vm.Require("hotshell", "")
	if err != nil {
		return err
	}
	value, err := hsModule.Object().Get("items")
	if err != nil {
		return err
	}
	i.rawResult, err = value.Export()
	return err
}

func (i *DslRunner) loadHotshellModule() error {
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

func (i *DslRunner) runMenuDef() error {
	_, err := motto.CreateLoaderFromSource(i.menuDef, "")(i.vm)
	return err
}

func (i *DslRunner) compileAndRun(filename string, content []byte) error {
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
