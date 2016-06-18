package main

import (
	"log"
	"net/http"

	"./handlers"
	"./sqlite"
)

const port = ":8080"

func main() {

	db, err := sqlite.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlite.CreateTables(db, sqlite.QuerysToMigrate); err != nil {
		log.Fatal(err.Error())
	}

	homeHandle := new(handlers.Home)
	podcastHandler := handlers.NewPodcast(db)

	// This handler serve all static files on the app
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public/assets"))))
	http.Handle("/", homeHandle)
	http.HandleFunc("/podcast_create", http.HandlerFunc(podcastHandler.Create))

	log.Println("Start app on port", port)
	http.ListenAndServe(port, nil)
}
