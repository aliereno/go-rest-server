package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Email        string  `json:"email" gorm:"not null;unique_index:idx_email"`
	PasswordHash string  `json:"-" gorm:"column:password;not null"`
	Name         string  `json:"name"`
	Role         int     `json:"-" gorm:"default:0"`
	Likes    	 []*Book `json:"likes" gorm:"many2many:user_book_likes;association_jointable_foreignkey:book_id"`
	Dislikes 	 []*Book `json:"dislikes" gorm:"many2many:user_book_dislikes;association_jointable_foreignkey:book_id"`
	Purchases	 []*Book `json:"purchases" gorm:"many2many:user_book_purchases;association_jointable_foreignkey:book_id"`
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
