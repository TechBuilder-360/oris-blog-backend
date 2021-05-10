package entity

import (
	"blog/domain"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PostEntity ...
type PostEntity struct {
	postRepo domain.PostRepository
}

// NewPostEntity will create new an articleUsecase object representation of domain.PostEntity interface
func NewPostEntity(a domain.PostRepository) domain.PostEntity {
	return &PostEntity{
		postRepo: a,
	}
}

//FetchPost retrives post record(s)...
func (a *PostEntity) FetchPost(c context.Context, ginContext *gin.Context) (res []primitive.M, err error) {
	res, err = a.postRepo.FetchPost(c, ginContext)
	return
}

//CreatePost creates a single post record...
func (a *PostEntity) CreatePost(c context.Context, post domain.Post) (res *mongo.InsertOneResult, err error) {
	res, err = a.postRepo.CreatePost(c, post)
	return
}

//UpdatePost creates a single post record...
func (a *PostEntity) UpdatePost(c context.Context, id string, post domain.Post) (res string, err error) {
	res, err = a.postRepo.UpdatePost(c, id, post)
	return
}

func (a *PostEntity) ValidatePostExistence(c context.Context, authorid string, postid string) (res bool) {
	res = a.postRepo.ValidatePostExistence(c, authorid, postid)
	return
}

func (a *PostEntity) LikePost(c context.Context, userId string, postid string, mode string) (res string, err error) {
	res, err = a.postRepo.LikePost(c, userId, postid, mode)
	return
}

//DeletePost creates a single post record...
func (a *PostEntity) DeletePost(c context.Context, id string) (res string, err error) {
	res, err = a.postRepo.DeletePost(c, id)
	return
}