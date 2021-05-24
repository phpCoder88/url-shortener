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

	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", handlers.RegisterEndpoint).Methods("POST")
	authRouter.HandleFunc("/login", handlers.LoginEndpoint).Methods("POST")
	authRouter.HandleFunc("/refresh-token", handlers.RefreshTokenEndpoint).Methods("POST")

	router.HandleFunc("/shorten", handlers.ShortenEndpoint).Methods("POST")
	router.HandleFunc("/report", handlers.ReportEndpoint).Methods("GET")

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
