package domain

import (
	"context"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

// Post struct
type Post struct {
	_id            string
	AuthorID       string `json:"authorID,omitempty"`
	Title          string `json:"title,omitempty"`
	Slug           string
	URL            string
	Categories     []string    `json:"categories,omitempty"`
	Likes          []uuid.UUID // user Id array
	Like_count     int
	Comments       []string // comment Id array
	Comments_count int
	Article        string `json:"article,omitempty"`
	DateCreated    time.Time
	DateUpdated    time.Time
	Status         string `json:"status,omitempty"`
}

// PostEntity interface
type PostEntity interface {
	CreatePost(ctx context.Context, post Post) (*mongo.InsertOneResult, error)
	UpdatePost(ctx context.Context, id string, post Post) (*mongo.UpdateResult, error)
	DeletePost(ctx context.Context, id string) (*mongo.DeleteResult, error)

	FetchPost(ctx context.Context, c *gin.Context) ([]primitive.M, error)
}

// PostRepository interface
type PostRepository interface {
	CreatePost(ctx context.Context, post Post) (*mongo.InsertOneResult, error)
	UpdatePost(ctx context.Context, id string, post Post) (*mongo.UpdateResult, error)
	DeletePost(ctx context.Context, id string) (*mongo.DeleteResult, error)

	FetchPost(ctx context.Context, c *gin.Context) ([]primitive.M, error)

	ValidatePostExistence(ctx context.Context, postid string) bool
	InsertPostComment(ctx context.Context, commentID *mongo.InsertOneResult, postid string) (resPost *mongo.UpdateResult, err error)
	RemovePostComment(ctx context.Context, postid string, commentid string) (resPost *mongo.UpdateResult, err error)
}
