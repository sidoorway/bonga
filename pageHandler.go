package main

import (
	"bonga/data"
	"hash/crc32"
	"strings"

	"github.com/valyala/fasthttp"
)

func pageHandler(ctx *fasthttp.RequestCtx) {
	page := data.ChatPage{}

	domain := string(ctx.Request.Host())
	username := strings.Replace(string(ctx.Request.RequestURI()), "/page?id=", "", 1)

	seed := int64(crc32.ChecksumIEEE(ctx.Request.RequestURI()))

	page.Lang = "ru"

	page.Domain = domain
	page.Color = "lightblue"

	d := data.Instance()

	page.Room, err = d.Room(username)
	if err != nil {
		ctx.NotFound()
		return
	}

	page.Title = page.Room.DisplayName + " - " + page.Room.Gender + " - " + domain

	page.Rooms = d.Rooms(page.Room.Gender, config.RoomsOnPage, seed)

	ctx.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)

	tpl.Lookup("page.tpl").Execute(ctx, page)
}
