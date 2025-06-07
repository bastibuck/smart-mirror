package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

type GoCache struct {
	c                 *cache.Cache
	defaultExpiration time.Duration
}

func (g *GoCache) Get(key string) (interface{}, bool) {
	return g.c.Get(key)
}

func (g *GoCache) Set(key string, value interface{}) {
	g.c.Set(key, value, g.defaultExpiration)
}

func NewCache(defaultExpiration time.Duration) Cache {
	return &GoCache{
		c:                 cache.New(defaultExpiration, defaultExpiration*2),
		defaultExpiration: defaultExpiration,
	}
}

// returns typed value from cache, or zero value and false if not found
func Get[T any](c Cache, key string) (T, bool) {
	value, found := c.Get(key)
	if !found {
		var zero T
		return zero, false
	}
	return value.(T), true
}
