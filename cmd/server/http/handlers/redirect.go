package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RedirectFullURL(res http.ResponseWriter, req *http.Request) {
	log.Printf("Requested token %s\n", mux.Vars(req)["token"])

	// TODO: Проверка на существование URL для переданного токена

	http.Redirect(res, req, "https://golang.org/", http.StatusSeeOther)
}
