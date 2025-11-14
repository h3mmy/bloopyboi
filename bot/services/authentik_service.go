package services

import (
	"context"
	"fmt"

	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	api "goauthentik.io/api/v3"
	"golang.org/x/sync/errgroup"
)

// AuthentikService is a service that interacts with the Authentik API.
type AuthentikService struct {
	client *api.APIClient
	logger *zap.Logger
}

// NewAuthentikService creates a new AuthentikService.
func NewAuthentikService(clientgen *AuthentikClientGenerator) *AuthentikService {
	return &AuthentikService{
		client: clientgen.generateClient(),
		logger: log.NewZapLogger().With(zapcore.Field{Key: ServiceLoggerFieldKey, Type: zapcore.StringType, String: "authentik"}),
	}
}

// GetClient returns the Authentik API client.
func (s *AuthentikService) GetClient() *api.APIClient {
	return s.client
}

// Verify verifies that the Authentik API is reachable and that the credentials are valid.
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
			return fmt.Errorf("non-2xx response code: %v", res)
		}
	})

	return errGroup.Wait() == nil
}
