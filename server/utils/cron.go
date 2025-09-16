package utils

import (
	"fmt"
	"time"
)

type Cron struct {
	Schedule func(taskName string, interval time.Duration, task func())
}

func NewCron(widget string) Cron {
	logger := NewLogger(fmt.Sprintf("CRON-%s", widget))

	return Cron{
		Schedule: func(taskName string, interval time.Duration, task func()) {
			logger.Info("Scheduled '%s' to run every %s", taskName, time.Duration(interval).String())

			ticker := time.NewTicker(interval)

			go func() {
				for range ticker.C {
					logger.Info("Running scheduled task '%s'", taskName)
					task()
				}
			}()
		},
	}
}
