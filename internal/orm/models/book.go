package models

type Book struct {
	BaseModel
	Price      int        `json:"price"`
	Title      string     `json:"title"`
	Content    *string    `json:"content"`
	Likes      []*User    `json:"likes" gorm:"many2many:user_book_likes;association_jointable_foreignkey:user_id"`
	Dislikes   []*User    `json:"dislikes" gorm:"many2many:user_book_dislikes;association_jointable_foreignkey:user_id"`
}
