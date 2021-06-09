package main

import (
	"blog/config"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	_ "blog/docs" // swagger support

	_entity "blog/app/entity"
	_handler "blog/app/handler"
	_repo "blog/app/repository"

	"github.com/spf13/viper"
)


func viperConfigVariable(key string) string {

	// name of config file (without extension)
	viper.SetConfigName("config")
	// look for config in the working directory
	viper.AddConfigPath("./")
  
	// Find and read the config file
	err := viper.ReadInConfig()
  
	if err != nil {
	  log.Fatalf("Error while reading config file %s", err)
	}
  
	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	
	value, ok := viper.Get(key).(string)
  
	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
	  log.Fatalf("Invalid type assertion")
	}
	return value
}

func sanityCheck() {
	envProps := []string{
		"BLOG_COLLECTION",
		"COMMENT_COLLECTION",
		"FOLLOW_COLLECTION",
		"DATABASE_NAME",
		"DATABASE_ADDRESS",
		"PORT",
	}
	for _, k := range envProps {
		if viperConfigVariable(k) == "" {
			log.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
	fmt.Println("Environment variables successfully loaded. Starting application...")
}

func main() {
	sanityCheck()
	gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.TestMode)

	app := gin.Default()

	DATABASE_NAME, _ := viper.Get("DATABASE_NAME").(string)
	DATABASE_ADDRESS, _ := viper.Get("DATABASE_ADDRESS").(string)
	BLOG_COLLECTION, _ := viper.Get("BLOG_COLLECTION").(string)
	COMMENT_COLLECTION, _ := viper.Get("COMMENT_COLLECTION").(string)
	FOLLOW_COLLECTION, _ := viper.Get("FOLLOW_COLLECTION").(string)
	BOOKMARK_COLLECTION, _ := viper.Get("BOOKMARK_COLLECTION").(string)
	// USER_INTEREST_COLLECTION, _ := viper.Get("USER_INTEREST_COLLECTION").(string)

	postCollection := config.GetEntityDbCollection(DATABASE_NAME, DATABASE_ADDRESS, BLOG_COLLECTION)
	commentCollection := config.GetEntityDbCollection(DATABASE_NAME, DATABASE_ADDRESS, COMMENT_COLLECTION)
	followCollection := config.GetEntityDbCollection(DATABASE_NAME, DATABASE_ADDRESS, FOLLOW_COLLECTION)
	bookmarkCollection := config.GetEntityDbCollection(DATABASE_NAME, DATABASE_ADDRESS, BOOKMARK_COLLECTION)
	// userInterestCollection := config.GetEntityDbCollection(DATABASE_NAME, DATABASE_ADDRESS, USER_INTEREST_COLLECTION)

	repoPost := _repo.NewPostRepository(postCollection)
	entityPost := _entity.NewPostEntity(repoPost)
	
	repoComment := _repo.NewCommentRepository(commentCollection)
	entityComment := _entity.NewCommentEntity(repoComment)
	
	repoFollow := _repo.NewFollowRepository(followCollection)
	entityFollow := _entity.NewFollowEntity(repoFollow)

	repoBookmark := _repo.NewBookmarkRepository(bookmarkCollection)
	entityBookmark := _entity.NewBookmarkEntity(repoBookmark)

	// repoUserInterest := _repo.NewUserInterestRepository(userInterestCollection)
	// entityUserInterest := _entity.NewUserInterestEntity(repoUserInterest)


	app.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "docs/")
	})
	
	app.StaticFS("/docs", http.Dir("./templates"))

	api := app.Group("/api/v1/blog")

	_handler.NewPostHandler(api, entityPost)
	_handler.NewCommentHandler(api, entityComment, repoPost)
	_handler.NewFollowHandler(api, entityFollow, repoFollow)
	_handler.NewBookmarkHandler(api, entityBookmark)

	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	app.Use(cors.New(config))

	//Port
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	vPORT, _ := viper.Get("PORT").(string)

    if PORT == ":" {
        PORT = vPORT
    } 

	app.Run(PORT)
}