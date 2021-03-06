package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Comment struct
type Comment struct {
	_id string
	PostID string `json:"postid"`
	Author string `json:"authorid"` // comment author id/name
	Text string `json:"text"`
	DateCreated time.Time
	Replies [] string //userID array
	RepliesCount int
	Likes [] string  //userID array
	LikeCount int
}


// PostEntity interface
type CommentEntity interface {

	CreateComment(ctx context.Context, comment Comment) (*mongo.InsertOneResult, error)
	UpdateComment(ctx context.Context, id string, comment Comment) (string, error)
	DeleteComment(ctx context.Context, commentid string) (string, error)

	FetchComment(ctx context.Context, c *gin.Context) ([]primitive.M, error)
}

// PostRepository interface
type CommentRepository interface {

	CreateComment(ctx context.Context, comment Comment) (*mongo.InsertOneResult, error)
	UpdateComment(ctx context.Context, id string, comment Comment) (string, error)
	DeleteComment(ctx context.Context, commentid string) (string, error)

	FetchComment(ctx context.Context, c *gin.Context) ([]primitive.M, error)
}