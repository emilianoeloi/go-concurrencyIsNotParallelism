package models

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	// UserTableName is the name of the table for the user model
	UserTableName = "user"
	// UserIDCol is the column name of the model's ID
	UserIDCol = "id_user"
	// UserNameCol is the column name of the model's name
	UserUrl1Col = "url1"
	// UserEmailCol is the column name of the model's email
	UserUrl2Col = "url2"
	// UserEmailCol is the column name of the model's email
	UserUrl3Col = "url3"
)

// User is the type to represent an user entity
type User struct {
	ID    int64  `json:"id"`
	Url1  string `json:"url1"`
	Url2  string `json:"url2"`
	Url3  string `json:"url3"`
}

// NewUser return an instance of user
func NewUser() *User {
	return &User{}
}

// Insert is the function to insert an user
func (u *User) Insert(db *sql.DB) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES(?, ?, ?)", UserTableName, UserUrl1Col, UserUrl2Col, UserUrl3Col),
		u.Url1,
		u.Url2,
		u.Url3,
	)
}

// SelectAll selects all users.
func (u *User) SelectAll(db *sql.DB) ([]User, error) {
	rows, err := db.Query(
		fmt.Sprintf(
			"SELECT %s, %s, %s, %s FROM %s ORDER BY %s ASC",
			UserIDCol,
			UserUrl1Col,
			UserUrl2Col,
			UserUrl3Col,
			UserTableName,
			UserUrl1Col,
		),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var retUrl1, retUrl2, retUrl3 string
		var retID int64
		if err := rows.Scan(&retID, &retUrl1, &retUrl2, &retUrl3); err != nil {
			return nil, err
		}
		log.Println(retUrl1)
		users = append(users, User{ID: retID, Url1: retUrl1, Url2: retUrl2, Url3: retUrl3})
	}
	return users, nil
}
