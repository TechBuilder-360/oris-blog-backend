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

// NewPostHandler ...
func NewPostHandler(r *gin.RouterGroup, pe domain.PostEntity) {
	handler := &PostHandler{
		PostEntity: pe,
	}

	//queryparams postid, authorid, category, status, slug (no params returns all posts)
	r.GET("/posts", handler.FindPost)

	r.POST("/posts/:authorid", handler.CreatePost)

	r.PUT("/posts/:postid", handler.UpdatePost)

	r.DELETE("/posts/:authorid/:postid", handler.DeletePost)

	r.PUT("/like/:userid/:postid/:operation", handler.LikePost)
}

// FindPost ... depending on query parameters
func (a *PostHandler) FindPost(c *gin.Context) {

	post, _ := a.PostEntity.FetchPost(c.Request.Context(), c)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// CreatePost ...
func (a *PostHandler) CreatePost(c *gin.Context) {
	var reqPost domain.Post 

	err := c.ShouldBind(&reqPost)

	if err != nil {
		c.JSON(http.StatusInternalServerError,  gin.H{"message": err.Error()})
		return
	}

	// check that required fields are available

	if reqPost.Article != "" && reqPost.Status != "" && reqPost.Title != "" && reqPost.AuthorID != "" && reqPost.Summary != "" {
		response, _ := a.PostEntity.CreatePost(c.Request.Context(), reqPost)
		c.JSON(http.StatusCreated, gin.H{"data": response.InsertedID})
	}else{
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "invalid request"})
		return
	}

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
	c.JSON(http.StatusCreated, gin.H{"status": response, "message": "Post updated successfully"})
}

// DeletePost ...
func (a *PostHandler) DeletePost(c *gin.Context) {
	// confirm post belongs to author else reject
	canDelete := a.PostEntity.ValidatePostExistence(c.Request.Context(), c.Param("authorid"), c.Param("postid"))
	// if it doesn't
	if !canDelete {
		c.JSON(http.StatusInternalServerError,  gin.H{"message": "Invalid Request"})
		return
	}

	response, _ := a.PostEntity.DeletePost(c.Request.Context(), c.Param("postid"))
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (a *PostHandler) LikePost(c *gin.Context) {
	mode := c.Param("operation")

	if mode != "add" && mode != "remove" {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "invalid operation"})
		return
	}

	response, err := a.PostEntity.LikePost(c.Request.Context(), c.Param("userid"), c.Param("postid"), c.Param("operation"))
	if err != nil {
		c.JSON(http.StatusInternalServerError,  gin.H{"message": err.Error()})
		return
	}

	if mode == "add"{
		c.JSON(http.StatusOK, gin.H{"status": response, "message": "Post liked successfully"})
	}else{
		c.JSON(http.StatusOK, gin.H{"status": response, "message": "Post unliked successfully"})
	}
}
