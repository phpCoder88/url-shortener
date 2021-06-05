// Package main URL shortener API.
//
// Open API for URL shortener service
//
// Terms Of Service:
//
//     Schemes: http
//     Host: localhost:8000
//     BasePath: /api
//     Version: 1.0.0
//     License: MIT https://opensource.org/licenses/MIT
//     Contact: Pavel Bobylev<p_bobylev@bk.ru> https://github.com/phpCoder88
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/phpCoder88/url-shortener/cmd/server/http/routes"
	"github.com/phpCoder88/url-shortener/internal/config"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	api := &http.Server{
		Addr:         net.JoinHostPort("", fmt.Sprint(conf.Port)),
		Handler:      routes.Routes(),
		IdleTimeout:  conf.IdleTimeout,
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
	}

	go func() {
		fmt.Printf("Server is listening %d port...\n", conf.Port)
		err = api.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	// Graceful shutdown
	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, syscall.SIGINT, syscall.SIGTERM)
	<-osSignalChan

	ctx, cancel := context.WithTimeout(context.Background(), conf.ShutdownTimeout)
	defer cancel()

	log.Println("Shutting down...")
	err = api.Shutdown(ctx)
	if err != nil {
		log.Println(err)
		return
	}
}
