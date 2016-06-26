package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"log"

	"strconv"

	"github.com/viniciusfeitosa/MyFirstAppWithGo/models"
)

type Language struct {
	db *sql.DB
}

func NewLanguage(db *sql.DB) *Language {
	return &Language{db: db}
}

func (u *Language) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("Invalid Operation")
		http.Error(w, "Invalid operation", http.StatusMethodNotAllowed)
		return
	}

	languageModel := models.NewLanguage()
	language, err := languageModel.SelectAll(u.db)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Users:", users)

	if err := serveTemplate(w, "list", language); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u *Language) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Invalid Operation")
		http.Error(w, "Invalid operation", http.StatusMethodNotAllowed)
		return
	}
	languageModel := models.NewLanguage()
	languageModel.type = r.FormValue("type")
	languageModel.name = r.FormValue("name")
  HEEEERE!!!!
	if _, err := userModel.Insert(u.db); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list", http.StatusMovedPermanently)
}

// Update change user on DB
func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		log.Println("Invalid Operation")
		http.Error(w, "Invalid operation", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var userModel models.User
	if err := json.NewDecoder(r.Body).Decode(&userModel); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := userModel.Update(u.db, id); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"Success\"}"))
}

// FindAll returns all users from DB
func (u *User) FindAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("Invalid Operation")
		http.Error(w, "Invalid operation", http.StatusMethodNotAllowed)
		return
	}
	userModel := models.NewUser()
	users, err := userModel.SelectAll(u.db)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(users)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

// FindOne returns all users from DB
func (u *User) FindOne(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("Invalid Operation")
		http.Error(w, "Invalid operation", http.StatusMethodNotAllowed)
		return
	}
	userModel := models.NewUser()
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := userModel.SelectOne(u.db, id); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(userModel)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

// Remove delete a new user from DB
func (u *User) Remove(w http.ResponseWriter, r *http.Request) {
	userModel := models.NewUser()
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = userModel.Delete(u.db, id); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("User id %d was removed\n", id)
	http.Redirect(w, r, "/list", http.StatusMovedPermanently)
}
