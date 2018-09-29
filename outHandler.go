package main

import (
	"bonga/data"
	"strings"

	"github.com/valyala/fasthttp"
)

func outHandler(ctx *fasthttp.RequestCtx) {

	username := strings.Replace(string(ctx.Request.RequestURI()), "/out?id=", "", 1)

	if !data.Instance().Check(username) {
		ctx.NotFound()
		return
	}

	ctx.Redirect("https://bongacams2.com/track?c="+config.GUID+"&csurl=https://bongacams2.com/"+username, 302)
}
