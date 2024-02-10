package bot

import (
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type Gateway struct {
	logger   *otelzap.Logger
	meta     models.BloopyMeta
	echoServ *echo.Echo
}

func NewGateway() *Gateway {
	echoServ := echo.New()
	lgr := otelzap.New(zap.L())
	return &Gateway{
		meta:     models.NewBloopyMeta(),
		echoServ: echoServ,
		logger:   lgr,
	}
}

func (g *Gateway) Start() error {
	return nil
}
