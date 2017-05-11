//go:generate go-bindata -nometadata -ignore \.go$ -pkg demo ./
package demo

import (
	"os"
	"strings"
	"text/template"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
)

type Generator struct {
	Item     *item.Item
	Filename string
	mainTmpl string
	itemTmpl string
	Css      string
	Js       string
}

func (g *Generator) Generate() error {
	if err := g.loadAssets(); err != nil {
		return err
	}

	funcMap := template.FuncMap{
		"Sanitize": func(str string) string {
			str = strings.Replace(str, "\n", "<br/>", -1)
			return strings.Replace(str, "'", "\\'", -1)
		},
	}

	tmpl, err := template.New("demo").Funcs(funcMap).Parse(g.mainTmpl + g.itemTmpl)
	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, g)
}

func (g *Generator) loadAssets() error {

	tmpl, err := Asset("demo.tmpl")
	if err != nil {
		return err
	}
	g.mainTmpl = string(tmpl)

	itemTmpl, err := Asset("item.tmpl")
	if err != nil {
		return err
	}
	g.itemTmpl = string(itemTmpl)

	css, err := Asset("demo.css")
	if err != nil {
		return err
	}
	g.Css = string(css)

	js, err := Asset("demo.js")
	if err != nil {
		return err
	}
	g.Js = string(js)

	return nil
}
