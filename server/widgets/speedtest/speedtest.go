package speedtest

import (
	"fmt"
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
	targets, _ := serverList.FindServer([]int{})

	if len(targets) == 0 {
		return speedtestResponse{}, fmt.Errorf("no speedtest servers found")
	}

	// run against first server only
	targetServer := targets[0]

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
			recent = append(recent, LastResult{
				SecondsAgo: int64(now.Sub(entry.Time).Seconds()),
				Download:   entry.Download,
				Upload:     entry.Upload,
				Ping:       entry.Ping,
			})
		}
	}

	return recent
}
