package utils

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"smartmirror.server/widgets/shared"
)

type Logger struct {
	Info func(formattedString string, args ...interface{})
}

var (
	client = &http.Client{}
)

func NewLogger(widget string) Logger {
	return Logger{
		Info: func(formattedString string, args ...interface{}) {
			timestamp := time.Now()
			logLine := fmt.Sprintf("[%s]: %s\n", widget, fmt.Sprintf(formattedString, args...))

			log.Print(logLine)

			logs := fmt.Appendf(nil,
				`{"streams": [{"stream": {"service_name": "Smartmirror backend", "widget": "%s", "level": "info"}, "values": [["%s", %q]]}]}`,
				widget,
				strconv.FormatInt(timestamp.UnixNano(), 10),
				logLine,
			)

			req, err := http.NewRequest("POST", "https://logs-prod-012.grafana.net/loki/api/v1/push", bytes.NewBuffer(logs))
			if err != nil {
				fmt.Println(err)
				return
			}

			req.Header.Set("Content-Type", "application/json")
			req.SetBasicAuth("1320688", shared.GetGrafanaToken())
			client.Do(req)
		},
	}
}
