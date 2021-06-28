package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RedirectFullURL(res http.ResponseWriter, req *http.Request) {
	h.logger.Infof("Requested token %s", mux.Vars(req)["token"])

	// TODO: Проверка на существование URL для переданного токена

	http.Redirect(res, req, "https://golang.org/", http.StatusSeeOther)
}
