package providers

import (
	"context"

	"github.com/h3mmy/bloopyboi/bot/services"
)

func GetBookService() *services.BookService {
	return services.NewBookService(context.Background())
}
