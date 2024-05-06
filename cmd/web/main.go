package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":4001"

// Will contain application config (dbhost, dbpool etc.)
type application struct{}

func main() {
	// app := application{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})

	fmt.Println("Start web application on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic(err)
	}
}
