package internal

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	embed "github.com/BatteredBunny/youtuee"
	"github.com/BatteredBunny/youtuee/internal/yt"
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

type Application struct {
	cachedVideoInfo map[string]yt.VideoInfo
	router          *gin.Engine
	ratelimiter     *limiter.Limiter

	youtubeApi *youtube.Service

	conf Config
}

func (app *Application) getVideoInfo(id string) (v yt.VideoInfo, err error) {
	cached, exists := app.cachedVideoInfo[id]
	if exists {
		log.Println("Got cached response for ", id)
		return cached, nil
	} else {
		if app.youtubeApi != nil {
			log.Println("Getting video info from yt api")
			v, err = yt.YtApiGetVideoInfo(app.youtubeApi, id)
			if err != nil {
				log.Println("Failed to fetch video:", err)
			} else {
				app.cachedVideoInfo[id] = v
				return
			}
		}

		// Fallback option
		v, err = yt.YtDlpGetVideoInfo(app.conf.ytdlpBinary, id)
		if err == nil {
			app.cachedVideoInfo[id] = v
		}

		return
	}
}

func (app *Application) setupRatelimiter() {
	app.ratelimiter = tollbooth.NewLimiter(4, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	if app.conf.BehindReverseProxy {
		app.ratelimiter.SetIPLookup(limiter.IPLookup{
			Name:           "X-Forwarded-For",
			IndexFromRight: 0,
		})
	} else {
		app.ratelimiter.SetIPLookup(limiter.IPLookup{
			Name:           "RemoteAddr",
			IndexFromRight: 0,
		})
	}
}

func NewApplication() Application {
	return Application{
		cachedVideoInfo: make(map[string]yt.VideoInfo),
	}
}

func (app *Application) setupRouter() {
	app.router = gin.Default()
	app.router.SetHTMLTemplate(template.Must(template.ParseFS(embed.Templates, "templates/video.gohtml")))

	app.router.Use(app.ratelimiterMiddleware())
	app.router.GET("/*path", app.pathHandler)
}

const YT_API_KEY_ENV = "YT_API"

func (app *Application) setupYtApi() {
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

func (app *Application) parseConfig() {
	flag.StringVar(&app.conf.ytdlpBinary, "yt-dlp", "yt-dlp", "Path to yt-dlp binary")
	flag.IntVar(&app.conf.Port, "port", 8080, "Port to listen on")
	flag.BoolVar(&app.conf.BehindReverseProxy, "reverse-proxy", false, "Set true if behind reverse proxy to make the ratelimiter work")
	flag.Parse()
}

func (app *Application) Start() {
	app.parseConfig()
	app.setupYtApi()

	app.setupRatelimiter()

	app.setupRouter()

	log.Printf("Starting api on port %d\n", app.conf.Port)
	log.Fatal(app.router.Run(fmt.Sprintf(":%d", app.conf.Port)))
}
