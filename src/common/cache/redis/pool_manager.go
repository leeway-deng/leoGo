package redis

import (
	"time"
	"os/exec"
	"github.com/garyburd/redigo/redis"
	"sync"
)

var (
	redisPool    *redis.Pool
	redisPoolMu  sync.Mutex
)

type RedisPoolConfig struct {
	host string
	port  int
	maxIdle int
	maxActive int
	idleTimeout time.Duration
	isTestOnBorrow bool
}

func NewPool(poolConfig RedisPoolConfig, args ...string) (*redis.Pool, error) {
	s := &Server{
		name: name,
		cmd:  exec.Command(*serverPath, args...),
		done: make(chan struct{}),
	}

	if err != nil {
		s.Stop()
		return nil, err
	}

	return s, nil
}