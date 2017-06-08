package interpreter

import (
	"github.com/mitchellh/mapstructure"
)

const (
	DescPropName  = "desc"
	CmdPropName   = "cmd"
	KeyPropName   = "key"
	ItemsPropName = "items"
)

type Ast struct {
	Key   string
	Desc  string
	Cmd   string
	Items []Ast
}

func NewAst(value interface{}) (list []Ast, err error) {
	err = mapstructure.WeakDecode(value, &list)
	return
}