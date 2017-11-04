package redis

import (
	"time"
	"github.com/garyburd/redigo/redis"
	"sync"
	"fmt"
)

var (
	redisPool    *redis.Pool
	redisPoolMu  sync.Mutex
)

type RedisPoolConfig struct {
	host string
	port  int
	password string
	maxIdle int
	maxActive int
	idleTimeout time.Duration
	readTimeout time.Duration
	writeTimeout time.Duration
	isTestOnBorrow bool
}

func NewPool(poolConfig RedisPoolConfig, args ...string) (*redis.Pool, error) {
	if (redisPool == nil) {
		redisPool = &redis.Pool {
			// Other pool configuration not shown in this example.
			Dial: func() (redis.Conn, error) {
				var addr string
				if (poolConfig.port > 0) {
					addr = fmt.Sprintf("%s:%d", poolConfig.host, poolConfig.port)
				} else {
					addr = poolConfig.host
				}
				c, err := redis.Dial("tcp", addr,
					redis.DialReadTimeout(poolConfig.readTimeout),
					redis.DialWriteTimeout(poolConfig.writeTimeout))
				if err != nil {
					return nil, err
				}
				if (poolConfig.password != "") {
					if _, err := c.Do("AUTH", poolConfig.password); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, nil
			},
			MaxIdle: poolConfig.maxIdle,
			MaxActive: poolConfig.maxActive,
			IdleTimeout: poolConfig.idleTimeout,
		}

		if (poolConfig.isTestOnBorrow) {
			redisPool.TestOnBorrow = func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			}
		}
	}

	return redisPool, nil
}