package log

// Logger todo
type Logger interface {
	Log(Level, string, ...interface{})
	Error(string, ...interface{})
	Warn(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
}

// Level todo
type Level int

const (
	_ Level = iota
	Error
	Warn
	Info
	Debug
)
