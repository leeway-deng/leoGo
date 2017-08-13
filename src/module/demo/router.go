package demo

import (
	"github.com/henrylee2cn/faygo"
	tgMiddleware "github.com/henrylee2cn/faygo/ext/middleware"
)

// Register the route in a tree style
func Route(frame *faygo.Framework) {
	frame.
	Filter(DoFilter).
		Route(
		//refer: https://github.com/henrylee2cn/faydoc/blob/master/zh/05.01.md
		frame.NewNamedAPI("index", "GET", "/index", Index()),
		frame.NewGroup("home",
			frame.NewNamedGET("html", "render", &Render{}),
		),
		frame.NewNamedAPI("zhijin", "GET", "api/zhijin/report", Report()),
		//frame.NewStatic("/sf", faygo.JoinStatic("")),
	).Use(tgMiddleware.CrossOrigin)
}
