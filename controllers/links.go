package controllers

import (
	"main/db"
	"main/utils"
	"net/http"
	"time"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("t")
	var link db.Link
	db.DB.First(&link, "token = ?", t)
	if link.ID == 0 {
		w.Write([]byte("404 Not Found"))
		return
	}

	http.Redirect(w, r, link.Original, http.StatusTemporaryRedirect)
}

func Add(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		utils.UUID()
	}
	original := r.URL.Query().Get("link")
	token := r.URL.Query().Get("token")
	if token == "" {
		token = utils.UUID()
	} else {
		var _link db.Link
		db.DB.First(&_link, "token = ?", token)
		if _link.ID != 0 {
			w.Write([]byte("Token exists"))
			return
		}
	}

	link := db.Link{
		Name:      name,
		Token:     token,
		Original:  original,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.DB.Create(&link)

	uri := "https://s.robet.my.id?t=" + token

	w.Write([]byte(uri))
}
