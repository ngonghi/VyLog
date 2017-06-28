package VyLog

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type LoggerInterface interface {

	Trace(string, ...interface{})

	Debug(string, ...interface{})

	Info(string, ...interface{})

	Warn(string, ...interface{})

	Error(string, ...interface{})

	Fatal(string, ...interface{})

	Log(int, string, ...interface{})
}