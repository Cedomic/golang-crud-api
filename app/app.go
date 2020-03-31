package app

import (
	"fmt"
	"log"
	"net/http"

	"golang-crud-api/app/handlers"
	"golang-crud-api/app/models"
	"golang-crud-api/config"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (app *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Name, config.DB.Password, config.DB.SslMode)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}

	app.DB = models.DBMigrate(db)
	app.Router = mux.NewRouter()
	app.setRouters()
}

func (app *App) setRouters() {
	// Artist routes
	app.Get("/artists", app.handleRequest(handlers.GetAllArtists))
	app.Post("/artists", app.handleRequest(handlers.CreateArtist))
	app.Get("/artists/{artistName}", app.handleRequest(handlers.GetArtist))
	app.Put("/artists/{artistName}", app.handleRequest(handlers.UpdateArtist))
	app.Delete("/artists/{artistName}", app.handleRequest(handlers.DeleteArtist))

	// Song routes
	app.Get("/artists/{artistName}/songs", app.handleRequest(handlers.GetAllSongs))
	app.Post("/artists/{artistName}/songs", app.handleRequest(handlers.CreateSong))
	app.Get("/artists/{artistName}/songs/{songTitle}", app.handleRequest(handlers.GetSong))
	app.Put("/artists/{artistName}/songs/{songTitle}", app.handleRequest(handlers.UpdateSong))
	app.Delete("/artists/{artistName}/songs/{songTitle}", app.handleRequest(handlers.DeleteSong))
}

// Get wraps the router for GET method
func (app *App) Get(path string, f func(responseWriter http.ResponseWriter, request *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (app *App) Post(path string, f func(responseWriter http.ResponseWriter, request *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (app *App) Put(path string, f func(responseWriter http.ResponseWriter, request *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (app *App) Delete(path string, f func(responseWriter http.ResponseWriter, request *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}

type RequestHandlerFunction func(db *gorm.DB, responseWriter http.ResponseWriter, request *http.Request)

func (app *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		handler(app.DB, responseWriter, request)
	}
}
