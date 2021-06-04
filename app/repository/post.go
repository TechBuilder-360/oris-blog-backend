package repository

import (
	"blog/domain"
	"blog/util"
	"context"
	"errors"
	"fmt"

	"net/url"
	"strings"
	"time"

	"github.com/JesusIslam/tldr"
	readingtime "github.com/begmaroman/reading-time"
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
	if ginContext.Query("authorid") != "" {
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
		filter = bson.M{"authorid": author, "status": status}
	}

	// get by author and slug
	if ginContext.Query("authorid") != "" && ginContext.Query("slug") != "" {
		author := ginContext.Query("authorid")
		slug := ginContext.Query("slug")
		filter = bson.M{"authorid": author, "slug": slug}
	}

	// get by status
	if ginContext.Query("status") != "" && ginContext.Query("authorid") == "" {
		status := ginContext.Query("status")
		filter = bson.M{"status": status}
	}

	// get by category //how to search for more than one option
	if ginContext.Query("category") != "" {
		categories := ginContext.Query("category")
		filter = bson.M{"categories": categories}
	}

	cur, err := c.Collection.Find(context.Background(), filter)

	if err != nil {
		return post, err
	}
	defer cur.Close(context.Background())

	cur.All(context.Background(), &post)

	return post, nil
}

func (c *PostRepository) CreatePost(ctx context.Context, reqPost domain.Post) (resPost *mongo.InsertOneResult, err error) {
	var post = reqPost

	post.Slug = createSlug(post.Title)

	// https://<base_ur>/username/<slug>/
	blogBaseURL := "https://oris-blog"

	// URL is unused and possibly removed from the post model in the future. A post is gotten by author and slug
	post.URL = fmt.Sprintf("%s/%s/%s/", blogBaseURL, url.QueryEscape(post.AuthorID), post.Slug)

	post.DateCreated = time.Now()
	post.DateUpdated = post.DateCreated
	post.Like_count = 0
	post.ReadTime = fmt.Sprintf("%v", readingtime.Estimate(post.Article).Text)
	
	intoSentences := 2
	bag := tldr.New()
	result, _ := bag.Summarize(post.Article, intoSentences)
	post.Summary = result[0]

	response, err := c.Collection.InsertOne(context.Background(), post)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *PostRepository) UpdatePost(ctx context.Context, id string, post domain.Post) (status string, err error) {

	post.DateUpdated = time.Now()

	update := bson.M{
		"$set": post,
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	response, err := c.Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		return "FAILED", err
	}

	if response.ModifiedCount == 0 {
		// update failed
		return "FAILED", errors.New("update failed")
	}

	return "SUCCESS", nil
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

	// commentIndex := indexOf(fmt.Sprintf("%v", commentObjID), post.Comments)
	_, commentIndex := util.Find(post.Comments, fmt.Sprintf("%v", commentObjID))

	errMessage := errors.New("comment does not exist for this post")

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

func (c *PostRepository) DeletePost(ctx context.Context, id string) (status string, err error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err = c.Collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		return "FAILED", err
	}

	return "SUCCESS", nil
}

func (c *PostRepository) ValidatePostExistence(ctx context.Context, authorid string, postid string) bool {
	var post domain.Post
	objID, _ := primitive.ObjectIDFromHex(postid)
	var filter bson.M = bson.M{"_id": objID}

	if authorid != ""{
		filter = bson.M{"_id": objID, "authorid" : authorid}
	}
	err := c.Collection.FindOne(context.Background(), filter).Decode(&post)

	return  err == nil
}

func (c *PostRepository) LikePost(ctx context.Context, userId string, postid string, mode string) (string, error){
	var post domain.Post
	postObjID, _ := primitive.ObjectIDFromHex(postid)

	err := c.Collection.FindOne(context.Background(), bson.M{"_id": postObjID}).Decode(&post)
	isElement, index := util.Find(post.Likes, userId)

	if err != nil {
		return "FAILED", err
	}

	if mode == "add"{

		if isElement {
			return "FAILED", errors.New("post liked already")
		}

		post.Likes = append(post.Likes, userId)
	}else if mode == "remove" {
		
		if isElement {
			post.Likes = util.RemoveIndex(post.Likes, index)
		}
	}
	
	post.Like_count = len(post.Likes)

	update := bson.M{
		"$set": post,
	}

	_, err = c.Collection.UpdateOne(context.Background(), bson.M{"_id": postObjID}, update)

	if err != nil {
		return "FAILED", err
	}

	return "SUCCESS", nil
}

func createSlug(title string) (result string) {
	// remove special characters from title
	str := title
	pseudoStr := str
	for i, char := range pseudoStr {
		ascii := int(char)
		if strings.Contains(str, string(pseudoStr[i])) {
			if !isAlphabet(ascii) {
				str = strings.Replace(str, string(str[i]), "", -1)
			}
		}
	}
	str = strings.ToLower(str)

	// replace whitespaces with hypen
	result = strings.ReplaceAll(str, " ", "-")
	
	split := strings.Split(time.Now().String(), ".")
	x := split[0]
	strCon := strings.Replace(x, "T", "", 1)
	strCon = strings.Replace(strCon, "-", "", 2)
	strCon = strings.Replace(strCon, ":", "", 2)
	strCon = strings.Replace(strCon, " ", "", 2)

	result = fmt.Sprintf("%s-%s",result, strCon)
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