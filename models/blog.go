// package models

// import (
// 	"blog/database"
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// // Post struct
// type Post struct {
// 	id string 
// 	Author string `json:"author,omitempty"`
// 	Title string `json:"title,omitempty"`
// 	// CoverImage 
// 	Slug string
// 	URL string 
// 	Categories []string `json:"categories,omitempty"`
// 	Likes int `json:"likes,omitempty"`
// 	Comments []Comment `json:"comments,omitempty"`
// 	Article string `json:"article,omitempty"`
// 	DateCreated time.Time
// 	DateUpdated time.Time
// 	Status string `json:"status,omitempty"`
// }

// /////////////////////////////////////////////////////////////////////////////////////////////////////

// // PostUsecase interface
// type PostUsecase interface {
// 	Fetch(ctx context.Context) ([]Post, error)
// 	GetByID(ctx context.Context, id string) (Post, error)
// }
// // PostRepository interface
// type PostRepository interface {
// 	Fetch(ctx context.Context) (res []Post, err error)
// 	GetByID(ctx context.Context, id string) (Post, error)
// }

// /////////////////////////////////////////////////////////////////////////////////////////////////////


// func getCollectionObject()(*mongo.Collection){
// 	collection := database.GetMongoDbCollection(os.Getenv("DATABASE_NAME"), os.Getenv("BLOG_COLLECTION"))
// 	return collection
// }

// func init (){
// 	Collection = database.GetMongoDbCollection(os.Getenv("DATABASE_NAME"), os.Getenv("BLOG_COLLECTION"))
// }

// // Collection blah blah
// var Collection *mongo.Collection

// // Collection blah blah
// // var Collection = database.GetMongoDbCollection(os.Getenv("DATABASE_NAME"), os.Getenv("BLOG_COLLECTION"))

// //GetPost returns a single post or all posts as the case may be.
// func GetPost(c *gin.Context) {
	
// 	var filter bson.M = bson.M{}

// 	// get by postId
// 	if c.Query("postId") != "" && c.Query("author") == "" {
// 		id := c.Query("postId")
// 		objID, _ := primitive.ObjectIDFromHex(id)
// 		filter = bson.M{"_id": objID}
// 	}

// 	// get by author
// 	if c.Query("author") != "" && c.Query("postId") == "" {
// 		author := c.Query("author")
// 		filter = bson.M{"author": author}
// 	}

// 	// get by author and postId
// 	if c.Query("author") != "" && c.Query("postId") != "" {
// 		author := c.Query("author")
// 		id := c.Query("postId")
// 		objID, _ := primitive.ObjectIDFromHex(id)
// 		filter = bson.M{"author": author, "_id": objID}
// 	}

// 	// get by author and status
// 	if c.Query("author") != "" && c.Query("status") != "" {
// 		author := c.Query("author")
// 		status := c.Query("status")
// 		filter = bson.M{"author": author, "status": status}
// 	}

// 	// get by category
// 	if c.Query("category") != ""{
// 		category := c.Query("category")
// 		filter = bson.M{"category": category}
// 	}

// 	var results []bson.M
// 	cur, err := Collection.Find(context.Background(), filter)
// 	defer cur.Close(context.Background())

// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	cur.All(context.Background(), &results)

// 	if results == nil {
// 		c.JSON(http.StatusNotFound, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, results)
// }

// // CreatePost record into collection
// func CreatePost(c *gin.Context) {

// 	var post Post 

// 	err := c.ShouldBind(&post)

// 	// What if not all required details are posted?
// 	// Post will still be successful, but this does not conform with the app requirement
// 	// Thus, a validation should be in place for this 

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	post.Slug = createSlug(post.Title)

// 	// https://<base_ur>/username/<slug>/
// 	blogBaseURL := "https://oris-blog"
	
// 	post.URL = fmt.Sprintf("%s/%s/%s/",blogBaseURL, url.QueryEscape(post.Author), post.Slug)
	
// 	post.DateCreated = time.Now()
// 	post.Likes = 0

// 	response, err := Collection.InsertOne(context.Background(), post)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// //UpdatePost record in collection
// func UpdatePost(c *gin.Context) {
// 	var post Post
// 	c.ShouldBind(&post)

// 	post.DateUpdated = time.Now()

// 	update := bson.M{
// 		"$set": post,
// 	}

// 	objID, _ := primitive.ObjectIDFromHex(c.Param("id"))
// 	response, err := Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// func createSlug(title string) (result string) {
// 	// remove special characters from title
// 	str := title
// 	pseudoStr := str
// 	for i, char := range pseudoStr {
// 		ascii := int(char)
// 		if strings.Contains(str, string(pseudoStr[i])) {
// 			if isAlphabet(ascii) == false {
// 				str = strings.Replace(str, string(str[i]), "", -1) 
// 			}
// 		}	 
// 	}
// 	str = strings.ToLower(str)

// 	// replace whitespaces with hypen
// 	result = strings.ReplaceAll(str, " ", "-")
// 	return
// }

// func isAlphabet(i int) bool {
// 	//lowercase check
// 	if (i >= 97) && (i <= 122) {
// 		return true
// 	} else if (i >= 65) && (i <= 90){ //uppercase check
// 		return true
// 	} else if i == 32 { //space character is acceptable for this check
// 		return true
// 	}

// 	return false
// }

// // DeletePost record from collection
// func DeletePost(c *gin.Context) {
// 	objID, _ := primitive.ObjectIDFromHex(c.Param("id"))
// 	response, err := Collection.DeleteOne(context.Background(), bson.M{"_id": objID})

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }


