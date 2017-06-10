//go:generate go-bindata -nometadata -ignore \.go$ -pkg settings ./
package settings

import (
	"errors"
	"fmt"
	"github.com/blang/vfs"
	"github.com/ddliu/motto"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	. "github.com/julienmoumne/hotshell/cmd/hs/jsinterpreter"
	"github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/robertkrimen/otto"
)

type Keys struct {
	Back   byte
	Bash   byte
	Repeat byte
	Reload byte
}

func Defaults() Settings {
	return Settings{
		Keys: Keys{
			Back:   item.KeyCodes.Space,
			Bash:   item.KeyCodes.Tab,
			Repeat: item.KeyCodes.Return,
			Reload: item.KeyCodes.Backspace,
		},
	}
}

type Settings struct {
	Keys Keys
}

type Loader struct {
	fs  vfs.Filesystem
	js  []byte
	cfg Settings
}

func (l *Loader) Load(fs vfs.Filesystem) (s Settings, err error) {
	defer func() {
		if err == nil {
			return
		}
		err = errors.New(fmt.Sprintf("Error while reading ~/.hsrc.js\n%s", err.Error()))
	}()
	l.fs = fs
	if err := l.retrieveFile(); err != nil {
		return Defaults(), nil
	}
	if err = l.interpretJs(); err != nil {
		return
	}
	return l.cfg, nil
}

func (l *Loader) retrieveFile() error {
	path, err := homedir.Expand("~/.hsrc.js")
	if err != nil {
		return err
	}
	l.js, err = vfs.ReadFile(l.fs, path)
	return err
}

func (l *Loader) interpretJs() error {
	cfgMod, err := l.createSettingsModule()
	if err != nil {
		return err
	}
	res, err := (&JsInterpreter{}).Run([]JsModule{cfgMod}, string(l.js), "hotshell-settings", "settings")
	if err != nil {
		return err
	}
	l.cfg = Defaults()
	return mapstructure.Decode(res, &l.cfg)
}

func (l *Loader) createSettingsModule() (JsModule, error) {
	js, err := Asset("settings.js")
	if err != nil {
		return JsModule{}, err
	}
	return JsModule{
		Name: "hotshell-settings",
		Loader: func(vm *motto.Motto) (otto.Value, error) {
			module, err := motto.CreateLoaderFromSource(string(js), "")(vm)
			if err != nil {
				return otto.Value{}, err
			}
			if err := module.Object().Set("keys", item.KeyCodes); err != nil {
				return otto.Value{}, err
			}
			return module, nil
		},
	}, nil
}
