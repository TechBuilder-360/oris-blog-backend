package repository

import (
	"blog/domain"
	"context"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CommentRepository ...
type CommentRepository struct {
	Collection *mongo.Collection
}

// NewCommentRepository will create an object that represent the article.Repository interface
func NewCommentRepository(Collection *mongo.Collection) domain.CommentRepository {
	return &CommentRepository{Collection}
}


// FetchComment ...
func (c *CommentRepository) FetchComment(ctx context.Context, ginContext *gin.Context) (res []primitive.M, err error) {
	var comment []bson.M

	var filter bson.M = bson.M{}

	// get by commentId 
	if ginContext.Query("commentid") != "" && ginContext.Query("authorid") == "" {
		id := ginContext.Query("commentid")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	// get by authorID
	if ginContext.Query("authorid") != "" && ginContext.Query("commentid") == "" {
		author := ginContext.Query("authorid")
		filter = bson.M{"authorid": author}
	}

	// get by author and commentId
	if ginContext.Query("authorid") != "" && ginContext.Query("commentid") != "" {
		author := ginContext.Query("authorid")
		id := ginContext.Query("commentid")
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
		return comment, err
	}

	cur.All(context.Background(), &comment)

	return comment, nil
}

func (c *CommentRepository) CreateComment(ctx context.Context, reqComment domain.Comment) (resComment *mongo.InsertOneResult, err error) {
	reqComment.DateCreated = time.Now()

	response, err := c.Collection.InsertOne(context.Background(), reqComment)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *CommentRepository) UpdateComment(ctx context.Context, id string, comment domain.Comment) (resComment string, err error) {
	
	// comment.DateUpdated = time.Now()

	update := bson.M{
		"$set": comment,
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	response, err := c.Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		return "FAILED", err
	}

	if response.ModifiedCount == 0 {
		return "FAILED", errors.New("update failed")
	}
	return "SUCCESS", nil
}

func (c *CommentRepository) DeleteComment(ctx context.Context, commentid string) (resComment string, err error) {
	commentObjID, _ := primitive.ObjectIDFromHex(commentid)

	response, err := c.Collection.DeleteOne(context.Background(), bson.M{"_id": commentObjID})

	if err != nil {
		return "FAILED", err
	}

	if response.DeletedCount == 0 {
		return "FAILED", errors.New("delete failed")
	}

	return "SUCCESS", nil
}