package server

import (
	"fmt"
	"log"
	"os"
)

var l = log.New(os.Stdout, "", log.LstdFlags)

func LogFatal(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Fatal] ", part))
	l.Fatalf(fmt.Sprintf("%s\n", format), v...)
}

func LogInfo(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Info] ", part))
	l.Printf(fmt.Sprintf("%s\n", format), v...)
}

func LogDebug(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Debug] ", part))
	l.Printf(fmt.Sprintf("%s\n", format), v...)
}

func LogPanic(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Panic] ", part))
	l.Panicf(fmt.Sprintf("%s\n", format), v...)
}

func LogError(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Error] ", part))
	l.Printf(fmt.Sprintf("%s\n", format), v...)
}

func LogWarning(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Warning] ", part))
	l.Printf(fmt.Sprintf("%s\n", format), v...)
}

func LogRecover(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Recover] ", part))
	l.Printf(fmt.Sprintf("%s\n", format), v...)
}
