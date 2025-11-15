package bot

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/h3mmy/bloopyboi/bot/discord"
	"github.com/h3mmy/bloopyboi/bot/handlers"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/api/pb"
	"github.com/labstack/echo-contrib/session"
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
	bot      *BloopyBoi
}

func NewGateway(cfg *models.GatewayConfig) *Gateway {
	echoServ := echo.New()
	echoServ.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
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
	return NewGateway(defaultGatewayConfig)
}

func (g *Gateway) Start() error {
	g.echoServ.GET("/info", GetAppInfo)
	dg := g.echoServ.Group("/discord")
	dg = RegisterDiscordSvcRoutes(dg, g.bot)
	g.logger.Debug("registered group", zap.Bool("isnil", dg == nil))
	return g.echoServ.Start(":8080")
}

func (g *Gateway) Shutdown(ctx context.Context) error {
	return g.echoServ.Shutdown(ctx)
}

func GetAppInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
}

func RegisterDiscordSvcRoutes(echoGroup *echo.Group, bot *BloopyBoi) *echo.Group {
	dg := echoGroup
	dg.GET("/manager/meta", GetDiscordManagerMeta(bot.DiscordManager))

	// Linked Roles
	lr := dg.Group("/linked-roles")
	lr.GET("", func(c echo.Context) error {
		return handlers.HandleLinkedRolesRedirect(c, &bot.OAuthConfig)
	})
	lr.GET("/callback", func(c echo.Context) error {
		return handlers.HandleLinkedRolesCallback(c, &bot.OAuthConfig)
	})

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
