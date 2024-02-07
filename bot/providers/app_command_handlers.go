package providers

import (
	"github.com/h3mmy/bloopyboi/bot/handlers"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"go.uber.org/zap"
)

func GetDiscordAppCommands() []models.DiscordAppCommand {
	handls := make([]models.DiscordAppCommand, 0, 3)
	handls = append(handls, handlers.NewInspiroCommand(GetInspiroService()))
	handls = append(handls, handlers.NewBlissfestCommand(GetBlissfestService()))
	bookSvc, err := GetBookService()
	if err != nil {
		logger.Error("failed to create book svc", zap.Error(err))
	} else {
		handls = append(handls, handlers.NewBookCommand(bookSvc))
	}
  return handls
}
