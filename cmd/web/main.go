package main

import (
	"flag"
	"fmt"
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
}

type appConfig struct {
	useCache bool
}

func main() {
	app := application{}
	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.Parse()

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

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
