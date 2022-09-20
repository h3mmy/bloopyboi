package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/log"
	"gitlab.com/h3mmy/bloopyboi/bot/providers"
	api "goauthentik.io/api/v3"
	"golang.org/x/sync/errgroup"
)

type AuthentikService struct {
	client *api.APIClient
	logger logrus.FieldLogger
}

func NewAuthentikService(clientgen *AuthentikClientGenerator) *AuthentikService {
	return &AuthentikService{
		client: clientgen.generateClient(),
		logger: log.DefaultBloopyFieldLogger().WithField(providers.ServiceLoggerFieldKey, "authentik"),
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
			s.logger.Error(err)
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
