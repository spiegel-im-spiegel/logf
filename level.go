package logf

//Level is log level
type Level int

const (
	//TRACE is TRACE level of logging
	TRACE Level = iota
	//DEBUG is DEBUG level of logging
	DEBUG
	//INFO is INFO level of logging
	INFO
	//WARN is WARN level of logging
	WARN
	//ERROR is ERROR level of logging
	ERROR
	//FATAL is FATAL level of logging
	FATAL
)
