package strava

import (
	"time"

	"smartmirror.server/cache"
)

var cacheKeys = struct {
	Annual       string
	LastActivity string
}{
	Annual:       "annual_stats",
	LastActivity: "last_activity",
}

type StravaCache struct {
	cache cache.Cache
}

var stravaCache = &StravaCache{
	cache: cache.NewCache(30 * time.Minute),
}

func (s *StravaCache) GetAnnualStats() (stravaStats, bool) {
	return cache.Get[stravaStats](s.cache, cacheKeys.Annual)
}

func (s *StravaCache) SetAnnualStats(stats stravaStats) {
	s.cache.Set(cacheKeys.Annual, stats)
}

func (s *StravaCache) GetLastActivity() (lastActivity, bool) {
	return cache.Get[lastActivity](s.cache, cacheKeys.LastActivity)
}

func (s *StravaCache) SetLastActivity(activity lastActivity) {
	s.cache.Set(cacheKeys.LastActivity, activity)
}
