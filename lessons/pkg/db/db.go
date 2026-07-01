package db

import (
	"context"
	"go/adv-demo/configs"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database := &Db{db}

	if err := database.CheckConnection(); err != nil {
		panic("database connection failed" + err.Error())
	}
	return database
}

func (d *Db) CheckConnection() error {
	// 1. Получаем стандартный *sql.DB из GORM
	sqlDb, err := d.DB.DB()
	if err != nil {
		return err
	}

	// 2. Создаем контекст с таймаутом, чтобы пинг не завис бесконечно,
	// если сеть «моргнула» или база лежит
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 3. Делаем Ping
	return sqlDb.PingContext(ctx)
}
