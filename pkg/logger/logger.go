package logger

import (
	"fmt"
	"log"
	"os"
)

var l = log.New(os.Stdout, "", log.LstdFlags)

func Fatal(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Fatal] ", part))
	l.Fatalf(fmt.Sprintf("%s\n", format), v...)
}

func Info(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Info] ", part))
	l.Printf(fmt.Sprintf("%s\n", format), v...)
}

func Debug(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Debug] ", part))
	l.Printf(fmt.Sprintf("%s\n", format), v...)
}

func Panic(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Panic] ", part))
	l.Panicf(fmt.Sprintf("%s\n", format), v...)
}

func Error(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Error] ", part))
	l.Printf(fmt.Sprintf("%s\n", format), v...)
}

func Warning(part, format string, v ...interface{}) {
	l.SetPrefix(fmt.Sprintf("[%s] [Warning] ", part))
	l.Printf(fmt.Sprintf("%s\n", format), v...)
}
