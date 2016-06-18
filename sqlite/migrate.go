package sqlite

import (
	"fmt"

	"../models"
)

var (
	// createUserTableQuery is the query to create user table
	createUserTableQuery = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s integer primary key autoincrement, %s varchar(255), %s varchar(255), %s varchar(255))",
		models.UserTableName,
		models.UserIDCol,
		models.UserUrl1Col,
		models.UserUrl2Col,
		models.UserUrl3Col,
	)

	createPodcastTableQuery = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s integer primary key autoincrement, %s varchar(255), %s varchar(255), %s varchar(255), %s varchar(255))",
		models.PodcastTableName,
		models.PodcastIDCol,
		models.PodcasterNameCol,
		models.PodcasterUrlCol,
		models.PodcastNameCol,
		models.PodcastUrlCol,
	)

	// QuerysToMigrate is the slice of queries to run on migration
	QuerysToMigrate = []string{createUserTableQuery, createPodcastTableQuery}
)
