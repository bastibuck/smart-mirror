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

func (s *StravaCache) getAnnualStats() (annualStatsModel, bool) {
	return cache.Get[annualStatsModel](s.cache, cacheKeys.Annual)
}

func (s *StravaCache) setAnnualStats(stats annualStatsModel) {
	s.cache.Set(cacheKeys.Annual, stats)
}

func (s *StravaCache) getLastActivity() (lastActivityModel, bool) {
	return cache.Get[lastActivityModel](s.cache, cacheKeys.LastActivity)
}

func (s *StravaCache) setLastActivity(activity lastActivityModel) {
	s.cache.Set(cacheKeys.LastActivity, activity)
}
