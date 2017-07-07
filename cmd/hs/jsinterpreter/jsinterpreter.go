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
	Name    string
	Factory func(jsInt *JsInterpreter) motto.ModuleLoader
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
	var resRaw otto.Value
	resRaw, err = i.RetrieveValueFromModule(resultModule, resultProperty)
	if err != nil {
		return
	}
	return resRaw.Export()
}

func convertOttoErrorToError(err error) error {
	switch err := err.(type) {
	case *otto.Error:
		return errors.New(strings.TrimSpace(err.String()))
	default:
		return err
	}
}

func (i *JsInterpreter) RetrieveValueFromModule(module string, prop string) (otto.Value, error) {
	hsModule, err := i.vm.Require(module, "")
	if err != nil {
		return otto.Value{}, err
	}
	return hsModule.Object().Get(prop)
}

func (i *JsInterpreter) loadModules(modules []JsModule) {
	for _, mod := range modules {
		i.vm.AddModule(mod.Name, mod.Factory(i))
	}
}

func (i *JsInterpreter) runJs(js string) error {
	js = "require.main = module\n" + js
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
