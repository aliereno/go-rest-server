package models

import (
	_ "github.com/jinzhu/gorm"
	"time"
)

// BaseModel defines the common columns that all db structs should hold, usually
// db structs based on this have no soft delete
type BaseModel struct {
	ID        int        `json:"id";gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time  `json:"created_at";gorm:"index;not null;default:CURRENT_TIMESTAMP"` // (My|Postgre)SQL
	UpdatedAt *time.Time `json:"-";gorm:"index"`
}
