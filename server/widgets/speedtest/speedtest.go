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
