package garmin

import (
	"time"

	"smartmirror.server/cache"
)

var cacheKeys = struct {
	sevenDaySteps string
}{
	sevenDaySteps: "seven_day_steps",
}

type GarminCache struct {
	cache cache.Cache
}

var garminCache = &GarminCache{
	cache: cache.NewCache(30 * time.Minute),
}

func (s *GarminCache) getSevenDaySteps() (sevenDayStepsModel, bool) {
	return cache.Get[sevenDayStepsModel](s.cache, cacheKeys.sevenDaySteps)
}

func (s *GarminCache) setSevenDaySteps(steps sevenDayStepsModel) {
	s.cache.Set(cacheKeys.sevenDaySteps, steps)
}
