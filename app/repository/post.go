package repository

import (
	"blog/domain"
	"context"
	"errors"
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
	if ginContext.Query("categories") != "" {
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

	post.URL = fmt.Sprintf("%s/%s/%s/", blogBaseURL, url.QueryEscape(post.AuthorID), post.Slug)

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

func (c *PostRepository) InsertPostComment(ctx context.Context, commentID *mongo.InsertOneResult, postid string) (resPost *mongo.UpdateResult, err error) {

	var post domain.Post
	objID, _ := primitive.ObjectIDFromHex(postid)

	err = c.Collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&post)

	if err != nil {
		return nil, err
	}

	post.Comments = append(post.Comments, fmt.Sprintf("%v", commentID.InsertedID))
	post.Comments_count = len(post.Comments)

	update := bson.M{
		"$set": post,
	}

	response, err := c.Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *PostRepository) RemovePostComment(ctx context.Context, postid string, commentid string) (resPost *mongo.UpdateResult, err error) {
	var post domain.Post
	postObjID, _ := primitive.ObjectIDFromHex(postid)
	commentObjID, _ := primitive.ObjectIDFromHex(commentid)

	err = c.Collection.FindOne(context.Background(), bson.M{"_id": postObjID}).Decode(&post)

	if err != nil {
		return nil, err
	}

	commentIndex := indexOf(fmt.Sprintf("%v", commentObjID), post.Comments)

	errMessage := errors.New("Comment does not exist for this post")
	if commentIndex == -1 {
		return nil, errMessage
	}

	post.Comments = append(post.Comments[:commentIndex], post.Comments[commentIndex+1:]...)

	post.Comments_count = len(post.Comments) - 1

	update := bson.M{
		"$set": post,
	}

	response, err := c.Collection.UpdateOne(context.Background(), bson.M{"_id": postObjID}, update)

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

func (c *PostRepository) ValidatePostExistence(ctx context.Context, postid string) bool {
	var post domain.Post
	objID, _ := primitive.ObjectIDFromHex(postid)
	err := c.Collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&post)

	if err != nil {
		return false
	}

	return true
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
	} else if (i >= 65) && (i <= 90) { //uppercase check
		return true
	} else if i == 32 { //space character is acceptable for this check
		return true
	}

	return false
}

func indexOf(element string, data []string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
 }