package config

type Config interface {
	Get(string) string
}
