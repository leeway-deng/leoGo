package test

import (
	"flag"
	"fmt"
	"io/ioutil"
	"lib/cache/redis"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"
)

/*
测试所有的文件 go test，将对当前目录下的所有*_test.go文件进行编译并自动运行测试。
测试某个文件使用”-file”参数。go test –file *.go 。例如：go test -file mysql_test.go，"-file"参数不是必须的，可以省略，如果你输入go test b_test.go也会得到一样的效果。
测试某个方法 go test -run='Test_xxx'
"-v" 参数 go test -v ... 表示无论用例是否测试通过都会显示结果，不加"-v"表示只显示未通过的用例结果
进行所有go文件的benchmark测试 go test -bench=".*" 或 go test . -bench=".*"
对某个go文件进行benchmark测试 go test mysql_b_test.go -bench=".*"
*/

var (
	redisHost    = flag.String("host", "redis-host", "Path to redis server binary")
	redisPort    = flag.Int("port", 16379, "Beginning of port range for test servers")
	redisLogName = flag.String("log", "", "Write Redis server logs to `filename`")
	redisLog     = ioutil.Discard //the io.writer without return

	redisPoolMutex  sync.Mutex
	redisPoolConfig = &redis.RedisPoolConfig{
		Host:           redisHost,
		Port:           redisPort,
		Password:       nil,
		MaxIdle:        3,
		MaxActive:      10,
		IdleTimeout:    240 * time.Second,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		IsTestOnBorrow: true,
	}
	redisPool, err = redis.CreatePool(*redisPoolConfig)

	testCommands = []struct {
		args     []interface{}
		expected interface{}
	}{
		{
			[]interface{}{"PING"},
			"PONG",
		},
		{
			[]interface{}{"SET", "foo", "bar"},
			"OK",
		},
		{
			[]interface{}{"GET", "foo"},
			[]byte("bar"),
		},
		{
			[]interface{}{"GET", "nokey"},
			nil,
		},
		{
			[]interface{}{"MGET", "nokey", "foo"},
			[]interface{}{nil, []byte("bar")},
		},
		{
			[]interface{}{"INCR", "mycounter"},
			int64(1),
		},
		{
			[]interface{}{"LPUSH", "mylist", "foo"},
			int64(1),
		},
		{
			[]interface{}{"LPUSH", "mylist", "bar"},
			int64(2),
		},
		{
			[]interface{}{"LRANGE", "mylist", 0, -1},
			[]interface{}{[]byte("bar"), []byte("foo")},
		},
		{
			[]interface{}{"MULTI"},
			"OK",
		},
		{
			[]interface{}{"LRANGE", "mylist", 0, -1},
			"QUEUED",
		},
		{
			[]interface{}{"PING"},
			"QUEUED",
		},
		{
			[]interface{}{"EXEC"},
			[]interface{}{
				[]interface{}{[]byte("bar"), []byte("foo")},
				"PONG",
			},
		},
	}
)

func TestDoCommands(t *testing.T) {
	c := redisPool.Get()
	defer c.Close()

	for _, cmd := range testCommands {
		actual, err := c.Do(cmd.args[0].(string), cmd.args[1:]...)
		if err != nil {
			t.Errorf("Do(%v) returned error %v", cmd.args, err)
			continue
		}
		if !reflect.DeepEqual(actual, cmd.expected) {
			t.Errorf("Do(%v) = %v, want %v", cmd.args, actual, cmd.expected)
		} else {
			t.Logf("Do(%v) = %v", cmd.args, actual)
		}
	}
}

//go test redis_test.go -host "127.0.0.1" -port "36379" -log "/Users/apple/Projects/go/leoGo/src/test/redis.log"
func TestMain(m *testing.M) {
	os.Exit(func() int {
		flag.Parse()
		fmt.Printf("redisHost=%s", *redisHost)
		var f *os.File
		if *redisLogName != "" {
			var err error
			f, err = os.OpenFile(*redisLogName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening redis-log: %v\n", err)
				return 1
			}
			defer f.Close()
			redisLog = f
		}

		defer destroy()

		return m.Run()
	}())
}

func destroy() {
	if nil != redisPool {
		redisPool.Close()
	}
}
