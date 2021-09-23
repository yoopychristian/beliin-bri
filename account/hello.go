package account

import (
	cfg "beliin-bri/configuration"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Hello : used for health check
func Hello(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := map[string]interface{}{
			"message": "Hello, World!",
			"time":    time.Now(),
		}
		c.JSON(http.StatusOK, res)
	}
}
