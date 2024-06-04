package dbConnection

import (
	"deli-blog/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN = "host=localhost user=postgres password=deliblogpassword dbname=deliblogdb port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"

func New() *gorm.DB {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(">> Connection fail: ", err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Post{},
	)
}

func MigrateBreakingAddress(db *gorm.DB) {
	if db.Migrator().HasTable(&model.User{}) {
		db.Transaction(func(tx *gorm.DB) error {
			if err := db.Model(&model.User{}).Where("address IS NULL").Update("address", "None").Error; err != nil {
				return err
			}
			return nil
		})
	}
}
