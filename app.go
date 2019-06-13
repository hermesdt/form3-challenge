package app

import (
	netHTTP "net/http"

	"github.com/hermesdt/form3-challenge/db"
	"github.com/hermesdt/form3-challenge/http"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	db     *db.Database
	router netHTTP.Handler
}

func NewApp() *App {
	db := db.Connect()
	return &App{
		db:     db,
		router: http.SetupRoutes(db),
	}
}

func (app *App) DB() *mongo.Database {
	return app.db.GetDB()
}

func (app *App) CloseDB() error {
	return app.db.Close()
}

func (app *App) Router() netHTTP.Handler {
	return app.router
}
