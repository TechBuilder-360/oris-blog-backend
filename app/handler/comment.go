package handler

import (
	"blog/domain"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CommentHandler ...
type CommentHandler struct {
	CommentEntity domain.CommentEntity
	PostRepo      domain.PostRepository
}

// NewBooksHandler ...
func NewCommentHandler(r *gin.RouterGroup, ce domain.CommentEntity, pr domain.PostRepository) {
	handler := &CommentHandler{
		CommentEntity: ce,
		PostRepo:      pr,
	}

	//queryparams  (no params returns all Comments)
	r.GET("/comments", handler.FindComment)

	r.POST("/comments", handler.CreateComment)

	r.PUT("/comments/:commentid", handler.UpdateComment)

	r.DELETE("/comments/:postid/:commentid", handler.DeleteComment)
}

// FindComment ... depending on query parameters
func (a *CommentHandler) FindComment(c *gin.Context) {

	comment, _ := a.CommentEntity.FetchComment(c.Request.Context(), c)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// CreateComment ...
func (a *CommentHandler) CreateComment(c *gin.Context) {
	var reqComment domain.Comment

	err := c.ShouldBind(&reqComment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILED", "message": err.Error()})
		return
	}

	// What if not all required details are posted?
	// Comment creation will still be successful, but this does not conform with the app requirement
	// Thus, a validation should be in place for this

	if reqComment.PostID == "" || reqComment.Author == "" || reqComment.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "FAILED", "message": "Invalid request"})
		return
	}


	// reqComment.PostID : check if this is valid before creating comment in docs
	// empty authorid parameter checks for ONLY post existence
	isValid := a.PostRepo.ValidatePostExistence(c.Request.Context(), "", reqComment.PostID)

	if !isValid {
		c.JSON(http.StatusNotAcceptable, gin.H{"status": "FAILED", "message": "invalid post id"})
		return
	}

	commentID, _ := a.CommentEntity.CreateComment(c.Request.Context(), reqComment)

	// update PostID: insert commentID into comment [] then increase commentCount
	response, err := a.PostRepo.InsertPostComment(c.Request.Context(), commentID, reqComment.PostID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILED", "message": err.Error()})
		return
	}

	if response.ModifiedCount == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILED", "message": "unable to match comment to post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "SUCCESS", "message": "comment created"})
}

// UpdateComment ...
func (a *CommentHandler) UpdateComment(c *gin.Context) {
	var Comment domain.Comment
	err := c.ShouldBind(&Comment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILED","message": err})
		return
	}

	_, err = a.CommentEntity.UpdateComment(c.Request.Context(), c.Param("commentid"), Comment)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "FAILED", "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": "SUCCESS", "message" : "record updated successfully"})
}

// DeleteComment ...
func (a *CommentHandler) DeleteComment(c *gin.Context) {

	// remove comment from post too 
	postResponse, err := a.PostRepo.RemovePostComment(c.Request.Context(), c.Param("postid"), c.Param("commentid"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILED", "message": err.Error()})
		return
	}
	
	errMessage := errors.New("unable to delete comment")

	if postResponse.ModifiedCount == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILED", "message": errMessage.Error()})
		return
	}

	_, err = a.CommentEntity.DeleteComment(c.Request.Context(), c.Param("commentid"))

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "FAILED", "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": "SUCCESS", "message" : "record deleted successfully"})
}
