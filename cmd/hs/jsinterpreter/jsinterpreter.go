package jsinterpreter

import (
	"errors"
	"github.com/ddliu/motto"
	"github.com/robertkrimen/otto"
	"strings"
)

type JsInterpreter struct {
	vm *motto.Motto
}

type JsModule struct {
	Name   string
	Loader motto.ModuleLoader
}

func (i *JsInterpreter) Run(modules []JsModule, js string, resultModule string, resultProperty string) (res interface{}, err error) {
	defer func() {
		err = convertOttoErrorToError(err)
	}()
	i.vm = motto.New()
	i.loadModules(modules)
	if err = i.runJs(js); err != nil {
		return
	}
	return i.retrieveResult(resultModule, resultProperty)
}

func convertOttoErrorToError(err error) error {
	switch err := err.(type) {
	case *otto.Error:
		return errors.New(strings.TrimSpace(err.String()))
	default:
		return err
	}
}

func (i *JsInterpreter) retrieveResult(module string, prop string) (interface{}, error) {
	hsModule, err := i.vm.Require(module, "")
	if err != nil {
		return nil, err
	}
	value, err := hsModule.Object().Get(prop)
	if err != nil {
		return nil, err
	}
	return value.Export()
}

func (i *JsInterpreter) loadModules(modules []JsModule) {
	for _, mod := range modules {
		i.vm.AddModule(mod.Name, mod.Loader)
	}
}

func (i *JsInterpreter) runJs(js string) error {
	_, err := motto.CreateLoaderFromSource(js, "")(i.vm)
	return err
}

func (i *JsInterpreter) compileAndRun(filename string, content []byte) error {
	script, err := i.vm.Otto.Compile(filename, content)
	if err != nil {
		return err
	}
	_, err = i.vm.Otto.Run(script)
	return err
}
