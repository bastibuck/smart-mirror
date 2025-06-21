package kptncook

import (
	"time"

	"smartmirror.server/cache"
)

var cacheKeys = struct {
	dailyRecipes string
}{
	dailyRecipes: "daily_recipes",
}

type KptnCookCache struct {
	cache cache.Cache
}

var kptnCookCache = &KptnCookCache{
	cache: cache.NewCache(1 * time.Hour),
}

func (s *KptnCookCache) getDailyRecipes() (dailyRecipesModel, bool) {
	return cache.Get[dailyRecipesModel](s.cache, cacheKeys.dailyRecipes)
}

func (s *KptnCookCache) setDailyRecipes(recipes dailyRecipesModel) {
	s.cache.Set(cacheKeys.dailyRecipes, recipes)
}
