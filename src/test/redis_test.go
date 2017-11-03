package test

import (
	"flag"
	"errors"
	"fmt"
	"testing"
	"os"
	"io/ioutil"
	"sync"
	"github.com/garyburd/redigo/redis"
	"os/exec"
	"strconv"
	"time"
)

var (
	redisHost     = flag.String("redis-host", "redis-host", "Path to redis server binary")
	redisPort = flag.Int("redis-port", 16379, "Beginning of port range for test servers")
	redisLogName  = flag.String("redis-log", "", "Write Redis server logs to `filename`")
	redisLog      = ioutil.Discard //the io.writer without return

	redisPoolMu  sync.Mutex
	redisPool    *redis.Pool
	defaultErr error
)

// DialDefaultServer starts the test server if not already started and dials a
// connection to the server.
func DialDefaultServer() (redis.Conn, error) {
	if err := startDefaultServer(); err != nil {
		return nil, err
	}
	c, err := redis.Dial("tcp", fmt.Sprintf(":%d", *serverBasePort),
		redis.DialReadTimeout(1*time.Second),
		redis.DialWriteTimeout(1*time.Second))
	if err != nil {
		return nil, err
	}
	c.Do("FLUSHDB")
	return c, nil
}

func TestMain(m *testing.M) {
	os.Exit(func() int {
		flag.Parse()

		var f *os.File
		if *serverLogName != "" {
			var err error
			f, err = os.OpenFile(*serverLogName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening redis-log: %v\n", err)
				return 1
			}
			defer f.Close()
			serverLog = f
		}

		defer stopDefaultServer()

		return m.Run()
	}())
}

// startDefaultServer starts the default server if not already running.
func startDefaultServer() error {
	defaultServerMu.Lock()
	defer defaultServerMu.Unlock()
	if defaultServer != nil || defaultServerErr != nil {
		return defaultServerErr
	}
	defaultServer, defaultServerErr = NewServer(
		"default",
		"--port", strconv.Itoa(*serverBasePort),
		"--save", "",
		"--appendonly", "no")
	return defaultServerErr
}

// stopDefaultServer stops the server created by DialDefaultServer.
func stopDefaultServer() {
	defaultServerMu.Lock()
	defer defaultServerMu.Unlock()
	if defaultServer != nil {
		defaultServer.Stop()
		defaultServer = nil
	}
}
