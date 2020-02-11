package auth

import (
	"github.com/aliereno/go-rest-server/internal/orm/models"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	mysupersecretpassword = "unicornsAreAwesome"
)

func GetToken(user models.User) (string, error) {
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	token.Claims = jwt_lib.MapClaims{
		"user_id": user.ID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // TOKEN EXPIRE TIME
	}
	tokenString, err := token.SignedString([]byte(mysupersecretpassword))
	if err != nil {
		return "Could not generate token", err
	}
	return tokenString, nil
}

func LookUserTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := []byte(mysupersecretpassword)
			userId := token.Claims.(jwt_lib.MapClaims)["user_id"]
			c.Set("user_id", userId)
			return b, nil
		})
		if err != nil {
			_ = c.AbortWithError(401, err)
		}
	}
}
