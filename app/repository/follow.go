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

func (c *FollowRepository) CreateFollow(ctx context.Context, follower domain.Follow) (resFollow *mongo.InsertManyResult, err error) {
	// create struct for followed
	var followed domain.Follow
	followed.UserID = follower.Following[0]
	followed.Followers = []string{follower.UserID}

	documents := []interface{}{follower, followed}
	response, err := c.Collection.InsertMany(context.Background(), documents)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *FollowRepository) UpdateFollowers(ctx context.Context, userId string, followerId string) (resFollow *mongo.UpdateResult, err error) {
	//find follower struct in db
	var userIdDoc domain.Follow
	var filter bson.M = bson.M{}
	filter = bson.M{"userid": userId}
	cur, _ := c.Collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	cur.All(context.Background(), &userIdDoc)

	userIdDoc.Followers = append(userIdDoc.Followers, followerId)
	
	update := bson.M{
		"$set": userIdDoc,
	}

	response, err := c.Collection.UpdateOne(context.Background(), bson.M{"userid": userId}, update)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *FollowRepository) UpdateFollowing(ctx context.Context, userId string, followedId string) (resFollow *mongo.UpdateResult, err error) {

	//find follower struct in db
	var userIdDoc domain.Follow
	var filter bson.M = bson.M{}
	filter = bson.M{"userid": userId}
	cur, _ := c.Collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	cur.All(context.Background(), &userIdDoc)

	userIdDoc.Following = append(userIdDoc.Following, followedId)
	
	update := bson.M{
		"$set": userIdDoc,
	}

	response, err := c.Collection.UpdateOne(context.Background(), bson.M{"userid": userId}, update)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *FollowRepository) ValidateRelationshipExistence(ctx context.Context, userId string, followerId string) bool {
	
	var filter bson.M = bson.M{}
	filter = bson.M{"userid": userId}
	var followStruct domain.Follow
 
	err := c.Collection.FindOne(context.Background(), filter).Decode(&followStruct)

	if err != nil {
		return false
	}
	
	if util.Find(followStruct.Following, followerId) {
		return true
	}

	if util.Find(followStruct.Followers, followerId) {
		return true
	}

	return false
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