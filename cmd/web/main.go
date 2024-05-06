package main

import (
	"flag"
	"fmt"
	"go-breeders/configuration"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4001"

// Will contain application settings (dbhost, dbpool etc.)
type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *configuration.Application // this is our singleton
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}
	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders_design_systems?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

	// Get DB
	db, err := initMySQLDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}

	// app.Models = *models.New(db) // hooking up the models with the database connection
	app.App = configuration.New(db)

	server := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	// This uses go's default mux router
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hi")
	// })

	fmt.Println("Start web application on port", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
