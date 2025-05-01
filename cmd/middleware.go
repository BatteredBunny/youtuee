package cmd

import (
	"github.com/didip/tollbooth/v8"
	"github.com/gin-gonic/gin"
)

func (app *Application) RatelimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(app.Ratelimiter, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, app.Ratelimiter.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
