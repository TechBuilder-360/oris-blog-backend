package main

import (
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
	commentCollection := config.GetEntityDbCollection(os.Getenv("DATABASE_NAME"), os.Getenv("COMMENT_COLLECTION"))

	repoPost := _repo.NewPostRepository(postCollection)
	entityPost := _entity.NewPostEntity(repoPost)
	
	repoComment := _repo.NewCommentRepository(commentCollection)
	entityComment := _entity.NewCommentEntity(repoComment)


	api := app.Group("/api/v1/blog")

	_handler.NewPostHandler(api, entityPost)
	_handler.NewCommentHandler(api, entityComment, repoPost)

	//Port
	app.Run(port)

}
