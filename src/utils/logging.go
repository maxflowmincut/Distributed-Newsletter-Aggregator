package utils

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo logs informational messages with a predefined format.
func LogInfo(message string) {
	InfoLogger.Println(message)
}

// LogError logs error messages with a predefined format.
func LogError(message string) {
	ErrorLogger.Println(message)
}