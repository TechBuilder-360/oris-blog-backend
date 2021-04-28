package repository

import (
	"blog/domain"
	"blog/util"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// FollowRepository ...
type FollowRepository struct {
	Collection *mongo.Collection
}

// NewFollowRepository will create an object that represent the article.Repository interface
func NewFollowRepository(Collection *mongo.Collection) domain.FollowRepository {
	return &FollowRepository{Collection}
}

func (c *FollowRepository) FetchFollows(ctx context.Context, ginContext *gin.Context) (res []primitive.M, err error) {
	var follow []bson.M

	var filter bson.M = bson.M{}

	// get by userId
	if ginContext.Query("userId") != "" {
		queryUserId := ginContext.Query("userId")
		filter = bson.M{"userid": queryUserId}
	}

	cur, err := c.Collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		return follow, err
	}

	cur.All(context.Background(), &follow)

	return follow, nil
}

func (c *FollowRepository) CreateFollow(ctx context.Context, follower domain.Follow, mode string) (status string, err error) {
	
	if mode == "one"{
		_, err := c.Collection.InsertOne(context.Background(), follower)

		if err != nil {
			return "FAILED", err
		} 
	}

	if mode == "many"{
		// create struct for followed
		var followed domain.Follow
		followed.UserID = follower.Following[0]
		followed.Followers = []string{follower.UserID}

		documents := []interface{}{follower, followed}
		_, err := c.Collection.InsertMany(context.Background(), documents)

		if err != nil {
			return "FAILED", err
		}
	}
	
	return "SUCCESS", nil
}

func (c *FollowRepository) UpdateFollowers(ctx context.Context, userId string, followerId string, mode string) (resFollow *mongo.UpdateResult, err error) {
	//find follower struct in db
	var userIdDoc domain.Follow
	var filter bson.M = bson.M{}
	filter = bson.M{"userid": userId}
	err = c.Collection.FindOne(context.Background(), filter).Decode(&userIdDoc)

	if err != nil {
		return nil, err
	}

	if mode == "follow" {
		userIdDoc.Followers = append(userIdDoc.Followers, followerId)
	}else if mode == "unfollow"{
		// remove followerId from userIdDoc.Followers
		isAnElement, index := util.Find(userIdDoc.Followers, followerId)
		if isAnElement {
			userIdDoc.Followers = util.RemoveIndex(userIdDoc.Followers, index)
		}
	}
	
	update := bson.M{
		"$set": userIdDoc,
	}

	response, err := c.Collection.UpdateOne(context.Background(), bson.M{"userid": userId}, update)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *FollowRepository) UpdateFollowing(ctx context.Context, userId string, followedId string, mode string) (resFollow *mongo.UpdateResult, err error) {

	//find follower struct in db
	var userIdDoc domain.Follow
	var filter bson.M = bson.M{}
	filter = bson.M{"userid": userId}
	err = c.Collection.FindOne(context.Background(), filter).Decode(&userIdDoc)

	if err != nil {
		return nil, err
	}

	if mode == "follow"{
		userIdDoc.Following = append(userIdDoc.Following, followedId)
	}else if mode == "unfollow"{
		// remove follwedId from userIdDoc.following
		isAnElement, index := util.Find(userIdDoc.Following, followedId)
		if isAnElement {
			userIdDoc.Following = util.RemoveIndex(userIdDoc.Following, index)
		}
	}
	
	update := bson.M{
		"$set": userIdDoc,
	}

	response, err := c.Collection.UpdateOne(context.Background(), bson.M{"userid": userId}, update)

	if err != nil {
		return response, err
	}

	return response, nil
}


func (c *FollowRepository) ValidateRelationshipExistence(ctx context.Context, userId string, followerId string) (isFollowing bool, isFollowed bool)  {
	
	var filter bson.M = bson.M{}
	filter = bson.M{"userid": userId}
	var followStruct domain.Follow
 
	err := c.Collection.FindOne(context.Background(), filter).Decode(&followStruct)

	if err != nil {
		return false, false
	}
	
	// check if follower is already following user
	isFollowing, _ = util.Find(followStruct.Following, followerId)

	// check if user is following follower
	isFollowed, _ = util.Find(followStruct.Followers, followerId)

	return isFollowing, isFollowed

}

func (c *FollowRepository) ValidateUserRecordExistence(ctx context.Context, userId string) bool {
	var filter bson.M = bson.M{}
	filter = bson.M{"userid": userId}
	var followStruct domain.Follow
 
	err := c.Collection.FindOne(context.Background(), filter).Decode(&followStruct)

	if err != nil {
		return false
	}

	return true
}

func (c *FollowRepository) DeleteFollowRecord(ctx context.Context, followid string) (resComment *mongo.DeleteResult, err error) {
	followObjID, _ := primitive.ObjectIDFromHex(followid)

	response, err := c.Collection.DeleteOne(context.Background(), bson.M{"_id": followObjID})

	if err != nil {
		return response, err
	}

	return response, nil
}