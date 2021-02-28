package main

import (
	"blog/models"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
)

var port string = os.Getenv("PORT")

func main() {

	app := gin.Default()

	//queryparams postId,author,category, status
	app.GET("/api/blog/post", models.GetPost) 

	app.POST("/api/blog/post", models.CreatePost)

	app.PUT("/api/blog/post/:id", models.UpdatePost)

	app.DELETE("/api/blog/post/:id", models.DeletePost)

	//Port
	app.Run(port)

}