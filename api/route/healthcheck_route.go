package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHealthCheckRouter(group *gin.RouterGroup) {
	group.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})
}