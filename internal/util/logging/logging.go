package logging

import (
	"bytes"
	"fmt"
	"log"
	"zootoma/internal/util/misc"
)

type CustomLogger struct {
	buf       *bytes.Buffer
	level     string
	loggers   [4](*log.Logger)
	loggerMap map[string](*log.Logger)
}

var logPriority [4]string = [4]string{"INFO", "WARNING", "ERROR", "DEBUG"}

func NewCustomLogger(level string) (customLogger *CustomLogger) {

	customLogger = new(CustomLogger)

	customLogger.buf = new(bytes.Buffer)

	customLogger.loggerMap = make(map[string](*log.Logger))

	flags := log.Ldate | log.Lmicroseconds | log.Lmsgprefix

	priority, err := misc.IndexOf(logPriority[:], level)
	if err != nil {
		fmt.Println("Error in logger creation, cant find log level")
	}

	for i := 0; i <= priority; {
		_level := logPriority[i]
		customLogger.loggers[i] = log.New(customLogger.buf, _level+": ", flags)
		customLogger.loggerMap[_level] = customLogger.loggers[i]
		i++
	}

	return customLogger

}

func (c CustomLogger) Log(msg string, level string) {
	if logger, ok := c.loggerMap[level]; ok {
		logger.Print(msg)
		fmt.Print(c.buf)
		c.buf.Reset()
	} else {
		fmt.Println("Log level does't exist")
	}
}
