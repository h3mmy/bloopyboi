package providers

import (
	"context"
	"fmt"

	"gitlab.com/h3mmy/bloopyboi/bot/internal/ent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

const (
	ServiceLoggerFieldKey = "service_name"
)


type BloopyEnt struct {
	client *ent.Client
	quit   *chan struct{}
	logger *zap.Logger
	active bool
}

func NewBloopyEnt() *BloopyEnt {
	lgr := NewZapLogger().With(zapcore.Field{
		Key:    ServiceLoggerFieldKey,
		Type:   zapcore.StringType,
		String: "bloopyEnt",
	})

	qch := make(chan struct{})
	return &BloopyEnt{
		logger: lgr,
		quit:   &qch,
	}
}

func (be *BloopyEnt) Start(ctx context.Context) error {

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		be.logger.Fatal("failed to create entgo client while initializing BloopyEnt")
		panic(err)
	}

	be.client = client

	// Run the auto migration tool.
	if err := be.client.Schema.Create(ctx); err != nil {
		be.logger.Sugar().Fatalf("failed creating schema resources: %v", err)
	}
	defer client.Close()
	defer be.client.Close()

	be.active = true

	<-*be.quit
	be.logger.Info("Received close signal. Exiting")
	be.active = false

	return nil
}

func (be *BloopyEnt) Quit(ctx context.Context) error {
	if be.active {
		be.logger.Info("client is active, sending close signal")
		close(*be.quit)
	} else {
		be.logger.Info("client not active. Nothing to do")
	}
	return nil
}


func(be *BloopyEnt) AddDiscordMessage(ctx context.Context, dmsg *discordgo.Message) error {
	mzField := zapcore.Field{Key: "func", Type: zapcore.StringType, String: "AddDiscordMessage"}
	if !be.active {
		be.logger.Error("client not initialized", mzField)
		return fmt.Errorf("bloopyent not initialized.")
	}

	dbentry, err := be.client.DiscordMessage.Create().SetID(dmsg.ID).SetRaw(*dmsg).Save(ctx)

	if err != nil {
		be.logger.Error("error adding discord message", mzField)
		return err
	}

	be.logger.Debug(fmt.Sprintf("Added Discord Message with ID %s", dbentry.ID), mzField)

	return nil
}
