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
			logger.Info("'%s' scheduled to run every %s", taskName, time.Duration(interval).String())

			ticker := time.NewTicker(interval)

			go func() {
				for range ticker.C {
					now := time.Now()

					logger.Info("'%s' scheduled task started", taskName)
					task()
					logger.Info("'%s' scheduled task finished (took %s)", taskName, time.Since(now).String())
				}
			}()
		},
	}
}
