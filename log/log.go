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
	Info
	Warn
	Debug
)

func (s Level) String() string {
	switch s {
	case Error: return "ERROR"
	case Warn: return "WARN"
	case Info: return "INFO"
	case Debug: return "DEBUG"
	default: return ""
	}
}

func (s Level) FormatString() string {
	switch s {
	case Error: return "ERROR"
	case Warn: return "WARN "
	case Info: return "INFO "
	case Debug: return "DEBUG"
	default: return "     "
	}
}
