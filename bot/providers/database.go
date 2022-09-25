package providers

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"gitlab.com/h3mmy/bloopyboi/bot/internal/config"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

var (
	sqlitePath string = "/var/go-bloopyboi/local.sqlite"
)

type BloopyDBManager struct {
	Type        string
	DB          *gorm.DB
	Logger      *zap.Logger
	DbLogConfig glogger.Interface
}

func NewBloopyDBManager(botConfig *config.BotConfig) *BloopyDBManager {
	// Increase verbosity of the database if the loglevel is higher than Info
	var logConfig glogger.Interface
	if GetLogLevel() > "DEBUG" {
		logConfig = glogger.Default.LogMode(glogger.Info)
	}
	return &BloopyDBManager{
		Logger:      logger,
		DbLogConfig: logConfig,
	}
}

// Initialized postgresDB connection
func (dbMgr *BloopyDBManager) WithPostgresDatabase(dbConfig *config.BloopyDBConfig) (*BloopyDBManager, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s, password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
		time.Local.String(),
	)

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{Logger: dbMgr.DbLogConfig})

	if err != nil {
		dbMgr.Logger.Sugar().Error("Unable to initialize postgres session", err)
		return dbMgr, err
	}

	dbMgr.Type = "postgres"
	dbMgr.DB = db
	return dbMgr, err
}

// Initializes sqliteDB. To be used for development
func (dbMgr *BloopyDBManager) WithSqliteDatabase() (*BloopyDBManager, error) {
	var sqlitePath string
	// Create the folder path if it doesn't exist
	_, err := os.Stat(sqlitePath)
	if errors.Is(err, fs.ErrNotExist) {
		dirPath := filepath.Dir(sqlitePath)
		if err := os.MkdirAll(dirPath, 0660); err != nil {
			dbMgr.Logger.Sugar().Error("unable to make directory path ", dirPath, " err: ", err)
			sqlitePath = "./local.db"
			return dbMgr, err
		}
	}
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{Logger: dbMgr.DbLogConfig})

	if err != nil {
		dbMgr.Logger.Sugar().Error("Unable to initialize sqlite: ", err)
		return dbMgr, err
	}

	dbMgr.Type = "sqlite"

	if err := db.AutoMigrate(&models.User{}); err != nil {
		dbMgr.Logger.Sugar().Error("Unable to migrate schemas: ", err)
		return dbMgr, err
	}

	dbMgr.DB = db
	return dbMgr, err
}

// Returns DB Ref
// Will do it naively atm, todo add nullcheck
func (dbMgr *BloopyDBManager) GetDB() (*gorm.DB, error) {
	if dbMgr.DB != nil {
		return dbMgr.DB, nil
	} else {
		return nil, errors.New("No Database Initialized")
	}

}
