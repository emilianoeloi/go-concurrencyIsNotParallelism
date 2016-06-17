package handlers

import (
	"database/sql"
	"net/http"
	"fmt"

	"log"

	"../models"
)

// User is the handler struct for user entity
type User struct {
	db *sql.DB
}

// NewUser create a new instance of user handler
func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

// Create insert a new user on DB
func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Invalid Operation")
		http.Error(w, "Invalid operation", http.StatusMethodNotAllowed)
		return
	}
	fmt.Printf(r.FormValue("url[0]"))
	userModel := models.NewUser()
	userModel.Url1 = r.FormValue("url[0]")
	userModel.Url2 = r.FormValue("url[1]")
	userModel.Url3 = r.FormValue("url[2]")
	if _, err := userModel.Insert(u.db); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list", http.StatusMovedPermanently)
}
