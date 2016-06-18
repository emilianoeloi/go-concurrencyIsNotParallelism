package handlers

import (
	"database/sql"
	"net/http"
	"log"

	"../models"
)

// User is the handler struct for user entity
type Podcast struct {
	db *sql.DB
}

func NewPodcast(db *sql.DB) *Podcast {
	return &Podcast{db: db}
}

func (u *Podcast) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Invalid Operation")
		http.Error(w, "Invalid operation", http.StatusMethodNotAllowed)
		return
	}
	podcastModel := models.NewPodcast()
	podcastModel.PodcasterName = r.FormValue("podcaster-name")
	podcastModel.PodcasterUrl = r.FormValue("podcaster-url")
	podcastModel.PodcastName = r.FormValue("podcast-name")
  podcastModel.PodcastUrl = r.FormValue("podcast-url")
	if _, err := podcastModel.Insert(u.db); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list", http.StatusMovedPermanently)
}
