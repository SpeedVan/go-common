package omap

// Map todo
type Map interface {
	Get(string) interface{}
	GetWithDefault(string, interface{}) interface{}
	Set(string, interface{}) Map
	Keys() []string
	Values() []interface{}
	ForEach(func(string, interface{}))
	F_map(func(string, interface{}) (string, interface{})) Map
	ToSourceMap() map[string]interface{}
	ToStringMap() map[string]string
}
