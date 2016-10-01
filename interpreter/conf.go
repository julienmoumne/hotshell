package interpreter

const BREADCRUMB_TYPE_KEY = "breadcrumbType"
const BREADCRUMB_TYPE_DEFAULT = "horizontal"

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
	
	value := builder.getScalar(BREADCRUMB_TYPE_KEY)
	if value != "horizontal" && value != "vertical" {
		return BREADCRUMB_TYPE_DEFAULT
	}
	return value
}

func (builder *confBuilder) getScalar(key string) string {
	
	value := builder.value[key]
	if value == nil {
		return ""
	}
	if stringValue, isString := value.(string); isString {
		return stringValue
	}
	return ""
}
