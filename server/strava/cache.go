package strava

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var statsCache = cache.New(30*time.Minute, 45*time.Minute)

const statsCacheKey = "strava_stats"

func getCachedStravaStats() (stravaStats, bool) {
	stats, found := statsCache.Get(statsCacheKey)

	if !found {
		return stravaStats{}, false
	}

	return stats.(stravaStats), true
}

func setCachedStravaStats(stats stravaStats) {
	statsCache.Set(statsCacheKey, stats, cache.DefaultExpiration)
}
