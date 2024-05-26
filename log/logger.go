package log

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type LogEntry struct {
	Application string
	Time        time.Time
	Level       string
	Message     string
}

func LogInfo(message string) {
	logEntry(LogEntry{
		Application: GetApplicationName(),
		Time:        time.Now(),
		Level:       "INFO",
		Message:     message,
	})
}

func LogSuccess(message string) {
	logEntry(LogEntry{
		Application: GetApplicationName(),
		Time:        time.Now(),
		Level:       "SUCCESS",
		Message:     message,
	})
}

func LogError(message string) {
	logEntry(LogEntry{
		Application: GetApplicationName(),
		Time:        time.Now(),
		Level:       "ERROR",
		Message:     message,
	})
}

func logEntry(entry LogEntry) {
	fmt.Printf("%s\n[%s] %s: %s\n", entry.Application, entry.Time.Format("2006-01-02 15:04:05"), entry.Level, entry.Message)
}

func GetApplicationName() string {
	content, err := ioutil.ReadFile("analytics.txt")
	// os.ReadFile("")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
