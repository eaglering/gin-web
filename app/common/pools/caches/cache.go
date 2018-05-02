package caches

import (
	"time"
	"gin-web/app/common/libraries/cache/persistence"
	. "gin-web/app/config"
	"sync"
)

var (
	cs persistence.CacheStore
	once sync.Once
)
const (
	Default			 = time.Duration(0)
	Forever			 = time.Duration(-1)
)

func Instance() persistence.CacheStore {
	once.Do(func() {
		switch Config.CacheStore {
		case REDIS:
			cs = persistence.NewRedisCache(Config.Cache.Address, Config.Cache.Password, Config.Cache.Expiration)
		case MEMCACHED:
			cs = persistence.NewMemcachedStore([]string{Config.Cache.Address}, Config.Cache.Expiration)
		default:
			cs = persistence.NewInMemoryStore(Config.Cache.Expiration)
		}
	})
	return cs
}
