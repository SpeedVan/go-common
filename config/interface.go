package config

// Config todo
type Config interface {
	WithPrefix(string) Config
	Get(string) string
}
