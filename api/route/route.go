package route

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/api/middleware"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/bootstrap"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, router *gin.Engine, logger *slog.Logger) {
	
	router.GET("/favicon.ico", func(c *gin.Context) {
        c.AbortWithStatus(http.StatusNoContent)
    })

	// Global Middleware
	router.Use(middleware.LoggingMiddleware(logger))
	router.Use(gin.Recovery())

	publicRouter := router.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	NewHealthCheckRouter(publicRouter)

	protectedRouter := router.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)
}
