//go:generate go-bindata -nometadata -ignore \.go$ -pkg interpreter ./
package interpreter

import (
	"fmt"
	"github.com/ddliu/motto"
	"github.com/julienmoumne/hotshell/formatter"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
	"os/exec"
	"regexp"
	"strings"
)

type Interpreter struct {
	Filename string
	Dsl      []byte
	vm       *motto.Motto
}

func (i *Interpreter) Interpret() (Conf, []Ast, error) {
	i.vm = motto.New()

	if err := i.registerNatives(); err != nil {
		return Conf{}, nil, err
	}

	if err := i.exec(); err != nil {
		return Conf{}, nil, err
	}

	return i.buildResult()
}

func (i *Interpreter) buildResult() (Conf, []Ast, error) {
	
	// Options
	
	options, err := i.vm.Get("options")
	if err != nil {
		return Conf{}, nil, err
	}

	opts, err := options.Export()
	if err != nil {
		return Conf{}, nil, err
	}
	
	// AST
	
	value, err := i.vm.Get("items")
	if err != nil {
		return Conf{}, nil, err
	}

	val, err := value.Export()
	if err != nil {
		return Conf{}, nil, err
	}

	return NewConf(opts), NewAst(val), nil
}

func (i *Interpreter) exec() error {
	// could be cached when reloading menu definition
	dslrunner := "dslrunner.js"
	js, err := Asset(dslrunner)
	if err != nil {
		return err
	}

	if err := i.compileAndRun(dslrunner, js); err != nil {
		return err
	}

	_, err = motto.CreateLoaderFromSource(string(i.Dsl), "")(i.vm)
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

func (i *Interpreter) registerNatives() error {
	return i.vm.Set("exec", nativeExec)
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
