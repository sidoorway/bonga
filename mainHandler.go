package main

import (
	"bonga/data"
	"hash/crc32"

	"github.com/valyala/fasthttp"
)

func mainHandler(ctx *fasthttp.RequestCtx) {
	page := data.MainPage{}

	domain := string(ctx.Request.Header.Peek("Host"))

	seed := int64(crc32.ChecksumIEEE(ctx.Request.RequestURI()))

	page.Lang = "ru"
	page.Title = domain
	page.Domain = domain
	page.Color = color(domain)

	d := data.Instance()

	page.Genders = make([]data.GenderSet, 0)

	for _, gender := range d.Genders() {
		page.Genders = append(page.Genders, data.GenderSet{
			Title: gender,
			Rooms: d.Rooms(gender, config.RoomsOnPage, seed),
		})
	}

	ctx.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)

	tpl.Lookup("main.tpl").Execute(ctx, page)
}
