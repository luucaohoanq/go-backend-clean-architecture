package bootstrap

import (
	"log/slog"
	"os"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
)

type Application struct {
	Env    *Env
	Mongo  mongo.Client
	Logger *slog.Logger
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)

	// Cách 1: Dùng logger mặc định (Text format)
	// app.Logger = slog.Default()

	// Cách 2: Dùng JSON format (Cực chuẩn cho dân Java chuyển sang làm Microservices)
	app.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
