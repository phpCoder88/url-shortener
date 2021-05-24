package handlers

import "net/http"

func LoginEndpoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
}
