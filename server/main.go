package main

import (
	"github.com/aliereno/go-rest-server/internal/handlers/auth"
	controllersMain "github.com/aliereno/go-rest-server/internal/handlers/controllers"
	log "github.com/aliereno/go-rest-server/internal/logger"
	"github.com/aliereno/go-rest-server/internal/orm"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var host, port string

func init() {
	host = "localhost"
	port = "7777"
}

func main() {
	orm, err := orm.Factory()
	controllers := controllersMain.Controller{ORM: orm}
	if err != nil {
		log.Panic(err)
	}
	endpoint := "http://" + host + ":" + port
	r := gin.Default()

	r.Use(cors.Default())

	r.Static("/files", "./files")

	r.POST("/login", controllers.UserLogin())
	r.POST("/register", controllers.UserRegister())

	r.GET("/books", controllers.BooksFetchAll())
	r.GET("/books/:id", controllers.BookFetchDetail())

	api := r.Group("/user")
	api.Use(auth.LookUserTokenHandler())
	{
		api.GET("/settings", controllers.UserFetchSetting())
		api.POST("/settings", controllers.UserUpdateSetting())
	}

	log.Info("Running @ " + endpoint)
	log.Fatal(r.Run(host + ":" + port))
}
