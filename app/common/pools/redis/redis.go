package redis

import (
	. "gin-web/app/config"
	"time"
	"github.com/garyburd/redigo/redis"
	"sync"
)

var (
	rs *redis.Pool
	once sync.Once
)

func Instance() *redis.Pool {
	once.Do(func() {
		rs = &redis.Pool {
			MaxIdle:     Config.Redis.MaxIdle,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				// the redis protocol should probably be made sett-able
				c, err := redis.Dial("tcp", Config.Redis.Address)
				if err != nil {
					return nil, err
				}
				if len(Config.Cache.Password) > 0 {
					if _, err := c.Do("AUTH", Config.Redis.Password); err != nil {
						c.Close()
						return nil, err
					}
				} else {
					// check with PING
					if _, err := c.Do("PING"); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, err
			},
			// custom connection test method
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if _, err := c.Do("PING"); err != nil {
					return err
				}
				return nil
			},
		}
	})
	return rs
}