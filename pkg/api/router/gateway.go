package api

import (
	"context"
	"net/http"

	"github.com/h3mmy/bloopyboi/bot"
	"github.com/h3mmy/bloopyboi/bot/discord"
	"github.com/h3mmy/bloopyboi/pkg/api/pb"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

var defaultGatewayConfig = &models.GatewayConfig{
	HttpPort: 8080,
	GrpcPort: 8081,
}

type AppInfoServer struct {
	pb.UnimplementedInfoServiceServer
}

type Gateway struct {
	logger   *otelzap.Logger
	meta     models.BloopyMeta
	echoServ *echo.Echo
	config   *models.GatewayConfig
	bot      *bot.BloopyBoi
}

func NewGateway(cfg *models.GatewayConfig) *Gateway {
	echoServ := echo.New()
	lgr := otelzap.New(zap.L())
	return &Gateway{
		meta:     models.NewBloopyMeta(),
		echoServ: echoServ,
		config:   cfg,
		logger:   lgr,
	}
}

func (g *Gateway) WithBotInstance(bot *bot.BloopyBoi) *Gateway {
	g.bot = bot
	return g
}

func NewDefaultGateway() *Gateway {
	return NewGateway(defaultGatewayConfig)
}

func (g *Gateway) Start() error {
	g.echoServ.GET("/info", GetAppInfo)
	dg := g.echoServ.Group("/discord")
	dg = RegisterDiscordSvcRoutes(dg, g.bot.DiscordManager)
	g.logger.Debug("registered group", zap.Bool("isnil", dg == nil))
	return g.echoServ.Start(":8080")
}

func (g *Gateway) Shutdown(ctx context.Context) error {
	return g.echoServ.Shutdown(ctx)
}

func GetAppInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
}

func RegisterDiscordSvcRoutes(echoGroup *echo.Group, discMgr *discord.DiscordManager) *echo.Group {
	dg := echoGroup
	dg.GET("/manager/meta", GetDiscordManagerMeta(discMgr))
	return dg
}

func GetDiscordManagerMeta(g *discord.DiscordManager) func(c echo.Context) error {
	return func(c echo.Context) error {
		if g == nil {
			return c.JSON(http.StatusServiceUnavailable, "Bot Instance Not Attached")
		}
		return c.JSON(http.StatusOK, g.GetDiscordService().GetMeta())
	}
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
