package bot

import (
	"context"

	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"github.com/h3mmy/bloopyboi/bot/servers"
	pmodels "github.com/h3mmy/bloopyboi/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

var defaultGatewayConfig = &pmodels.GatewayConfig{
	HttpPort: 8080,
	GrpcPort: 8081,
}

type AppInfoServer struct {
	servers.UnimplementedInfoServiceServer
}

type Gateway struct {
	logger   *otelzap.Logger
	meta     models.BloopyMeta
	echoServ *echo.Echo
	config   *pmodels.GatewayConfig
}

func NewGateway(cfg *pmodels.GatewayConfig) *Gateway {
	echoServ := echo.New()
	lgr := otelzap.New(zap.L())
	return &Gateway{
		meta:     models.NewBloopyMeta(),
		echoServ: echoServ,
		config:   cfg,
		logger:   lgr,
	}
}

func NewDefaultGateway() *Gateway {
	return NewGateway(defaultGatewayConfig)
}

func (g *Gateway) Start() error {
	g.echoServ.GET("/info", GetAppInfo)

	return g.echoServ.Start(":8080")
}

func (g *Gateway) Shutdown(ctx context.Context) error {
	return g.echoServ.Shutdown(ctx)
}

func GetAppInfo(c echo.Context) error {
	return c.JSON(200, "Hello World")
}

// func (g *Gateway) startGRPC() error {
// 	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", g.config.GrpcPort))
// 	if err != nil {
// 		g.logger.Error("failed to listen", zap.Error(err))
// 		return err
// 	}

// 	gserv := grpc.NewServer()
// 	servers.RegisterInfoServiceServer(gserv, &AppInfoServer{})
// 	return gserv.Serve(lis)
// }
