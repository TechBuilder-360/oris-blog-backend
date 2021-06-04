package entity

import (
	"blog/domain"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CommentEntity ...
type CommentEntity struct {
	commentRepo domain.CommentRepository
}

// NewCommentEntity will create new an articleUsecase object representation of domain.CommentEntity interface
func NewCommentEntity(a domain.CommentRepository) domain.CommentEntity {
	return &CommentEntity{
		commentRepo: a,
	}
}

//FetchComment retrives Comment record(s)...
func (a *CommentEntity) FetchComment(c context.Context, ginContext *gin.Context) (res []primitive.M, err error) {
	res, err = a.commentRepo.FetchComment(c, ginContext)
	return
}

//CreateComment creates a single Comment record...
func (a *CommentEntity) CreateComment(c context.Context, Comment domain.Comment) (res *mongo.InsertOneResult, err error) {
	res, err = a.commentRepo.CreateComment(c, Comment)
	return
}

//UpdateComment creates a single Comment record...
func (a *CommentEntity) UpdateComment(c context.Context, id string, Comment domain.Comment) (res string, err error) {
	res, err = a.commentRepo.UpdateComment(c, id, Comment)
	return
}

//DeleteComment creates a single Comment record...
func (a *CommentEntity) DeleteComment(c context.Context, commentid string) (res string, err error) {
	res, err = a.commentRepo.DeleteComment(c, commentid)
	return
}