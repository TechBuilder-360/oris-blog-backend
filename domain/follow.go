package domain

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Follow struct
type Follow struct {
	UserID string `json:"userid,omitempty"`
	Followers [] string `json:"followers,omitempty"`
	Following [] string `json:"following,omitempty"`
}

type FollowEntity interface {

	Follow(ctx context.Context, follow Follow) (*mongo.InsertManyResult, error)
	UpdateFollowers(ctx context.Context, userId string, followerId string) (resFollow *mongo.UpdateResult, err error)
	UpdateFollowing(ctx context.Context, userId string, followedId string) (resFollow *mongo.UpdateResult, err error)
	// UnFollow(ctx context.Context, id string, follow Follow) (*mongo.UpdateResult, error)

	FetchFollows(ctx context.Context, ginContext *gin.Context) (res []primitive.M, err error)
	DeleteFollowRecord(ctx context.Context, commentid string) (*mongo.DeleteResult, error)
}

// PostRepository interface
type FollowRepository interface {

	CreateFollow(ctx context.Context, follow Follow) (*mongo.InsertManyResult, error)
	UpdateFollowers(ctx context.Context, userId string, followerId string) (resFollow *mongo.UpdateResult, err error)
	UpdateFollowing(ctx context.Context, userId string, followedId string) (resFollow *mongo.UpdateResult, err error)
	
	// DeleteFollow(ctx context.Context, id string, follow Follow) (*mongo.UpdateResult, error)

	FetchFollows(ctx context.Context, ginContext *gin.Context) (res []primitive.M, err error)

	ValidateRelationshipExistence(ctx context.Context, userID string, followerID string) bool
	ValidateUserRecordExistence(ctx context.Context, userID string) bool
	DeleteFollowRecord(ctx context.Context, commentid string) (*mongo.DeleteResult, error)

}