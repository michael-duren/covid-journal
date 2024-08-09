package logging

import (
	"log"
	"os"
)

func getLogLevel(level string) LogLevel {
	switch level {
	case "TRACE":
		return Trace
	case "DEBUG":
		return Debug
	case "INFO":
		return Information
	case "WARN":
		return Warning
	case "FATAL":
		return Fatal
	}
	return Information
}

var level = os.Getenv("LOG_LEVEL")

func NewDefaultLogger() Logger {
	logLevel := getLogLevel(level)
	flags := log.Lmsgprefix | log.Ltime
	return &DefaultLogger{
		minLevel: logLevel,
		loggers: map[LogLevel]*log.Logger{
			Trace:       log.New(os.Stdout, "TRACE ", flags),
			Debug:       log.New(os.Stdout, "DEBUG ", flags),
			Information: log.New(os.Stdout, "INFO ", flags),
			Warning:     log.New(os.Stdout, "WARN ", flags),
			Fatal:       log.New(os.Stdout, "FATAL ", flags),
		},
		triggerPanic: true,
	}
}
