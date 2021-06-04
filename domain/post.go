package domain

import (
	"context"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Post struct
type Post struct {
	_id            string
	AuthorID       string `json:"authorID"`
	Title          string `json:"title"`
	Summary		   string
	CoverImage	   string `json:"coverimage"`
	Slug           string
	URL            string
	Categories     []string    `json:"categories"`
	Likes          []string// user Id array
	Like_count     int
	Comments       []string // comment Id array
	Comments_count int
	Article        string `json:"article"`
	DateCreated    time.Time
	DateUpdated    time.Time
	Status         string `json:"status"`
	ReadTime	   string
}

// PostEntity interface
type PostEntity interface {
	CreatePost(ctx context.Context, post Post) (*mongo.InsertOneResult, error)
	UpdatePost(ctx context.Context, id string, post Post) (string, error)
	ValidatePostExistence(ctx context.Context, authorid string, postid string) bool
	DeletePost(ctx context.Context, id string) (string, error)

	FetchPost(ctx context.Context, c *gin.Context) ([]primitive.M, error)
	LikePost(ctx context.Context, userid string, postid string, mode string) (string, error)
}

// PostRepository interface
type PostRepository interface {
	CreatePost(ctx context.Context, post Post) (*mongo.InsertOneResult, error)
	UpdatePost(ctx context.Context, id string, post Post) (string, error)
	DeletePost(ctx context.Context, id string) (string, error)

	FetchPost(ctx context.Context, c *gin.Context) ([]primitive.M, error)
	LikePost(ctx context.Context, userid string, postid string, mode string) (string, error)

	ValidatePostExistence(ctx context.Context, authorid string, postid string) bool
	InsertPostComment(ctx context.Context, commentID *mongo.InsertOneResult, postid string) (resPost *mongo.UpdateResult, err error)
	RemovePostComment(ctx context.Context, postid string, commentid string) (resPost *mongo.UpdateResult, err error)
}
