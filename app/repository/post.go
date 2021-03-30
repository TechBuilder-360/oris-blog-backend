package repository

import (
	"blog/domain"
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PostRepository ...
type PostRepository struct {
	Collection *mongo.Collection
}

// NewPostRepository will create an object that represent the article.Repository interface
func NewPostRepository(Collection *mongo.Collection) domain.PostRepository {
	return &PostRepository{Collection}
}


// FetchPost ...
func (c *PostRepository) FetchPost(ctx context.Context, ginContext *gin.Context) (res []primitive.M, err error) {
	var post []bson.M

	var filter bson.M = bson.M{}

	// get by postId 
	if ginContext.Query("postid") != "" && ginContext.Query("authorid") == "" {
		id := ginContext.Query("postid")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	// get by authorID
	if ginContext.Query("authorid") != "" && ginContext.Query("postid") == "" {
		author := ginContext.Query("authorid")
		filter = bson.M{"authorid": author}
	}

	// get by author and postId
	if ginContext.Query("authorid") != "" && ginContext.Query("postid") != "" {
		author := ginContext.Query("authorid")
		id := ginContext.Query("postid")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"authorid": author, "_id": objID}
	}

	// get by author and status
	if ginContext.Query("authorid") != "" && ginContext.Query("status") != "" {
		author := ginContext.Query("authorid")
		status := ginContext.Query("status")
		filter = bson.M{"author": author, "status": status}
	}

	// get by status
	if ginContext.Query("status") != "" && ginContext.Query("authorid") == "" {
		status := ginContext.Query("status")
		filter = bson.M{"status": status}
	}

	// get by category //how to search for more than one option
	if ginContext.Query("categories") != ""{
		categories := ginContext.Query("categories")
		filter = bson.M{"categories": categories}
	}

	cur, err := c.Collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		return post, err
	}

	cur.All(context.Background(), &post)

	return post, nil
}

func (c *PostRepository) CreatePost(ctx context.Context, reqPost domain.Post) (resPost *mongo.InsertOneResult, err error) {
	var post = reqPost

	post.Slug = createSlug(post.Title)

	// https://<base_ur>/username/<slug>/
	blogBaseURL := "https://oris-blog"
	
	post.URL = fmt.Sprintf("%s/%s/%s/",blogBaseURL, url.QueryEscape(post.AuthorID), post.Slug)
	
	post.DateCreated = time.Now()
	post.Like_count = 0

	response, err := c.Collection.InsertOne(context.Background(), post)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *PostRepository) UpdatePost(ctx context.Context, id string, post domain.Post) (resPost *mongo.UpdateResult, err error) {
	
	post.DateUpdated = time.Now()

	update := bson.M{
		"$set": post,
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	response, err := c.Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *PostRepository) DeletePost(ctx context.Context, id string) (resPost *mongo.DeleteResult, err error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	response, err := c.Collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		return response, err
	}

	return response, nil
}

func createSlug(title string) (result string) {
	// remove special characters from title
	str := title
	pseudoStr := str
	for i, char := range pseudoStr {
		ascii := int(char)
		if strings.Contains(str, string(pseudoStr[i])) {
			if isAlphabet(ascii) == false {
				str = strings.Replace(str, string(str[i]), "", -1) 
			}
		}	 
	}
	str = strings.ToLower(str)

	// replace whitespaces with hypen
	result = strings.ReplaceAll(str, " ", "-")
	return
}

func isAlphabet(i int) bool {
	//lowercase check
	if (i >= 97) && (i <= 122) {
		return true
	} else if (i >= 65) && (i <= 90){ //uppercase check
		return true
	} else if i == 32 { //space character is acceptable for this check
		return true
	}

	return false
}