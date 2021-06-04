package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Bookmark struct
type Bookmark struct {
	UserID string `json:"userid"`
	Bookmarks [] string `json:"bookmarks"`
}

type BookmarkEntity interface {

	UpdateBookmark(ctx context.Context, userId string, postId string, mode string) (status string, err error)
	ValidateUserRecordExistence(ctx context.Context, userId string) bool 

	FetchBookmarks(ctx context.Context, userId string) (res []primitive.M, err error)
	DeleteBookmarkRecord(ctx context.Context, bookmarkId string) (status string, err error)
}

// PostRepository interface
type BookmarkRepository interface {

	UpdateBookmark(ctx context.Context, userId string, postId string, mode string) (status string, err error)
	ValidateUserRecordExistence(ctx context.Context, userId string) bool 

	FetchBookmarks(ctx context.Context, userId string) (res []primitive.M, err error)
	DeleteBookmarkRecord(ctx context.Context, bookmarkId string) (status string, err error)

}