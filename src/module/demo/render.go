package demo

import (
	"github.com/henrylee2cn/faygo"
	// "time"
	"module/demo/dao/zhijin"
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

func Report() faygo.HandlerFunc {
	return func(ctx *faygo.Context) error {
		result := make(faygo.Map)
		result["code"] = 10000
		result["msg"] = "OK"
		//reports, err := zhijin.GetDataBySqlId("mssql/zhijin","select_tbUser", "")
		reports, err := zhijin.GetDataBySqlId("mysql/leotao", "select_tbUser", "")
		if err != nil {
			faygo.Error(err.Error())
			result["code"] = 10001
			result["msg"] = "error in inner system"
		} else {
			result["data"] = reports
		}
		return ctx.JSON(200, result, true)
	}
}
