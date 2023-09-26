package customlog

import (
	"log"
	"os"
)

type MyLog struct {
	*log.Logger
	level int
}

const (
	Debug = iota
	Info
	Warn
	Error
	// FATAL
)

func (l *MyLog) SetLevel(level int) {
	l.level = level
}

func (l *MyLog) print(level int, msgArr ...interface{}) {
	var stringLv string
	switch level {
	case 0:
		stringLv = "Debug "
	case 1:
		stringLv = "Info "
	case 2:
		stringLv = "Warn "
	case 3:
		stringLv = "Error "
	default:
		stringLv = "Fatal "
	}
	if l.level <= level {
		msgArr = append([]interface{}{stringLv}, msgArr...)
		l.Logger.Print(msgArr...)
	}
}

func (l *MyLog) Debug(msgArr ...interface{}) {
	l.SetOutput(os.Stdout)
	l.print(Debug, msgArr...)
}
func (l *MyLog) Info(msgArr ...interface{}) {
	os.MkdirAll("logs", os.ModePerm)
	infoFile, err := os.OpenFile(
		"logs/info.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666,
	)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	l.SetOutput(infoFile)
	l.print(Info, msgArr...)
}
func (l *MyLog) Warn(msgArr ...interface{}) {
	os.MkdirAll("logs", os.ModePerm)
	warnFile, err := os.OpenFile(
		"logs/warn.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666,
	)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	l.SetOutput(warnFile)
	l.print(Warn, msgArr...)
}
func (l *MyLog) Error(msgArr ...interface{}) {
	os.MkdirAll("logs", os.ModePerm)
	errFile, err := os.OpenFile(
		"logs/error.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666,
	)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	l.SetOutput(errFile)
	l.print(Error, msgArr...)
}
