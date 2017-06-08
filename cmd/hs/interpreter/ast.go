package interpreter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
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
	items, err := getGenericMapArray(value)
	if err != nil {
		return
	}
	for _, i := range items {
		builder := &astBuilder{i, Ast{}}
		if err = builder.build(); err != nil {
			return
		}
		list = append(list, builder.ast)
	}
	return
}

type astBuilder struct {
	value map[string]interface{}
	ast   Ast
}

func (a *astBuilder) build() (err error) {
	a.ast.Key, err = a.getKey()
	if err != nil {
		return
	}
	a.ast.Desc, err = a.getDesc()
	if err != nil {
		return
	}
	a.ast.Cmd, err = a.getCmd()
	if err != nil {
		return
	}
	a.ast.Items, err = NewAst(a.value[ItemsPropName])
	return
}

func getGenericMapArray(candidate interface{}) (m []map[string]interface{}, err error) {
	if candidate == nil {
		return
	}
	m, typeIsMapArray := candidate.([]map[string]interface{})
	if !typeIsMapArray {
		array, typeIsArray := candidate.([]interface{})
		if !typeIsArray || len(array) > 0 {
			err = errors.New(fmt.Sprintf("type is '%s', expected '[]map[string]interface{}", reflect.TypeOf(candidate)))
		}
	}
	return
}

func (a *astBuilder) getDesc() (string, error) {
	return a.getScalar(DescPropName)
}

func (a *astBuilder) getKey() (string, error) {
	return a.getScalar(KeyPropName)
}

func (a *astBuilder) getCmd() (string, error) {
	return a.getScalar(CmdPropName)
}

func (a *astBuilder) getScalar(key string) (string, error) {
	value := a.value[key]
	if value == nil {
		return "", nil
	}
	if intValue, isInt := value.(int); isInt {
		return strconv.Itoa(intValue), nil
	}
	if intValue, isInt := value.(int64); isInt {
		return strconv.FormatInt(intValue, 10), nil
	}
	if floatValue, isFloat := value.(float64); isFloat {
		return strconv.FormatFloat(floatValue, 'f', 0, 64), nil
	}
	str, strConv := value.(string)
	if !strConv {
		return "", errors.New(fmt.Sprintf("error while parsing property '%s', type '%s' is not convertible to string", key, reflect.TypeOf(value)))
	}
	return str, nil
}
