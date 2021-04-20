package handler

import (
	"blog/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CommentHandler ...
type FollowHandler struct {
	FollowEntity domain.FollowEntity
	FollowRepo      domain.FollowRepository
}

// NewBooksHandler ...
func NewFollowHandler(r *gin.RouterGroup, fe domain.FollowEntity, fr domain.FollowRepository) {
	handler := &FollowHandler{
		FollowEntity: fe,
		FollowRepo: fr,
	}

	r.POST("/follow/:followerId/:followedId", handler.Follow)

	// r.PUT("/unfollow/:followerId/:followedId", handler.Unfollow)

	// query param userId
	r.GET("/follows", handler.FetchFollow)

	r.DELETE("/follows/:followid", handler.DeleteFollowRecord)
}

// FindFollow ... all or by userID
func (a *FollowHandler) FetchFollow(c *gin.Context) {

	follows, _ := a.FollowEntity.FetchFollows(c.Request.Context(), c)

	c.JSON(http.StatusOK, gin.H{"data": follows})
}

// CreateComment ...
func (a *FollowHandler) Follow(c *gin.Context) {
	var follower domain.Follow
	
	followerId := c.Param("followerId")
	followedId := c.Param("followedId")

	// check if both IDs are same

	if followerId == followedId {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Invalid request"})
		return
	}
	
	// check if both IDs are valid userIDs

	// if !isValid {
	// 	c.JSON(http.StatusNotAcceptable, gin.H{"error": "Unable to create follow"})
	// 	return
	// }

	haveRelationship := a.FollowRepo.ValidateRelationshipExistence(c.Request.Context(), followerId, followedId)

	if haveRelationship {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Relationship exists already"})
		return
	}

	// find out if users have existing record in the document.

	followerHasRecord := a.FollowRepo.ValidateUserRecordExistence(c.Request.Context(), followerId)
	followedHasRecord := a.FollowRepo.ValidateUserRecordExistence(c.Request.Context(), followedId)

	if followerHasRecord {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "follower has record already, you need to update it not create new one"})
		// update follower record.
		response, err := a.FollowEntity.UpdateFollowing(c.Request.Context(), followerId, followedId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	
		c.JSON(http.StatusCreated, gin.H{"data": response})
		return
	}
	
	if followedHasRecord {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "followed has record already, you need to update it not create new one"})
		// update followed record.
		response, err := a.FollowEntity.UpdateFollowers(c.Request.Context(), followedId, followerId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	
		c.JSON(http.StatusCreated, gin.H{"data": response})
		return
	}

	if 1==1 {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Done bro"})
		return
	}
	
	// if it does update its document : add follwedId as an element of its following
	
	// otherwise, create its new struct
	// creates struct for the follower
	follower.UserID = followerId
	follower.Following = []string{followedId}

	response, err := a.FollowEntity.Follow(c.Request.Context(), follower)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": response})
}

func (a *FollowHandler) DeleteFollowRecord(c *gin.Context) {

	commResponse, _ := a.FollowEntity.DeleteFollowRecord(c.Request.Context(), c.Param("followid"))
	c.JSON(http.StatusOK, gin.H{"data": commResponse})
}