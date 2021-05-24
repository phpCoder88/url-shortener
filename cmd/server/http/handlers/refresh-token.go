package handlers

import "net/http"

func RefreshTokenEndpoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
}
