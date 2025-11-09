package app

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func setupObservability() {
	sentryDsn := getSentryDsn()

	if sentryDsn == "" {
		fmt.Println("Warning: SENTRY_DSN not set, Sentry tracking will be disabled")
		return
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDsn,
		// Adds request headers and IP for users,
		// visit: https://docs.sentry.io/platforms/go/data-management/data-collected/ for more info
		SendDefaultPII: true,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)
}
