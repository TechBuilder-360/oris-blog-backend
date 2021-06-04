package entity

import (
	"blog/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookmarkEntity ...
type BookmarkEntity struct {
	bookmarkRepo domain.BookmarkRepository
}

// NewBookmarkEntity will create new an articleUsecase object representation of domain.BookmarkEntity interface
func NewBookmarkEntity(a domain.BookmarkRepository) domain.BookmarkEntity {
	return &BookmarkEntity{
		bookmarkRepo: a,
	}
}


func (a *BookmarkEntity) FetchBookmarks(ctx context.Context, userId string) (response []primitive.M, err error){
	response, err = a.bookmarkRepo.FetchBookmarks(ctx, userId)
	return 
}


func (a *BookmarkEntity) UpdateBookmark(ctx context.Context, userId string, postId string, mode string) (status string, err error){
	status, err = a.bookmarkRepo.UpdateBookmark(ctx, userId, postId, mode)
	return 
}

func (a *BookmarkEntity) ValidateUserRecordExistence(ctx context.Context, userId string) (userHasRecord bool) {
	userHasRecord = a.bookmarkRepo.ValidateUserRecordExistence(ctx, userId)
	return 
}


func (a *BookmarkEntity) DeleteBookmarkRecord(c context.Context, commentid string) (status string, err error) {
	status, err = a.bookmarkRepo.DeleteBookmarkRecord(c, commentid)
	return
}