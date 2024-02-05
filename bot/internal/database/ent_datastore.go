package database

import (
	"context"
	"time"

	"entgo.io/ent/dialect"
	"github.com/h3mmy/bloopyboi/ent"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go.uber.org/zap"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

var logger *zap.Logger = zap.L().Named("ent_store")

func Open() (*ent.Client, error) {
	// TODO: Make this configurable
	oteldb, err := otelsql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1&_journal=WAL")
	if err != nil {
		logger.Error("failed to create entgo client while initializing BloopyEnt")
		return nil, err
	}
	drv := entsql.OpenDB(dialect.SQLite, oteldb)
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	client := ent.NewClient(ent.Driver(drv))

	err = migrateSchema(client, context.Background())

	return client, err
}

func migrateSchema(client *ent.Client, ctx context.Context) error {
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		logger.Error("failed creating schema resources", zap.Error(err))
		return err
	}
	return nil
}
