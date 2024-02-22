package services

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/bot/internal/database"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	pmodels "github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/ent"
	"github.com/h3mmy/bloopyboi/ent/book"
	"github.com/h3mmy/bloopyboi/ent/bookauthor"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
	"github.com/h3mmy/bloopyboi/ent/mediarequest"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	books "google.golang.org/api/books/v1"
	"google.golang.org/api/option"
)

type BookService struct {
	bloopyMeta models.BloopyMeta
	logger     *zap.Logger
	svc        *books.Service
	db         *ent.Client
	dbEnabled  bool
}

func NewBookService(ctx context.Context, options ...option.ClientOption) (*BookService, error) {
	meta := models.NewBloopyMeta()
	lgr := log.NewZapLogger().
		Named("book_service").
		With(
			zapcore.Field{Type: zapcore.StringType, Key: "bloopy_id", String: meta.Id.String()},
			zapcore.Field{Type: zapcore.TimeFullType, Key: "created_at", Interface: meta.CreatedAt},
			zapcore.Field{Type: zapcore.StringType, Key: ServiceLoggerFieldKey, String: "book_service"},
		)
	dbEnabled := true
	bookSvc, err := books.NewService(ctx, options...)
	if err != nil {
		lgr.Error("failed to create book service", zap.Error(err))
		return nil, err
	}
	dbClient, err := database.Open()
	if err != nil {
		lgr.Error("failed to open database", zap.Error(err))
		dbEnabled = false
	}
	return &BookService{
		svc:        bookSvc,
		logger:     lgr,
		bloopyMeta: meta,
		db:         dbClient,
		dbEnabled:  dbEnabled,
	}, nil
}

func (b *BookService) IsReady() bool {
	if b.svc == nil {
		return false
	}
	return b.bloopyMeta.Id != uuid.Nil
}

func (b *BookService) SearchBook(ctx context.Context, req *models.BookSearchRequest) (*books.Volumes, error) {
	b.logger.Debug(fmt.Sprintf("context: %v", ctx))
	// Google's full text string can have special keywords
	// See https://developers.google.com/books/docs/v1/using#PerformingSearch
	q := b.buildSearchString(req)
	b.logger.Info(fmt.Sprintf("book req, %v", req))
	b.logger.Info(fmt.Sprintf("searching for book %s", q))
	volume, err := b.svc.Volumes.
		List(q).
		Context(context.Background()).
		MaxResults(4).
		Do()
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
	if req.ISBN != "" {
		q += "isbn:" + req.ISBN
	}
	if req.TextSnippet != "" {
		q += "intext:" + req.TextSnippet
	}
	return q
}

func (b *BookService) GetVolume(volumeId string) (*books.Volume, error) {
	volume, err := b.svc.Volumes.Get(volumeId).Context(context.TODO()).Do()
	if err != nil {
		b.logger.Error("failed to get book", zap.Error(err))
	}
	return volume, err
}

func (b *BookService) SubmitBookRequest(ctx context.Context, discUser *discordgo.User, volumeId string) error {
	volume, err := b.GetVolume(volumeId)
	if err != nil {
		return err
	}
	if b.dbEnabled {

		err := database.WithTx(ctx, b.db, func(tx *ent.Tx) error {
			return tx.DiscordUser.
				Create().
				SetID(uuid.New()).
				SetDiscordid(discUser.ID).
				SetUsername(discUser.Username).
				SetEmail(discUser.Email).
				SetDiscriminator(discUser.Discriminator).
				OnConflict(
					sql.ConflictColumns(discorduser.FieldDiscordid),
				).
				UpdateNewValues().
				Exec(ctx)
		})
		if err != nil {
			return fmt.Errorf("failed to save discord user with id %s: %w", discUser.ID, err)
		} else {
			b.logger.Debug(fmt.Sprintf("saved discord user id: %s", discUser.ID))
		}
		discordUserId, err := b.db.DiscordUser.
			Query().
			Where(discorduser.DiscordidEQ(discUser.ID)).
			FirstID(ctx)

		if err != nil {
			return fmt.Errorf("failed to find discord user with id %s: %w", discUser.ID, err)
		} else {
			b.logger.Debug(fmt.Sprintf("found discord user id: %s", discordUserId))
		}

		bookid, err := b.SaveBook(ctx, volume)

		if err != nil {
			return fmt.Errorf("failed to save book: %w", err)
		}

		err = database.WithTx(ctx, b.db, func(tx *ent.Tx) error {
			return tx.MediaRequest.
				Create().
				SetID(uuid.New()).
				AddBookIDs(bookid).
				SetDiscordUserID(discordUserId).
				SetStatus(string(pmodels.MediaRequestStatusRequested)).
				Exec(ctx)
		})
		if err != nil {
			return fmt.Errorf("failed to save media request: %w", err)
		} else {
			b.logger.Debug(fmt.Sprintf("saved media request id: %s", volumeId))
		}

		for _, author := range volume.VolumeInfo.Authors {
			err = database.WithTx(ctx, b.db, func(tx *ent.Tx) error {
				return b.db.BookAuthor.
					Create().
					SetID(uuid.New()).
					SetFullName(author).
					AddBookIDs(bookid).
					OnConflict(sql.ConflictColumns(bookauthor.FieldFullName)).
					UpdateNewValues().
					Exec(ctx)
			})
			if err != nil {
				return fmt.Errorf("failed to save book author: %w", err)
			}
		}
		return nil
	}
	return nil
}

func (b *BookService) SaveBook(ctx context.Context, volume *books.Volume) (uuid.UUID, error) {
	// TODO: Parse and include ISBNs
	err := database.WithTx(ctx, b.db, func(tx *ent.Tx) error {
		return tx.Book.
			Create().
			SetID(uuid.New()).
			SetGoogleVolumeID(volume.Id).
			SetDescription(volume.VolumeInfo.Description).
			SetTitle(volume.VolumeInfo.Title).
			SetPublisher(volume.VolumeInfo.Publisher).
			SetImageURL(volume.VolumeInfo.ImageLinks.Thumbnail).
			OnConflict(sql.ConflictColumns(book.FieldGoogleVolumeID)).
			UpdateNewValues().
			Exec(ctx)
	})

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed persisting book to db: %w", err)
	} else {
		b.logger.Debug(fmt.Sprintf("saved book id: %s", volume.Id))
	}

	bookid, err := b.db.Book.Query().
		Where(book.GoogleVolumeIDEQ(volume.Id)).
		FirstID(ctx)

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to find book with volume id %s: %w", volume.Id, err)
	}
	b.logger.Debug(fmt.Sprintf("found book id: %s", bookid))
	return bookid, nil
}

// GetAllBookRequestsForUser returns all book requests for a user.
func (b *BookService) GetAllBookRequestsForUser(ctx context.Context, userId string) ([]*ent.MediaRequest, error) {
	if b.dbEnabled {
		requests, err := b.db.MediaRequest.
			Query().
			WithBook().
			Where(
				mediarequest.HasDiscordUserWith(
					discorduser.DiscordidEQ(userId),
				),
			).
			All(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get book requests for user %s: %w", userId, err)
		}
		return requests, nil
	}
	b.logger.Warn("database not enabled")
	return nil, nil
}
