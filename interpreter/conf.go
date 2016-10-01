package interpreter

import "strconv"

const BREADCRUMB_TYPE_KEY = "breadcrumbType"

type Conf struct {
	BreadcrumbType  string
}

func NewConf(value interface{}) Conf {

	val, validType := value.(map[string]interface{})
	if !validType {
		return Conf{}
	}
	return (&confBuilder{val}).build()
}

type confBuilder struct {
	value map[string]interface{}
}

func (builder *confBuilder) build() Conf {
	conf := Conf{}
	conf.BreadcrumbType = builder.getBreadcrumbType()
	return conf
}

func (builder *confBuilder) getBreadcrumbType() string {
	return builder.getScalar(BREADCRUMB_TYPE_KEY)
}

func (builder *confBuilder) getScalar(key string) string {
	value := builder.value[key]
	if value == nil {
		return ""
	}
	if intValue, isInt := value.(int); isInt {
		return strconv.Itoa(intValue)
	}
	if intValue, isInt := value.(int64); isInt {
		return strconv.FormatInt(intValue, 10)
	}
	if floatValue, isFloat := value.(float64); isFloat {
		return strconv.FormatFloat(floatValue, 'f', 0, 64)
	}
	return builder.value[key].(string)
}
