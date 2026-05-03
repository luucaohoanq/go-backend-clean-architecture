package main

import (
	"time"

	route "github.com/amitshekhariitbhu/go-backend-clean-architecture/api/route"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	ginEngine := gin.Default()

	//if err := ginEngine.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
	//	app.Logger.Error("failed to set trusted proxies", "err", err)
	//}

	route.Setup(env, timeout, db, ginEngine, app.Logger)

	err := ginEngine.Run(env.ServerAddress)
	if err != nil {
		return
	}
}
