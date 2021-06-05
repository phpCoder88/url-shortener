package routes

import (
	"net/http"

	"github.com/phpCoder88/url-shortener/cmd/server/http/handlers"
	"github.com/phpCoder88/url-shortener/cmd/server/http/middlewares"

	gHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func Routes() http.Handler {
	standardMiddleware := alice.New(middlewares.RecoverPanic)

	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/shorten", handlers.ShortenEndpoint).Methods("POST")
	api.HandleFunc("/report", handlers.ReportEndpoint).Methods("GET")
	router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./cmd/swaggerui"))))
	router.HandleFunc("/", handlers.RedirectFullURL).Methods("GET").Queries("t", "{token}")

	methods := gHandlers.AllowedMethods([]string{
		"GET",
		"POST",
	})
	headers := gHandlers.AllowedHeaders([]string{
		"Content-Type",
		"Authorization",
		"X-Requested-With",
	})
	origins := gHandlers.AllowedOrigins([]string{"*"})

	return gHandlers.CORS(headers, methods, origins)(standardMiddleware.Then(router))
}
