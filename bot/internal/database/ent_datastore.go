package database

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"github.com/h3mmy/bloopyboi/bot/internal/config"
	"github.com/h3mmy/bloopyboi/ent"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go.uber.org/zap"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var logger *zap.Logger = zap.L().Named("ent_store")

func Open() (*ent.Client, error) {
	if cfg := config.GetConfig().DBConfig; cfg == nil {
		return nil, fmt.Errorf("database config is nil. No can persist")
	}
	dsnString := config.GetConfig().DBConfig.GetDSN()
	oteldb, err := otelsql.Open("pgx", dsnString)
	if err != nil {
		logger.Error("failed to create entgo client while initializing BloopyEnt")
		return nil, err
	}
	drv := entsql.OpenDB(dialect.Postgres, oteldb)
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

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			err = tx.Rollback()
			zap.L().Error(fmt.Sprintf("rollback panic recovery: %v", v), zap.Error(err))
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		return Rollback(tx, err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
