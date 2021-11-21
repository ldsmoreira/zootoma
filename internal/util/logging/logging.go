package logging

import (
	"bytes"
	"fmt"
	"log"
)

type CustomLogger struct {
	buf           *bytes.Buffer
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	loggerMap     map[string](*log.Logger)
}

func NewCustomLogger() (customLogger *CustomLogger) {

	customLogger = new(CustomLogger)

	customLogger.buf = new(bytes.Buffer)

	flags := log.Ldate | log.Lmicroseconds | log.Lmsgprefix

	customLogger.debugLogger = log.New(customLogger.buf, "DEBUG: ", flags)
	customLogger.infoLogger = log.New(customLogger.buf, "INFO: ", flags)
	customLogger.warningLogger = log.New(customLogger.buf, "WARNING: ", flags)
	customLogger.errorLogger = log.New(customLogger.buf, "ERROR: ", flags)

	customLogger.loggerMap = make(map[string](*log.Logger))

	customLogger.loggerMap["DEBUG"] = customLogger.debugLogger
	customLogger.loggerMap["INFO"] = customLogger.infoLogger
	customLogger.loggerMap["WARNING"] = customLogger.warningLogger
	customLogger.loggerMap["ERROR"] = customLogger.errorLogger

	return customLogger

}

func (c CustomLogger) Log(msg string, level string) {
	logger := c.loggerMap[level]
	logger.Print(msg)
	fmt.Print(c.buf)
	c.buf.Reset()
}
