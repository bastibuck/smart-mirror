package speedtest

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
)

type speedtestResponse struct {
	Download float64 `json:"download"`
	Upload   float64 `json:"upload"`
	Ping     int64   `json:"ping"`
}

// TODO: run this on a regular basis like every 15 minutes and store results somewhere
func runSpeedtest() (speedtestResponse, error) {
	var speedtestClient = speedtest.New()

	serverList, _ := speedtestClient.FetchServers()
	targets, _ := serverList.FindServer([]int{})

	if len(targets) == 0 {
		logger.Info("No speedtest servers found")
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
	SecondsAgo int64   `json:"seconds_ago"`
	Download   float64 `json:"download"`
	Upload     float64 `json:"upload"`
	Ping       int64   `json:"ping"`
}

func getSpeedTestResults(hours int64) (lastResults []LastResult) {
	// TODO: read from a persistant place like a file, DB or smth

	return []LastResult{
		{SecondsAgo: 3600 * 3, Download: 50.5, Upload: 10.2, Ping: 20},
		{SecondsAgo: 3600 * 2, Download: 52.3, Upload: 11.0, Ping: 18},
		{SecondsAgo: 3600 * 1, Download: 48.7, Upload: 9.8, Ping: 22},
		{SecondsAgo: 0, Download: 51.0, Upload: 10.5, Ping: 19},
	}

}
