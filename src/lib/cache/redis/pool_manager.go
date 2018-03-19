package redis

import (
	"time"
	"github.com/garyburd/redigo/redis"
	"sync"
	"fmt"
)

var (
	redisPool      *redis.Pool
	redisPoolMutex sync.Mutex
	once           sync.Once
)

type RedisPoolConfig struct {
	Host           *string
	Port           *int
	Password       *string
	MaxIdle        int
	MaxActive      int
	IdleTimeout    time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IsTestOnBorrow bool
}

func CreatePool(poolConfig RedisPoolConfig, args ...string) (*redis.Pool, error) {
	//if (redisPool == nil) { // Singleton way1
	//	redisPoolMutex.Lock()
	//	defer redisPoolMutex.Unlock()
	once.Do(func() { // Singleton way2
		redisPool = &redis.Pool{
			// Other pool configuration not shown in this example.
			Dial: func() (redis.Conn, error) {
				var addr string
				if (*poolConfig.Port > 0) {
					addr = fmt.Sprintf("%s:%d", *poolConfig.Host, *poolConfig.Port)
				} else {
					addr = fmt.Sprintf("%s", *poolConfig.Host)
				}
				c, err := redis.Dial("tcp", addr,
					redis.DialReadTimeout(poolConfig.ReadTimeout),
					redis.DialWriteTimeout(poolConfig.WriteTimeout))
				if err != nil {
					return nil, err
				}
				if (poolConfig.Password != nil) {
					if _, err := c.Do("AUTH", fmt.Sprintf("%s", *poolConfig.Password)); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, nil
			},
			MaxIdle:     poolConfig.MaxIdle,
			MaxActive:   poolConfig.MaxActive,
			IdleTimeout: poolConfig.IdleTimeout,
		}

		if (poolConfig.IsTestOnBorrow) {
			redisPool.TestOnBorrow = func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			}
		}
	})

	return redisPool, nil
}