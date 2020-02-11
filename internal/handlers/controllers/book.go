package controllers

import (
	"github.com/aliereno/go-rest-server/internal/orm/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (con Controller) BooksFetchAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := con.ORM.DB.New()
		var dbRecords []*models.Book
		db = db.Preload("Likes").Preload("Dislikes").Find(&dbRecords)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dbRecords})
	}
}

func (con Controller) BookFetchDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		db := con.ORM.DB.New()
		var dbRecord models.Book
		err := db.Preload("Likes").Preload("Dislikes").Where("id = ?", id).Find(&dbRecord).Error
		if err == nil {
			db.Save(&dbRecord)
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dbRecord})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": ""})
		}
	}
}