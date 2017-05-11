package doc

import (
	"bytes"
	"fmt"
	"github.com/julienmoumne/hotshell/item"
	"strings"
	"text/template"
)

type Generator struct {
	Item     *item.Item
	Filename string
	buffer   bytes.Buffer
	itemTmpl *template.Template
	depth    int
}

func (g *Generator) Generate() error {
	if err := g.parseTemplate(); err != nil {
		return err
	}
	g.buffer.WriteString(fmt.Sprintf("%s\n", g.Item.Desc))
	g.generateSubitems(g.Item.Items)
	g.buffer.WriteString("\n\\* *generated using [hotshell](https://github.com/julienmoumne/hotshell)*")
	fmt.Print(g.buffer.String())
	return nil
}

func (g *Generator) parseTemplate() error {
	var err error
	var tmpl = "- {{if .Desc}}{{.Desc}} {{if not .Items}}:{{end}} {{end}}{{if .Cmd}}`{{.Cmd}}`{{end}}\n"
	g.itemTmpl, err = template.New("itemTmpl").Parse(tmpl)
	return err
}

func (g *Generator) generateSubitems(items []*item.Item) {
	for _, i := range items {
		if len(i.Cmd) == 0 && len(i.Items) == 0 {
			continue
		}
		g.buffer.WriteString(strings.Repeat(" ", g.depth*2))
		g.itemTmpl.Execute(&g.buffer, i)
		g.depth++
		g.generateSubitems(i.Items)
		g.depth--
	}
}
