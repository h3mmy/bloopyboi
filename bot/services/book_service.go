package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	books "google.golang.org/api/books/v1"
	"google.golang.org/api/option"
)

type BookService struct {
	bloopyMeta models.BloopyMeta
	logger     *zap.Logger
	svc        *books.Service
}

func NewBookService(ctx context.Context, options ...option.ClientOption) *BookService {
	lgr := log.NewZapLogger().With(
		zapcore.Field{Type: zapcore.StringType, Key: ServiceLoggerFieldKey, String: "book_service"},
	)
	bookSvc, _ := books.NewService(ctx, options...)

	return &BookService{
		svc:        bookSvc,
		logger:     lgr,
		bloopyMeta: models.NewBloopyMeta(),
	}
}

func (b *BookService) SearchBook(ctx context.Context, req *models.BookSearchRequest) (*books.Volumes, error) {
	// Google's full text string can have special keywords
	// See https://developers.google.com/books/docs/v1/using#PerformingSearch
	q := b.buildSearchString(req)
	b.logger.Info(fmt.Sprintf("book req, %v", req))
	b.logger.Info(fmt.Sprintf("searching for book %s",q))
	volume, err := b.svc.Volumes.List(q).Context(ctx).MaxResults(4).Do()
	if err != nil {
		b.logger.Error("failed to get book", zap.Error(err))
		return nil, err
	}
	b.logger.Info(fmt.Sprintf("found %d books", len(volume.Items)))
	for _, book := range volume.Items {
		b.logger.Debug(fmt.Sprintf("book: %s, by %s", book.VolumeInfo.Title, strings.Join(book.VolumeInfo.Authors, "")))
	}
	return volume, nil
}

func (b *BookService) buildSearchString(req *models.BookSearchRequest) string {
	var q string = ""
	if req.Title != "" {
		q += "intitle:" + req.Title
	}
	if req.Author != "" {
		q += "inauthor:" + req.Author
	}
	if req.Publisher != "" {
		q += "inpublisher:" + req.Publisher
	}
	if  req.ISBN != "" {
		q += "isbn:" + req.ISBN
	}
	if  req.TextSnippet !="" {
		q += "intext:" + req.TextSnippet
	}
	return q
}

func (b *BookService) GetVolume(volumeId string) (*books.Volume, error) {
	volume, err := b.svc.Volumes.Get(volumeId).Context(context.TODO()).Do()
	if err!= nil {
		b.logger.Error("failed to get book", zap.Error(err))
	}
	return volume, err
}
