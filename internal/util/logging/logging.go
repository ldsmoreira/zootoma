package logging

import (
	"bytes"
	"fmt"
	"log"
)

const (
	ERROR int = iota
	WARNING
	INFO
	DEBUG
)

type CustomLogger struct {
	buf    *bytes.Buffer
	logger *log.Logger
	level  int
}

type ICustomLogger interface {
	Error(msg string)
	Warning(msg string)
	Info(msg string)
	Debug(msg string)
}

func (cl CustomLogger) log(msg string) {
	cl.logger.Print(msg)
	fmt.Print(cl.buf)
	cl.buf.Reset()
}

func (cl CustomLogger) Error(msg string) {
	fmt.Println(cl.level)
	if cl.level >= ERROR {
		cl.logger.SetPrefix("Error: ")
		cl.log(msg)
	}
}

func (cl CustomLogger) Warning(msg string) {
	fmt.Println(cl.level)
	if cl.level >= WARNING {
		cl.logger.SetPrefix("Warning: ")
		cl.log(msg)
	}
}

func (cl CustomLogger) Info(msg string) {
	fmt.Println(cl.level)
	if cl.level >= INFO {
		cl.logger.SetPrefix("INFO: ")
		cl.log(msg)
	}
}

func (cl CustomLogger) Debug(msg string) {
	fmt.Println(cl.level)
	if DEBUG <= cl.level {
		cl.logger.SetPrefix("DEBUG: ")
		cl.log(msg)
	}
}

func NewCustomLogger(level int) (customLogger *CustomLogger) {

	customLogger = new(CustomLogger)

	customLogger.buf = new(bytes.Buffer)

	customLogger.level = level

	flags := log.Ldate | log.Lmicroseconds | log.Lmsgprefix

	customLogger.logger = log.New(customLogger.buf, "", flags)

	return customLogger

}
