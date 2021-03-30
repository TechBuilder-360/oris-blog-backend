package main

import (
	// "blog/models"
	"blog/config"
	"os"

	// "github.com/gin-gonic/contrib/cors"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"

	_ "blog/docs" // swagger support

	_repo "blog/app/repository"
	_entity "blog/app/entity"
	_handler "blog/app/handler"
)

var port string = os.Getenv("PORT")

func main() {
	// gin.SetMode(gin.ReleaseMode)

	app := gin.Default()

	// app.Use(cors.Default())

	postCollection := config.GetEntityDbCollection(os.Getenv("DATABASE_NAME"), os.Getenv("BLOG_COLLECTION"))

	repoPost := _repo.NewPostRepository(postCollection)
	entityPost := _entity.NewPostEntity(repoPost)	

	//queryparams postId,author,category, status
	// app.GET("/api/blog/post", models.GetPost)

	// app.POST("/api/blog/post", models.CreatePost)

	// app.PUT("/api/blog/post/:id", models.UpdatePost)

	// app.DELETE("/api/blog/post/:id", models.DeletePost)

	api := app.Group("/api/v1/blog")

	_handler.NewBooksHandler(api, entityPost)

	//Port
	app.Run(port)

}
