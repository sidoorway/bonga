package main

import (
	"bonga/data"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

var (
	tpl *template.Template

	config struct {
		Listen      string
		GUID        string
		RoomsOnPage int
		Colors      []string
	}

	err error
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})

	log := logrus.WithField("func", "main")

	log.Info("Loading config")

	cd, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(cd, &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Parsing template")

	var re = regexp.MustCompile(`\/\/[^!\/]+\/(.+)`)

	funcs := template.FuncMap{
		"timeonline": func(d int) string {

			ds := strconv.Itoa(d) + "s"
			dur, err := time.ParseDuration(ds)
			if err != nil {
				return ds
			}
			return dur.String()
		},
		"size": func(s string) string {
			switch s {
			case "Small":
				return "Маленькая"
			case "Medium":
				return "Средняя"
			case "Large":
				return "Большая"
			}
			return s
		},
		"height": func(h string) string {
			return h
		},
		"weight": func(w string) string {
			return w
		},
		"image": func(img string) string {
			return re.ReplaceAllString(img, `/images/$1`)
		},
		"out": func(username string) string {
			return "/out/" + username
		},
		"link": func(username string) string {
			return "/room/" + username
		},
	}

	tpl, err = template.New("html").Funcs(funcs).ParseFiles("tpl/page.tpl", "tpl/main.tpl")
	if err != nil {
		log.Fatal(err)
	}

	d := data.Instance()

	log.Info("Starting rooms reload goroutine")

	go func(d *data.Storage) {
		for {
			time.Sleep(time.Minute * 1)
			d.Reload()
		}
	}(d)

	log.Info("Starting fastHttp server on ", config.Listen)

	err = fasthttp.ListenAndServe(config.Listen, func(ctx *fasthttp.RequestCtx) {

		switch string(ctx.Path()) {
		case "/main":
			mainHandler(ctx)
		case "/page":
			pageHandler(ctx)
		case "/out":
			outHandler(ctx)
		default:
			ctx.NotFound()
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}
