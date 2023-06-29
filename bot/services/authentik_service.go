package services

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/h3mmy/bloopyboi/bot/internal/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	api "goauthentik.io/api/v3"
	"golang.org/x/sync/errgroup"
)

type AuthentikService struct {
	client *api.APIClient
	logger *zap.Logger
}

func NewAuthentikService(clientgen *AuthentikClientGenerator) *AuthentikService {
	return &AuthentikService{
		client: clientgen.generateClient(),
		logger: log.NewZapLogger().With(zapcore.Field{Key: ServiceLoggerFieldKey, Type: zapcore.StringType, String: "authentik"}),
	}
}

func (s *AuthentikService) GetClient() *api.APIClient {
	return s.client
}

func (s *AuthentikService) Verify(ctx context.Context) bool {
	errGroup, ctx := errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		_, res, err := s.client.CoreApi.CoreUsersList(ctx).Execute()
		if err != nil {
			s.logger.Sugar().Error(err)
			return err
		}
		if res.StatusCode >= 200 && res.StatusCode < 300 {
			return nil
		} else {
			return errors.New(fmt.Sprintf("non-2xx response code: %v", res))
		}
	})

	return errGroup.Wait() == nil
}
