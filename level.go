package logf

//Level is log level
type Level int

//Values of Level
const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var lavelMap = map[Level]string{
	TRACE: "TRACE",
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

func (lv Level) String() string {
	if s, ok := lavelMap[lv]; ok {
		return s
	}
	return ""
}
