package dlog

import (
	"log"
	"os"
)

var debugEnabled = os.Getenv("DEBUG") != ""

// Debugf logs a formatted debug message if debugging is enabled.
func Debugf(format string, args ...any) {
	if debugEnabled {
		log.Printf("[DEBUG] "+format, args...)
	}
}

// EnableDebug enables debug logging.
func EnableDebug() {
	debugEnabled = true
}

// DisableDebug disables debug logging.
func DisableDebug() {
	debugEnabled = false
}
