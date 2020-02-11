package jobs

import (
	"fmt"
	"github.com/aliereno/go-rest-server/internal/logger"
	"github.com/aliereno/go-rest-server/internal/orm/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

// SeedUsers inserts the first users
var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		name := ""
		email := ""
		for i := 1; i < 3; i++ {
			name = fmt.Sprintf("User%d", i)
			email = fmt.Sprintf("user%d@gmail.com", i)
			dbo := &models.User{
				Name:  name,
				Email: email,
			}
			if err := dbo.SetPassword("123"); err != nil {
				logger.Error("password 123")
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
