package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/henrylee2cn/faygo"

	"module/demo"
)

func main() {
	// only for pprof
	go pprofServer()

	faygo.SetShutdown(time.Minute, func() error {
		faygo.Debug("Before services close: wait 1s...")
		time.Sleep(1 * time.Second)
		faygo.Debug("Before services close: 1s end!")
		return nil
	}, func() error {
		faygo.Debug("After services are closed: wait 2s...")
		time.Sleep(2 * time.Second)
		faygo.Debug("After services are closed: 2s end!")
		return nil
	})

	faygo.SetUpload("/data/leoGo/upload/", false, false)
	faygo.SetStatic("/data/leoGo/static", false, false)

	{
		demoApp := faygo.New("demo", "1.0")
		demo.Route(demoApp)
		go demoApp.Run()
	}
	select {}
}

// http://localhost:7777/debug/pprof
func pprofServer() {
	http.ListenAndServe("0.0.0.0:7777", nil)
}
