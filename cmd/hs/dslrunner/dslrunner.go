//go:generate go-bindata -nometadata -ignore \.go$ -pkg dslrunner ./
package dslrunner

import (
	"errors"
	"fmt"
	"github.com/ddliu/motto"
	_ "github.com/ddliu/motto/underscore"
	"github.com/julienmoumne/hotshell/cmd/hs/formatter"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	. "github.com/julienmoumne/hotshell/cmd/hs/jsinterpreter"
	"github.com/robertkrimen/otto"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

type DslRunner struct {
}

func (i *DslRunner) Run(dsl string) (it *item.Item, err error) {
	defer func() {
		if err == nil {
			return
		}
		err = errors.New(fmt.Sprintf("Error while reading the menu definition\n%s", err.Error()))
	}()
	hsMod, err := i.createHotshellModule()
	if err != nil {
		return
	}
	res, err := (&JsInterpreter{}).Run([]JsModule{hsMod}, dsl, "hotshell", "items")
	if err != nil {
		return
	}
	return (&mapper{}).mapp(res)
}

func (i *DslRunner) createHotshellModule() (JsModule, error) {
	js, err := Asset("dslrunner.js")
	if err != nil {
		return JsModule{}, err
	}
	return JsModule{
		Name: "hotshell",
		Factory: func(jsInt *JsInterpreter) motto.ModuleLoader {
			return func(vm *motto.Motto) (otto.Value, error) {
				module, err := motto.CreateLoaderFromSource(string(js), "")(vm)
				if err != nil {
					return otto.Value{}, err
				}
				if err := module.Object().Set("exec", i.nativeExec(jsInt)); err != nil {
					return otto.Value{}, err
				}
				return module, nil
			}
		},
	}, nil
}

func (d *DslRunner) nativeExec(jsInt *JsInterpreter) func(call otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		cmd := exec.Command("bash", "-c", call.Argument(0).String())
		cmd.Dir = getWD(jsInt)
		outBytes, execErr := cmd.CombinedOutput()
		out := strings.TrimSpace(string(outBytes))
		handleErrorInNative(call.Otto, cmd, execErr, out)
		outOtto, errConv := call.Otto.ToValue(out)
		if errConv != nil {
			panic(errConv) // not sure what to do in this case
		}
		return outOtto
	}
}

func getWD(jsInt *JsInterpreter) string {
	currItem, err := jsInt.RetrieveValueFromModule("hotshell", "current")
	if err != nil {
		panic(err)
	}
	wdRaw, err := currItem.Object().Get("wd")
	if err != nil {
		panic(err)
	}
	wd, _ := wdRaw.Export()
	osCwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path.Clean(fmt.Sprintf("%s/%s", osCwd, wd.(string)))
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
