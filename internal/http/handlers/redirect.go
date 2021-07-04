package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RedirectFullURL(res http.ResponseWriter, req *http.Request) {
	token := mux.Vars(req)["token"]
	h.logger.Infof("Requested token %s", mux.Vars(req)["token"])

	fullURL, err := h.container.ShortenerService.GetFullURL(token)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		h.logger.Error(err)
		return
	}

	http.Redirect(res, req, fullURL, http.StatusSeeOther)
}
