package models

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	PodcastTableName = "podcasts"
	PodcastIDCol = "podcast_id"
	PodcasterNameCol = "podcaster_name"
	PodcasterUrlCol = "podcaster_url"
	PodcastNameCol = "podcast_name"
  PodcastUrlCol = "podcast_url"
)

type Podcast struct {
	ID    int64  `json:"id"`
	PodcasterName  string `json:"podcasterName"`
	PodcasterUrl  string `json:"podcasterUrl"`
	PodcastName  string `json:"podcastName"`
  PodcastUrl  string `json:"podcastUrl"`
}

func NewPodcast() *Podcast {
	return &Podcast{}
}

func (u *Podcast) Insert(db *sql.DB) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s) VALUES(?, ?, ?, ?)", PodcastTableName, PodcasterNameCol, PodcasterUrlCol, PodcastNameCol, PodcastUrlCol),
		u.PodcasterName,
		u.PodcasterUrl,
		u.PodcastName,
    u.PodcastUrl,
	)
}

func (u *Podcast) SelectAll(db *sql.DB) ([]Podcast, error) {
	rows, err := db.Query(
		fmt.Sprintf(
			"SELECT %s, %s, %s, %s FROM %s ORDER BY %s ASC",
			PodcastIDCol,
			PodcasterNameCol,
			PodcasterUrlCol,
			PodcastNameCol,
			PodcastUrlCol,
			PodcastTableName,
			PodcasterNameCol,
		),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var podcasts []Podcast
	for rows.Next() {
		var retPodcasterName, retPodcasterUrl, retPodcastName, retPodcastUrl string
		var retID int64
		if err := rows.Scan(&retID, &retPodcasterName, &retPodcasterUrl, &retPodcastName, &retPodcastUrl); err != nil {
			return nil, err
		}
		log.Println(retPodcasterName)
		podcasts = append(podcasts, Podcast{ID: retID, PodcasterName: retPodcasterName, PodcasterUrl: retPodcasterUrl, PodcastName: retPodcastName, PodcastUrl: retPodcastUrl})
	}
	return podcasts, nil
}
