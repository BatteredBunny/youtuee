package internal

import (
	"github.com/didip/tollbooth/v8"
	"github.com/gin-gonic/gin"
)

func (app *Application) ratelimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(app.ratelimiter, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, app.ratelimiter.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
