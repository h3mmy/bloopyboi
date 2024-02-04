package providers

import (
	"github.com/h3mmy/bloopyboi/bot/handlers"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
)

func GetDiscordAppCommands() []models.DiscordAppCommand {
  return []models.DiscordAppCommand{
		handlers.NewInspiroCommand(GetInspiroService()),
		handlers.NewBlissfestCommand(GetBlissfestService()),
		handlers.NewBookCommand(GetBookService()),
	}
}
