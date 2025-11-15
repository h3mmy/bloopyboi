package bot

import (
	"context"
	"fmt"
	"net/http"

	"github.com/h3mmy/bloopyboi/bot/discord"
	"github.com/h3mmy/bloopyboi/bot/handlers"
	"github.com/h3mmy/bloopyboi/bot/providers"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/api/pb"
	"github.com/h3mmy/bloopyboi/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type AppInfoServer struct {
	pb.UnimplementedInfoServiceServer
}

type Gateway struct {
	logger   *otelzap.Logger
	meta     models.BloopyMeta
	echoServ *echo.Echo
	config   *config.GatewayConfig
	bot      *BloopyBoi
}

func NewGateway(cfg *config.GatewayConfig) *Gateway {
	echoServ := echo.New()
	lgr := otelzap.New(zap.L())
	return &Gateway{
		meta:     models.NewBloopyMeta(),
		echoServ: echoServ,
		config:   cfg,
		logger:   lgr,
	}
}

func (g *Gateway) WithBotInstance(bot *BloopyBoi) *Gateway {
	g.bot = bot
	return g
}

func NewDefaultGateway() *Gateway {
	return NewGateway(providers.GetGatewayConfig())
}

func (g *Gateway) Start() error {
	g.logger.Debug("starting gateway with config", zap.Any("config", g.config))
	g.echoServ.GET("/info", GetAppInfo)
	dg := g.echoServ.Group("/discord")
	dg = RegisterDiscordSvcRoutes(dg, func () *discord.DiscordManager {
		if g.bot == nil {
			return nil
		}
		return g.bot.GetDiscordManager()
	})
	g.logger.Debug("registered group", zap.Bool("isnil", dg == nil))
	return g.echoServ.Start(fmt.Sprintf(":%d", g.config.HttpServerConfig.Port))
}

func (g *Gateway) Shutdown(ctx context.Context) error {
	return g.echoServ.Shutdown(ctx)
}

func GetAppInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
}

func RegisterDiscordSvcRoutes(echoGroup *echo.Group, discMgrFunc func() *discord.DiscordManager) *echo.Group {
	dg := echoGroup
	dg.GET("/manager/meta", GetDiscordManagerMeta(discMgrFunc))

	if providers.GetDiscordOauthConfig().ClientSecret != "" {
		// Linked Roles
		lr := dg.Group("/linked-roles")
		lr.GET("", func(c echo.Context) error {
			return handlers.HandleLinkedRolesRedirect(c, providers.GetDiscordOauthConfig())
		})
		lr.GET("/callback", func(c echo.Context) error {
			return handlers.HandleLinkedRolesCallback(c, providers.GetDiscordOauthConfig(), discMgrFunc().GetDiscordService())
		})
	}
	return dg
}

func GetDiscordManagerMeta(g func () *discord.DiscordManager) func(c echo.Context) error {
	return func(c echo.Context) error {
		if g == nil {
			return c.JSON(http.StatusServiceUnavailable, "Bot Instance Not Attached")
		}
		return c.JSON(http.StatusOK, g().GetDiscordService().GetMeta())
	}
}

// TODO: finish this

// func GetRoleConnectionInfo(g *discord.DiscordManager) func(c echo.Context) error{
// return func(c echo.Context) error {
// 		if g == nil {
// 			return c.JSON(http.StatusServiceUnavailable, "Bot Instance Not Attached")
// 		}
// 		return c.JSON(http.StatusOK, g.GetDiscordService().GetDiscordUserRoleConnection())
// 	}
// }

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
