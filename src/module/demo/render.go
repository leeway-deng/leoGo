package demo

import (
	"github.com/henrylee2cn/faygo"
	// "time"
)

type Render struct {
	Title     string   `param:"<in:query> <desc:标题，不为0> <nonzero>"`
	Paragraph []string `param:"<in:query> <desc:参数，长度10个字符> <name:p> <len: 1:10> <regexp: ^[\\w]*$>"`
}

func (r *Render) Serve(ctx *faygo.Context) error {
	return ctx.Render(200, faygo.JoinStatic("demo/render.html"), faygo.Map{
		"title": r.Title,
		"p":     r.Paragraph,
	})
}

func init() {
	faygo.RenderVar("__RES__", "/static/demo/res")
}

func Index() faygo.HandlerFunc {
	return func(ctx *faygo.Context) error {
		// time.Sleep(10e9)
		return ctx.Render(200, faygo.JoinStatic("demo/index.html"), faygo.Map{
			"TITLE":   "io-ok",
			"VERSION": faygo.VERSION,
			"CONTENT": "Welcome To io-ok",
			"AUTHOR":  "Leeway.Deng",
		})
	}
}