package jobs

import (
	"fmt"
	"github.com/aliereno/go-rest-server/internal/orm/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var SeedBooks *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_BOOKS",
	Migrate: func(db *gorm.DB) error {
		title := ""
		for i := 1; i < 3; i++ {
			title = fmt.Sprintf("Book Title %d", i)
			dbo := &models.Book{
				Price:     100,
				Title:     title,
			}
			tx := db.New().Begin()
			tx = tx.Create(dbo).First(dbo)
			tx = tx.Commit()
		}
		return nil
	},
	Rollback: func(db *gorm.DB) error {
		return nil
	},
}
