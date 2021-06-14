package handlers

import (
	"github.com/phpCoder88/url-shortener/internal/config"
	"github.com/phpCoder88/url-shortener/internal/ioc"
	"go.uber.org/zap"
)

type Handler struct {
	logger    *zap.SugaredLogger
	conf      *config.Config
	container *ioc.Container
}

func NewHandler(logger *zap.SugaredLogger, conf *config.Config, container *ioc.Container) *Handler {
	return &Handler{
		logger:    logger,
		conf:      conf,
		container: container,
	}
}
