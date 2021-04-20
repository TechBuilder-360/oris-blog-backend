package entity

import (
	"blog/domain"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// FollowEntity ...
type FollowEntity struct {
	followRepo domain.FollowRepository
}

// NewFollowEntity will create new an articleUsecase object representation of domain.FollowEntity interface
func NewFollowEntity(a domain.FollowRepository) domain.FollowEntity {
	return &FollowEntity{
		followRepo: a,
	}
}

func (a *FollowEntity) Follow(ctx context.Context, follow domain.Follow) (res *mongo.InsertManyResult, err error){
	res, err = a.followRepo.CreateFollow(ctx, follow)
	return 
}


func (a *FollowEntity) FetchFollows(ctx context.Context, ginContext *gin.Context) (res []primitive.M, err error){
	res, err = a.followRepo.FetchFollows(ctx, ginContext)
	return 
}


func (a *FollowEntity) UpdateFollowers(ctx context.Context, userId string, followerId string) (res *mongo.UpdateResult, err error){
	res, err = a.followRepo.UpdateFollowers(ctx, userId, followerId)
	return 
}

func (a *FollowEntity) UpdateFollowing(ctx context.Context, userId string, followedId string) (res *mongo.UpdateResult, err error){
	res, err = a.followRepo.UpdateFollowing(ctx, userId, followedId)
	return 
}


func (a *FollowEntity) DeleteFollowRecord(c context.Context, commentid string) (res *mongo.DeleteResult, err error) {
	res, err = a.followRepo.DeleteFollowRecord(c, commentid)
	return
}