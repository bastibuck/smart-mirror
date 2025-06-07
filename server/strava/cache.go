package strava

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var cacheKeys = struct {
	Annual       string
	LastActivity string
}{
	Annual:       "annual_stats",
	LastActivity: "last_activity",
}

var stravaCache = cache.New(30*time.Minute, 45*time.Minute)

func getCachedStravaStats() (stravaStats, bool) {
	stats, found := stravaCache.Get(cacheKeys.Annual)

	if !found {
		return stravaStats{}, false
	}

	return stats.(stravaStats), true
}

func setCachedStravaStats(stats stravaStats) {
	stravaCache.Set(cacheKeys.Annual, stats, cache.DefaultExpiration)
}

func getCachedStravaLastActivity() (lastActivity, bool) {
	activity, found := stravaCache.Get(cacheKeys.LastActivity)

	if !found {
		return lastActivity{}, false
	}

	return activity.(lastActivity), true
}

func setCachedStravaLastActivity(activity lastActivity) {
	stravaCache.Set(cacheKeys.LastActivity, activity, cache.DefaultExpiration)
}
