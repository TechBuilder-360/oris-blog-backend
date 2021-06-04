package repository

import (
	"blog/domain"
	"blog/util"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type BookmarkRepository struct {
	Collection *mongo.Collection
}

// NewBookmarkRepository will create an object that represent the bookmark.Repository interface
func NewBookmarkRepository(Collection *mongo.Collection) domain.BookmarkRepository {
	return &BookmarkRepository{Collection}
}

func (c *BookmarkRepository) UpdateBookmark(ctx context.Context, userId string, postId string, mode string) (status string, err error){
	
	var userIdDoc domain.Bookmark

	//find user struct in db
	var filter bson.M = bson.M{}
	filter = bson.M{"userid": userId}
	errFindUserRecord := c.Collection.FindOne(context.Background(), filter).Decode(&userIdDoc)

	// mode == "add"
	/* check if bookmark doc exists
	if it does, check if postid is in []bookmark, if it is, reject with error="Already in bookmark"
	if it isn't, append the new record to the bookmark array*/

	if mode == "add" {
		if errFindUserRecord != nil {
			// means user has no record, thus we create a new record
			userIdDoc.UserID = userId
			userIdDoc.Bookmarks = []string{postId}
			_, err := c.Collection.InsertOne(context.Background(), userIdDoc)
	
			if err != nil {
				return "FAILED", err
			}
		}else{
			// user has a record, check if the post isn't in the record already
			isAnElement, _ := util.Find(userIdDoc.Bookmarks, postId)
			
			if isAnElement {
				return "FAILED", errors.New("post already in bookmark")
			}

			// post not in bookmark, update bookmark
			userIdDoc.Bookmarks = append(userIdDoc.Bookmarks, postId)
			update := bson.M{
				"$set": userIdDoc,
			}

			_, err := c.Collection.UpdateOne(context.Background(), bson.M{"userid": userId}, update)

			if err != nil {
				return "FAILED", err
			}
		}
	}

	// mode == "remove"
	/* check if bookmark doc exists
	if it does, remove the record from the bookmark array*/

	if mode == "remove" {
		// confirm user has record: errFindUserRecord == nil, means user has record, otherwise, user has no record
		if errFindUserRecord == nil {
			// check that post is in bookmark for sure
			isAnElement, index := util.Find(userIdDoc.Bookmarks, postId)

			if isAnElement {
				userIdDoc.Bookmarks = util.RemoveIndex(userIdDoc.Bookmarks, index)
			}else{
				return "FAILED", errors.New("post not in bookmark")
			}
			
			update := bson.M{
				"$set": userIdDoc,
			}

			_, err := c.Collection.UpdateOne(context.Background(), bson.M{"userid": userId}, update)

			if err != nil {
				return "FAILED", err
			}
		}
	}	
	
	return "SUCCESS", nil
}

func (c *BookmarkRepository) FetchBookmarks(ctx context.Context, userId string) (res []primitive.M, err error) {
	var bookmark []bson.M

	var filter bson.M = bson.M{"userid": userId}

	cur, err := c.Collection.Find(context.Background(), filter)
	
	if err != nil {
		return bookmark, err
	}

	defer cur.Close(context.Background())

	cur.All(context.Background(), &bookmark)

	return bookmark, nil
}

func (c *BookmarkRepository) DeleteBookmarkRecord(ctx context.Context, bookmarkId string) (status string, err error) {
	bookmarkObjID, _ := primitive.ObjectIDFromHex(bookmarkId)

	_, err = c.Collection.DeleteOne(context.Background(), bson.M{"_id": bookmarkObjID})

	if err != nil {
		return "FAILED", err
	}

	return "SUCCESS", nil
}

func (c *BookmarkRepository) ValidateUserRecordExistence(ctx context.Context, userId string) bool {
	var filter bson.M = bson.M{}
	filter = bson.M{"userid": userId}
	var userBookmark domain.Bookmark
 
	err := c.Collection.FindOne(context.Background(), filter).Decode(&userBookmark)

	return err == nil
}
