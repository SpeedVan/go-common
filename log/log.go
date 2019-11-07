package log

// Logger todo
type Logger interface {
	Log(Level, string)
	Error(string)
	Warn(string)
	Info(string)
	Debug(string)
	LogF(Level, string, ...interface{})
	ErrorF(string, ...interface{})
	WarnF(string, ...interface{})
	InfoF(string, ...interface{})
	DebugF(string, ...interface{})
	SetLevel(Level) Logger
}

// Level todo
type Level int

const (
	_ Level = iota
	Debug
	Warn
	Info
	Error
)

func (s Level) String() string {
	switch s {
	case Error:
		return "ERROR"
	case Warn:
		return "WARN"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	default:
		return ""
	}
}

// FormatString same length to petty format print
func (s Level) FormatString() string {
	switch s {
	case Error:
		return "ERROR"
	case Warn:
		return "WARN "
	case Info:
		return "INFO "
	case Debug:
		return "DEBUG"
	default:
		return "     "
	}
}
