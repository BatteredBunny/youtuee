package cmd

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (app *Application) PathHandler(c *gin.Context) {
	path := c.Param("path")
	if path == "/" {
		c.String(http.StatusOK, "https://github.com/BatteredBunny/youtuee")
	} else {
		// YouTube video IDs are always 11 characters {a-zA-Z0-9-_}
		id := strings.TrimPrefix(path, "/")

		if !VerifyPath(id) {
			c.Redirect(http.StatusTemporaryRedirect, "https://www.youtube.com")
			return
		}

		// Cut away junk after the ID, helps with caching
		id = id[:11] // VerifyPath made sure this is 11 or longer

		v, err := app.GetVideoInfo(id)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.HTML(http.StatusOK, "video.gohtml", v)
	}
}
