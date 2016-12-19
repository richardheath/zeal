package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/richardheath/zeal/config"
)

var logFilePath string

// Initialize Initialize logger for CLI app.
func Initialize() error {
	logFile, err := newLogFile()
	if err != nil {
		return err
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(logFile)

	return nil
}

// GetLogFilePath Get current log file.
func GetLogFilePath() string {
	return logFilePath
}

// Info Print a message to screen and file.
func Info(a ...interface{}) {
	fmt.Fprint(os.Stdout, a...)
	log.Print(a...)
}

// Infof Print a message to screen and file.
func Infof(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, format, a...)
	log.Printf(format, a...)
}

// Debug Print a message to file.
func Debug(a ...interface{}) {
	log.Print(a...)
}

// Debugf Print a message to file.
func Debugf(format string, a ...interface{}) {
	log.Printf(format, a...)
}

func newLogFile() (io.Writer, error) {
	globalConfig := config.GetConfig()
	logsPath := globalConfig.LogsPath

	if _, err := os.Stat(logsPath); os.IsNotExist(err) {
		os.MkdirAll(logsPath, os.ModePerm)
	}

	now := time.Now()
	logFile := fmt.Sprintf("zeal_%d-%d-%d_%d-%d-%d.log", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	logFilePath = path.Join(logsPath, logFile)
	return os.Create(logFilePath)
}
