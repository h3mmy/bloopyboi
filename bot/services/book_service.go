package services

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/ent"
	"github.com/h3mmy/bloopyboi/ent/book"
	"github.com/h3mmy/bloopyboi/ent/bookauthor"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
	"github.com/h3mmy/bloopyboi/ent/mediarequest"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/database"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	books "google.golang.org/api/books/v1"
	"google.golang.org/api/option"
)

// BookService is a service that interacts with the Google Books API and the local database.
type BookService struct {
	bloopyMeta models.BloopyMeta
	logger     *zap.Logger
	svc        *books.Service
	db         *ent.Client
	dbEnabled  bool
}

// NewBookService creates a new BookService.
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
		lgr.Warn("failed to open database", zap.Error(err))
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

// IsReady returns true if the service is ready.
func (b *BookService) IsReady() bool {
	if b.svc == nil {
		return false
	}
	return b.bloopyMeta.Id != uuid.Nil
}

// IsDatabaseEnabled returns true if the database is enabled.
func (b *BookService) IsDatabaseEnabled() bool {
	return b.dbEnabled
}

// Shutdown shuts down the service.
func (b *BookService) Shutdown() {
	b.logger.Info("shutting down")
	if b.db != nil {
		b.logger.Info("closing database")
		err := b.db.Close()
		if err != nil {
			b.logger.Error("failed to close database", zap.Error(err))
		}
	}
}

// RefreshDatabaseConnection refreshes the database connection.
func (b *BookService) RefreshDatabaseConnection() {
	b.logger.Debug("refreshing DB Client")
	dbClient, err := database.Open()
	if err != nil {
		b.logger.Warn("failed to open database. Marking disabled", zap.Error(err))
	} else {
		b.logger.Info("database connected successfully")
		b.db = dbClient
		b.dbEnabled = true
	}
}

// SearchBook searches for a book.
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

// buildSearchString builds a search string for the Google Books API.
func (b *BookService) buildSearchString(req *models.BookSearchRequest) string {
	var q = ""
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

// GetVolume gets a volume from the Google Books API.
func (b *BookService) GetVolume(volumeId string) (*books.Volume, error) {
	volume, err := b.svc.Volumes.Get(volumeId).Context(context.TODO()).Do()
	if err != nil {
		b.logger.Error("failed to get book", zap.Error(err))
	}
	return volume, err
}

// SubmitBookRequest submits a book request.
func (b *BookService) SubmitBookRequest(ctx context.Context, discUser *discordgo.User, volumeId string) (*ent.MediaRequest, error) {
	volume, err := b.GetVolume(volumeId)
	if err != nil {
		return nil, err
	}
	if !b.dbEnabled {
		return nil, nil
	}

	err = database.WithTx(ctx, b.db, func(tx *ent.Tx) error {
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
		return nil, fmt.Errorf("failed to save discord user with id %s: %w", discUser.ID, err)
	} else {
		b.logger.Debug(fmt.Sprintf("saved discord user id: %s", discUser.ID))
	}
	discordUserId, err := b.db.DiscordUser.
		Query().
		Where(discorduser.DiscordidEQ(discUser.ID)).
		FirstID(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to find discord user with id %s: %w", discUser.ID, err)
	} else {
		b.logger.Debug(fmt.Sprintf("found discord user id: %s", discordUserId))
	}

	bookid, err := b.SaveBook(ctx, volume)

	if err != nil {
		return nil, fmt.Errorf("failed to save book: %w", err)
	}

	// check book to see if there is an existing request

	mediareq, err := b.db.MediaRequest.Query().
		Where(
			mediarequest.HasBookWith(book.IDEQ(bookid)),
		).
		WithDiscordUsers().
		WithBook().
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			// Mostly expected scenario
			b.logger.Debug(fmt.Sprintf("no existing request for book %s and user %s", bookid, discordUserId))
			err = database.WithTx(ctx, b.db, func(tx *ent.Tx) error {
				return tx.MediaRequest.
					Create().
					SetID(uuid.New()).
					SetBookID(bookid).
					AddDiscordUserIDs(discordUserId).
					SetStatus(models.MediaRequestStatusRequested).
					Exec(ctx)
			})
			if err != nil {
				return nil, fmt.Errorf("failed to save media request: %w", err)
			} else {
				b.logger.Debug(fmt.Sprintf("saved media request id: %s", volumeId))
			}
		} else {
			return nil, fmt.Errorf("error checking for existing request for book %s and user %s: %w", bookid, discordUserId, err)
		}
	} else {
		// media request already exists
		b.logger.Debug(fmt.Sprintf("existing request for book %s with id %v", bookid, mediareq.ID))
		// Add user to request if not already in it
		dUser, err := mediareq.QueryDiscordUsers().
			Where(discorduser.DiscordidEQ(discUser.ID)).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				b.logger.Debug(fmt.Sprintf("user %s not in existing request %s. Adding...", discordUserId, mediareq.ID))
				err = database.WithTx(ctx, b.db, func(tx *ent.Tx) error {
					return tx.MediaRequest.
						UpdateOneID(mediareq.ID).
						AddDiscordUserIDs(discordUserId).
						Exec(ctx)
				})
				if err != nil {
					return nil, fmt.Errorf("failed to add user %s to existing request %s: %w", discordUserId, mediareq.ID, err)
				}
			} else {
				return nil, fmt.Errorf("error checking for user %s in existing request %s: %w", discordUserId, mediareq.ID, err)
			}
		} else {
			b.logger.Debug(fmt.Sprintf("user %s already in existing request %s", dUser.Username, mediareq.ID))
			// message user about it
		}
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
			return mediareq, fmt.Errorf("failed to save book author: %w", err)
		}
	}
	return mediareq, nil
}

// SaveBook saves a book to the database.
func (b *BookService) SaveBook(ctx context.Context, volume *books.Volume) (uuid.UUID, error) {
	// Check for existing volume first
	bookid, err := b.db.Book.Query().
		Where(book.GoogleVolumeIDEQ(volume.Id)).
		FirstID(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			b.logger.Debug(fmt.Sprintf("volume %s not found in DB. Adding...", volume.Id))
			bookid = uuid.New()
			// TODO: Parse and include ISBNs
			err = database.WithTx(ctx, b.db, func(tx *ent.Tx) error {
				return tx.Book.
					Create().
					SetID(bookid).
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
				b.logger.Debug(fmt.Sprintf("saved book for volume id: %s", volume.Id))
			}
		} else {
			return uuid.Nil, fmt.Errorf("something went wrong when finding a book with volume id %s: %w", volume.Id, err)
		}

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
				mediarequest.HasDiscordUsersWith(
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

// BuildBookRequestStatusAsEmbed builds an embed for a book request.
func (b *BookService) BuildBookRequestStatusAsEmbed(ctx context.Context, req *ent.MediaRequest) *discordgo.MessageEmbed {
	book := req.Edges.Book
	authors, err := book.QueryBookAuthor().Select(bookauthor.FieldFullName).Strings(ctx)
	if err != nil {
		b.logger.Warn("could not retrieve book authors", zap.Error(err))
		authors = []string{}
	}
	return &discordgo.MessageEmbed{
		Image: &discordgo.MessageEmbedImage{
			URL: book.ImageURL,
		},
		Title: fmt.Sprintf("%s by %s", book.Title, strings.Join(authors, "")),

		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Publisher",
				Value: book.Publisher,
			},
			{
				Name:  "Volume ID",
				Value: book.GoogleVolumeID,
			},
			{
				Name:  "Requested",
				Value: req.CreateTime.Format("2006-01-02"),
			},
			{
				Name:  "Status",
				Value: string(req.Status),
			},
		},
	}
}
