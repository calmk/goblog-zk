package main

import (
	"database/sql"
	"goblogCalmk/app/http/middlewares"
	"goblogCalmk/bootstrap"
	"goblogCalmk/pkg/database"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()
var db *sql.DB

func main() {

	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	http.ListenAndServe(":8000", middlewares.RemoveTrailingSlash(router))
}
