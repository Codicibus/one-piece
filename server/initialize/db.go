package initialize

import (
	"fmt"
	"log"
	"opiece/server/global"
	"opiece/server/model"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func NewGorm() *gorm.DB {
	switch global.OConfig.DB.Type {
	case "postgres":
		return GormPostgres()
	case "mysql":
		return GormMYSQL()
	default:
		return nil
	}
}

func GormMYSQL() *gorm.DB {
	return nil
}

func MigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(model.User{})
	if err != nil {
		log.Fatal(err)
	}
}

func GormPostgres() *gorm.DB {
	pg := global.OConfig.Postgres
	if pg.DBName == "" {
		return nil
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%v TimeZone=%s",
		pg.DBHost, pg.DBUsername, pg.DBPassword, pg.DBName, pg.DBPort, pg.SSLMode, pg.Timezone)
	pgc := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(pgc, &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("connect to postgres failed: %s", err.Error()))
	}
	return db
}
