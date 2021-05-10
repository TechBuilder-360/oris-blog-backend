package domain

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserInterest struct
type UserInterest struct {
	UserID string `json:"userid"`
	Categories [] string `json:"categories"`
}

type UserInterestEntity interface {

	AddToUserInterest(ctx context.Context, userId string, followerId string, mode string) (resFollow *mongo.UpdateResult, err error)
	RemoveFromUserInterest(ctx context.Context, userId string, followedId string, mode string) (resFollow *mongo.UpdateResult, err error)

	FetchUserInterests(ctx context.Context, ginContext *gin.Context) (res []primitive.M, err error)
	DeleteUserInterestRecord(ctx context.Context, commentid string) (*mongo.DeleteResult, error)
}

// PostRepository interface
type UserInterestRepository interface {

	AddToUserInterest(ctx context.Context, userId string, followerId string, mode string) (resFollow *mongo.UpdateResult, err error)
	RemoveFromUserInterest(ctx context.Context, userId string, followedId string, mode string) (resFollow *mongo.UpdateResult, err error)

	FetchUserInterests(ctx context.Context, ginContext *gin.Context) (res []primitive.M, err error)
	DeleteUserInterestRecord(ctx context.Context, commentid string) (*mongo.DeleteResult, error)

}