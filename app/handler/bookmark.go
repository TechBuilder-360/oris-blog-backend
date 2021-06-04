package handler

import (
	"blog/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BookmarkHandler ...
type BookmarkHandler struct {
	BookmarkEntity domain.BookmarkEntity
}

// NewBooksHandler ...
func NewBookmarkHandler(r *gin.RouterGroup, ce domain.BookmarkEntity) {
	handler := &BookmarkHandler{
		BookmarkEntity: ce,
	}

	r.GET("/bookmarks/:userId", handler.FindBookmark)

	r.PUT("/bookmarks/:userId/:postId/:operation", handler.UpdateBookmark)

	// r.DELETE("/bookmarks/:userId/:bookmarkId", handler.DeleteBookmark)
}

// FindBookmark ... depending on query parameters
func (a *BookmarkHandler) FindBookmark(c *gin.Context) {

	bookmark, err := a.BookmarkEntity.FetchBookmarks(c.Request.Context(), c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bookmark})
}

// UpdateBookmark ...
func (a *BookmarkHandler) UpdateBookmark(c *gin.Context) {
	mode := c.Param("operation")

	if mode != "add" && mode != "remove" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid operation"})
		return
	}
	
	response, err := a.BookmarkEntity.UpdateBookmark(c.Request.Context(), c.Param("userId"), c.Param("postId"), mode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": response, "message": err.Error()})
		return
	}
	
	if mode == "add"{
		c.JSON(http.StatusOK, gin.H{"status": response, "message": "Post added to bookmark successfully"})
	}else{
		c.JSON(http.StatusOK, gin.H{"status": response, "message": "Post removed from bookmark successfully"})
	}
}

// DeleteBookmark ...
func (a *BookmarkHandler) DeleteBookmark(c *gin.Context) {

	// remove post from book 

	commResponse, err := a.BookmarkEntity.DeleteBookmarkRecord(c.Request.Context(), c.Param("bookmarkid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": commResponse, "message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": commResponse, "message" : "Record Deleted Successfully"})
}
