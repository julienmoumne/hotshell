//go:generate go-bindata -nometadata -ignore \.go$ -pkg generator ./
package generator

import (
	"github.com/julienmoumne/hotshell/cmd/hs/item"
	"os"
	"strings"
	"text/template"
)

type Demo struct {
	Item     *item.Item
	Filename string
	mainTmpl string
	itemTmpl string
	Css      string
	Js       string
}

func (g *Demo) Generate() error {
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

func (g *Demo) loadAssets() error {

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
