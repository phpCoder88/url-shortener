package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				log.Println(fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
