package database

import (
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/access-control-system/model"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(filepath.Join("database", "main.db")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{}, &model.Record{})

	return db, nil
}
