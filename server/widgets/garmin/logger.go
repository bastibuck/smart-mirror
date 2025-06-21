package garmin

import (
	"fmt"
	"log"
	"time"
)

func logger(format string, args ...interface{}) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("[GARMIN] %s: %s\n", ts, fmt.Sprintf(format, args...))
}
