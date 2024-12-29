package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (app *Application) PathHandler(c *gin.Context) {
	if c.Param("path") == "/" {
		c.String(http.StatusOK, "https://github.com/BatteredBunny/youtuee")
	} else {
		id := strings.TrimPrefix(c.Param("path"), "/")

		v, err := app.GetVideoInfo(id)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.HTML(http.StatusOK, "video.gohtml", v)
	}
}
