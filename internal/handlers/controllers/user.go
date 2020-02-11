package controllers

import (
	"github.com/aliereno/go-rest-server/internal/handlers/auth"
	"github.com/aliereno/go-rest-server/internal/orm/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (con Controller) UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginModel struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		err := c.BindJSON(&loginModel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		user := models.User{}
		db := con.ORM.DB.New()
		er := db.Where("email = ?", loginModel.Email).First(&user).Error
		if er != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
			return
		}
		if err := user.CheckPassword(loginModel.Password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		token, errr := auth.GetToken(user)
		if errr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "token": token, "id": user.ID})
	}
}

func (con Controller) UserRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerModel struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		err := c.BindJSON(&registerModel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		dbo := &models.User{
			Name:  registerModel.Name,
			Email: registerModel.Email,
		}
		if err := dbo.SetPassword(registerModel.Password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		tx := con.ORM.DB.New().Begin()
		tx = tx.Create(dbo).First(dbo)
		tx = tx.Commit()
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}

func (con Controller) UserFetchSetting() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId,flag := c.Get("user_id")
		if flag{
			db := con.ORM.DB.New()
			var dbRecord models.User
			err := db.Where("id = ?", userId).Preload("Likes").Preload("Dislikes").Preload("Purchases").Find(&dbRecord).Error
			if err == nil {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dbRecord})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			}
		}else{
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error":"Token error."})
		}
	}
}

func (con Controller) UserUpdateSetting() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "on development..."})
	}
}