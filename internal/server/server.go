package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/phpCoder88/url-shortener/internal/http/routes"

	"github.com/phpCoder88/url-shortener/internal/config"
	"github.com/phpCoder88/url-shortener/internal/ioc"

	"go.uber.org/zap"
)

// Server struct
type Server struct {
	server    http.Server
	logger    *zap.SugaredLogger
	conf      *config.Config
	container *ioc.Container
}

func NewServer(logger *zap.SugaredLogger, conf *config.Config, container *ioc.Container) *Server {
	return &Server{
		server: http.Server{
			Addr:         net.JoinHostPort("", fmt.Sprint(conf.Server.Port)),
			Handler:      routes.Routes(logger, conf, container),
			IdleTimeout:  conf.Server.IdleTimeout,
			ReadTimeout:  conf.Server.ReadTimeout,
			WriteTimeout: conf.Server.WriteTimeout,
		},
		logger:    logger,
		conf:      conf,
		container: container,
	}
}

func (s *Server) Run() error {
	go func() {
		s.logger.Infof("Server is listening on PORT: %d...", s.conf.Server.Port)
		err := s.server.ListenAndServe()
		if err != nil {
			s.logger.Error(err)
		}
	}()

	// Graceful shutdown
	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, syscall.SIGINT, syscall.SIGTERM)
	<-osSignalChan

	ctx, cancel := context.WithTimeout(context.Background(), s.conf.Server.ShutdownTimeout)
	defer cancel()

	s.logger.Info("Starting to shutdown the server...")
	return s.server.Shutdown(ctx)
}
