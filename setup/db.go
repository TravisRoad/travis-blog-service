package setup

import (
	"log/slog"

	"github.com/TravisRoad/travis-blog-service/global"
	"github.com/TravisRoad/travis-blog-service/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	config := global.Config.Sqlite
	db, err := gorm.Open(sqlite.Open(config.Path), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	db.AutoMigrate(
		&model.Like{},
		&model.LikeOp{},
	)

	global.DB = db
}
