package speedtest

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

type speedtestResponse struct {
	Download float64 `json:"download"`
	Upload   float64 `json:"upload"`
	Ping     int64   `json:"ping"`
}

func runSpeedtest() (speedtestResponse, error) {
	var speedtestClient = speedtest.New()

	serverList, _ := speedtestClient.FetchServers()
	targets := *serverList.Available()

	if len(targets) == 0 {
		return speedtestResponse{}, fmt.Errorf("no speedtest servers found")
	}

	// get random server from list
	targetServer := targets[rand.Intn(len(targets))]

	logger.Info("Using server: [%s] %s by %s", targetServer.ID, targetServer.Name, targetServer.Sponsor)

	targetServer.PingTest(nil)
	targetServer.DownloadTest()
	targetServer.UploadTest()

	return speedtestResponse{
		Download: targetServer.DLSpeed.Mbps(),
		Upload:   targetServer.ULSpeed.Mbps(),
		Ping:     targetServer.Latency.Milliseconds(),
	}, nil
}

type LastResult struct {
	SecondsAgo int64   `json:"secondsAgo"`
	Download   float64 `json:"download"`
	Upload     float64 `json:"upload"`
	Ping       int64   `json:"ping"`
}

func getSpeedTestResults(cutoffTime time.Duration) (lastResults []LastResult) {
	now := time.Now()
	cutoff := now.Add(cutoffTime)

	recent := make([]LastResult, 0)
	for _, entry := range speedtestHistory {
		if entry.Time.After(cutoff) {
			recent = append([]LastResult{{
				SecondsAgo: int64(now.Sub(entry.Time).Seconds()),
				Download:   entry.Download,
				Upload:     entry.Upload,
				Ping:       entry.Ping,
			}}, recent...) // chronological order
		}

		if entry.Time.Before(cutoff) {
			break // history is in reverse chronological order, so we can stop checking once we hit an old entry
		}
	}

	return recent
}
