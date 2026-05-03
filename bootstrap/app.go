package bootstrap

import (
	"log/slog"
	"os"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"github.com/lmittmann/tint"
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
	// app.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Cách 3: Tint + Handler
	handler := tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: "15:04:05",
	})

	app.Logger = slog.New(handler)

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
