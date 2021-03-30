package handler

import (
	"blog/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostHandler ...
type PostHandler struct {
	PostEntity domain.PostEntity
}

// NewBooksHandler ...
func NewBooksHandler(r *gin.RouterGroup, pe domain.PostEntity) {
	handler := &PostHandler{
		PostEntity: pe,
	}

	//queryparams postid, authorid, category, status (no params returns all posts)
	r.GET("/post", handler.FindPost)

	r.POST("/post", handler.MakePost)

	r.PUT("/post/:postid", handler.UpdatePost)

	r.DELETE("/post/:postid", handler.DeletePost)
}

// FindPost ... depending on query parameters
func (a *PostHandler) FindPost(c *gin.Context) {

	post, _ := a.PostEntity.FetchPost(c.Request.Context(), c)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// MakePost ...
func (a *PostHandler) MakePost(c *gin.Context) {
	var reqPost domain.Post 

	err := c.ShouldBind(&reqPost)

	// What if not all required details are posted?
	// Post will still be successful, but this does not conform with the app requirement
	// Thus, a validation should be in place for this 

	if err != nil {
		c.JSON(http.StatusInternalServerError,  gin.H{"error": err})
		return
	}
	postID, _ := a.PostEntity.CreatePost(c.Request.Context(), reqPost)
	c.JSON(http.StatusCreated, gin.H{"data": postID})
}

// UpdatePost ...
func (a *PostHandler) UpdatePost(c *gin.Context) {
	var post domain.Post
	err := c.ShouldBind(&post)

	if err != nil {
		c.JSON(http.StatusInternalServerError,  gin.H{"error": err})
		return
	}

	response, _ := a.PostEntity.UpdatePost(c.Request.Context(), c.Param("postid"), post)
	c.JSON(http.StatusCreated, gin.H{"data": response})
}

// DeletePost ...
func (a *PostHandler) DeletePost(c *gin.Context) {

	response, _ := a.PostEntity.DeletePost(c.Request.Context(), c.Param("postid"))
	c.JSON(http.StatusOK, gin.H{"data": response})
}
