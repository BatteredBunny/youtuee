package cmd

import (
	"context"
	"embed"
	"flag"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/BatteredBunny/youtuee/cmd/yt"
	"github.com/didip/tollbooth/v8"
	"github.com/didip/tollbooth/v8/limiter"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Config struct {
	ytdlpBinary   string // Path to yt-dlp binary
	Port          int    // HTTP Port
	YoutubeApiKey *string

	BehindReverseProxy bool // Makes the ratelimiter work behind a reverse proxy
}

//go:embed templates
var Templates embed.FS

type Application struct {
	CachedVideoInfo map[string]yt.VideoInfo

	Router      *gin.Engine
	Ratelimiter *limiter.Limiter

	youtubeApi *youtube.Service

	Conf Config
}

func (app *Application) GetVideoInfo(id string) (v yt.VideoInfo, err error) {
	cached, exists := app.CachedVideoInfo[id]
	if exists {
		log.Println("Got cached response for ", id)
		return cached, nil
	} else {
		if app.youtubeApi == nil {
			v, err = yt.YtDlpGetVideoInfo(app.Conf.ytdlpBinary, id)
			if err == nil {
				app.CachedVideoInfo[id] = v
			}
		} else {
			log.Println("Getting video info from yt api")
			v, err = yt.YtApiGetVideoInfo(app.youtubeApi, id)
			if err == nil {
				app.CachedVideoInfo[id] = v
			}
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
		CachedVideoInfo: make(map[string]yt.VideoInfo),
	}
}

func (app *Application) SetupRouter() {
	gin.SetMode(gin.ReleaseMode)
	app.Router = gin.Default()
	app.Router.SetHTMLTemplate(template.Must(template.ParseFS(Templates, "templates/video.gohtml")))

	app.Router.Use(app.RatelimiterMiddleware())
	app.Router.GET("/*path", app.PathHandler)
}

const YT_API_KEY_ENV = "YT_API"

func (app *Application) SetupYtApi() {
	key, existing := os.LookupEnv(YT_API_KEY_ENV)
	if existing {
		log.Printf("Env %s found, enabling youtube api usage", YT_API_KEY_ENV)

		ctx := context.Background()
		var err error
		app.youtubeApi, err = youtube.NewService(ctx, option.WithAPIKey(key))
		if err != nil {
			app.youtubeApi = nil
			log.Println(err, "disabling youtube api")
		}
	}
}

func (app *Application) ParseConfig() {
	flag.StringVar(&app.Conf.ytdlpBinary, "yt-dlp", "yt-dlp", "Path to yt-dlp binary")
	flag.IntVar(&app.Conf.Port, "port", 8080, "Port to listen on")
	flag.BoolVar(&app.Conf.BehindReverseProxy, "reverse-proxy", false, "Set true if behind reverse proxy to make the ratelimiter work")
	flag.Parse()
}
