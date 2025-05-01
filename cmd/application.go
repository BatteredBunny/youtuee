package cmd

import (
	"embed"
	"flag"
	"html/template"
	"log"
	"time"

	"github.com/BatteredBunny/youtuee/cmd/youtube"
	"github.com/didip/tollbooth/v8"
	"github.com/didip/tollbooth/v8/limiter"
	"github.com/gin-gonic/gin"
)

type Config struct {
	ytdlpBinary string
	Port        int

	BehindReverseProxy bool
}

//go:embed templates
var Templates embed.FS

type Application struct {
	CachedVideoInfo map[string]youtube.VideoInfo

	Router      *gin.Engine
	Ratelimiter *limiter.Limiter

	Conf Config
}

func (app *Application) GetVideoInfo(id string) (v youtube.VideoInfo, err error) {
	cached, exists := app.CachedVideoInfo[id]
	if exists {
		log.Println("Got cached response for ", id)
		return cached, nil
	} else {
		v, err = youtube.GetVideoInfo(app.Conf.ytdlpBinary, id)
		if err == nil {
			app.CachedVideoInfo[id] = v
		}
		return
	}
}

func (app *Application) SetupRatelimiter() {
	app.Ratelimiter = tollbooth.NewLimiter(4, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	if app.Conf.BehindReverseProxy {
		app.Ratelimiter.SetIPLookup(limiter.IPLookup{
			Name:           "X-Forwarded-For",
			IndexFromRight: 0,
		})
	} else {
		app.Ratelimiter.SetIPLookup(limiter.IPLookup{
			Name:           "RemoteAddr",
			IndexFromRight: 0,
		})
	}
}

func NewApplication() Application {
	return Application{
		CachedVideoInfo: make(map[string]youtube.VideoInfo),
	}
}

func (app *Application) SetupRouter() {
	gin.SetMode(gin.ReleaseMode)
	app.Router = gin.Default()
	app.Router.SetHTMLTemplate(template.Must(template.ParseFS(Templates, "templates/video.gohtml")))

	app.Router.Use(app.RatelimiterMiddleware())
	app.Router.GET("/*path", app.PathHandler)
}

func (app *Application) ParseConfig() {
	flag.StringVar(&app.Conf.ytdlpBinary, "yt-dlp", "yt-dlp", "Path to yt-dlp binary")
	flag.IntVar(&app.Conf.Port, "port", 8080, "Port to listen on")
	flag.BoolVar(&app.Conf.BehindReverseProxy, "reverse-proxy", false, "Set true if behind reverse proxy to make the ratelimiter work")
	flag.Parse()
}
