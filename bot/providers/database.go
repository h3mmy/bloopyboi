package providers

import (
	"fmt"
	"log"
	"time"

	"gitlab.com/h3mmy/bloopyboi/bot/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func GetDatabase(dbConfig *config.BloopyDBConfig) *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s, password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
		time.Local.String(),
)

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
