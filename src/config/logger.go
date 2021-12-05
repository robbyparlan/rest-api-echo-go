package config

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
)

var LogWriteFile io.Writer
var LogDb io.Writer

func init() {
	if !strings.HasSuffix(os.Args[0], ".test") && (len(os.Args) <= 1 || !strings.Contains(os.Args[1], "test.run")) {
		// LogWriteFile ...
		LogWriteFile = io.MultiWriter(os.Stdout, LogFile())

		// LogDb ...
		LogDb = io.MultiWriter(os.Stdout, DBLogFile())

		log.SetHeader(`[${time_rfc3339}] : [INFO]`)
		log.SetOutput(LogWriteFile)
	}

}

// prefixFilename ...
func prefixFilename() string {
	today := time.Now().Format("2006-01-02")
	return today + ".log"
}

// LogFile ...
func LogFile() *os.File {
	filename := "logs/" + prefixFilename()
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	return f
}

// DBLogFile ...
func DBLogFile() *os.File {
	filename := "logs/DB-" + prefixFilename()
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}

	return f
}
