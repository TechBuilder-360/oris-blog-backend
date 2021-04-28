package handler

import (
	"blog/domain"
	"fmt"
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

	r.POST("/follow/:userId/:followedUserId", handler.Follow)

	r.PUT("/unfollow/:userId/:unfollowedUserId", handler.UnFollow)

	// query param userId
	r.GET("/follows", handler.FetchFollow)

	r.DELETE("/follows/:followid", handler.DeleteFollowRecord)
}

// FindFollow ... all or by userID
func (a *FollowHandler) FetchFollow(c *gin.Context) {

	follows, _ := a.FollowEntity.FetchFollows(c.Request.Context(), c)

	c.JSON(http.StatusOK, gin.H{"data": follows})
}

// CreateFollow ...
func (a *FollowHandler) Follow(c *gin.Context) {
	var follower domain.Follow
	var followed domain.Follow
	
	userId := c.Param("userId")
	followedId := c.Param("followedUserId")

	// check if both IDs are same

	if userId == followedId {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Invalid request"})
		return
	}
	
	// check if both IDs are valid userIDs

	// if !isValid {
	// 	c.JSON(http.StatusNotAcceptable, gin.H{"error": "Unable to create follow"})
	// 	return
	// }
	
	// check if follower is already following user : followerFollowingUser
	// check if user is following follower : userFollowingFollower
	userFollowingFollowed, followedFollowingUser := a.FollowRepo.ValidateRelationshipExistence(c.Request.Context(), userId, followedId)

	if followedFollowingUser && userFollowingFollowed {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Relationship exists already"})
		return
	}

	// actual request is for user (who happens to be desired follower) to follow the second user 'followedUser'
	// this request will pass if an only if user is NOT already following 'followerUser' i.e userFollowingFollower is false

	if userFollowingFollowed {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Relationship exists already"})
		return
	}

	// user can follow the desired user

	// check if BOTH have records first.
	userHasRecord := a.FollowRepo.ValidateUserRecordExistence(c.Request.Context(), userId)
	followedHasRecord := a.FollowRepo.ValidateUserRecordExistence(c.Request.Context(), followedId)

	
	if followedHasRecord || userHasRecord {
		
		// we check if userHasRecord has a record
		if userHasRecord {
			// if he does, we update the record by appending the followed user to user.following list
			_, err := a.FollowEntity.UpdateFollowing(c.Request.Context(), userId, followedId, "follow")
	
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
		}else{
			// if he doesn't we create a record for user, then assign followed user to user.following list
			follower.UserID = userId
			follower.Following = []string{followedId}

			response, err := a.FollowEntity.Follow(c.Request.Context(), follower, "one")
			fmt.Println(response)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
		}

		// we check if followedUser has a record
		if followedHasRecord {
			// if he does we update the user.Followers list with user
			_, err := a.FollowEntity.UpdateFollowers(c.Request.Context(), followedId, userId, "follow")
	
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
		}else{
			// if he doesnt, we create a struct for him and assign user to user.followers list
			followed.UserID = userId
			followed.Followers = []string{userId}

			_, err := a.FollowEntity.Follow(c.Request.Context(), follower, "one")

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}

		}

		c.JSON(http.StatusCreated, gin.H{"status": "SUCCESS"})
		return

	}
	// if they dont. default course runs through
	
	// creates struct for the follower
	follower.UserID = userId
	follower.Following = []string{followedId}

	// to create followed struct too, set mode as "many"
	response, err := a.FollowEntity.Follow(c.Request.Context(), follower, "many")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": response})
}

func (a *FollowHandler) UnFollow(c *gin.Context){

	/*
	work flow.
	get the userids.
	userId - user who is unfollowing another user
	unfollowedId - user who is being unfollowed

	unfollow defination - a process where by a userId is removed from unfollowedId.followers list

	flow: 

	check if userId follows unfollowed exist, else reject request
	if above is true, remove userId from unfollowedId.followers list
	then remove unfollowedId from userID.following list
	*/
	
	userId := c.Param("userId")
	followedId := c.Param("unfollowedUserId")

	// check if both IDs are same

	if userId == followedId {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Invalid request"})
		return
	}
	
	// check if both IDs are valid userIDs

	// if !isValid {
	// 	c.JSON(http.StatusNotAcceptable, gin.H{"error": "Unable to create follow"})
	// 	return
	// }
	

	// check if BOTH have record
	userHasRecord := a.FollowRepo.ValidateUserRecordExistence(c.Request.Context(), userId)
	followedHasRecord := a.FollowRepo.ValidateUserRecordExistence(c.Request.Context(), followedId)

	if userHasRecord && followedHasRecord {
		// check if follower is already following user : followerFollowingUser
		// check if user is following follower : userFollowingFollower
		userFollowingFollowed, _ := a.FollowRepo.ValidateRelationshipExistence(c.Request.Context(), userId, followedId)

		// user can unfollow the desired user
		// remove user from unfollowed.folowers list
		if !userFollowingFollowed {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Invalid request"})
			return
		}
		_, err := a.FollowEntity.UpdateFollowers(c.Request.Context(), followedId, userId, "unfollow")
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		// remove unfollowed from user.following list
		_, err = a.FollowEntity.UpdateFollowing(c.Request.Context(), userId, followedId, "unfollow")
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}else{
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "SUCCESS"})
}

func (a *FollowHandler) DeleteFollowRecord(c *gin.Context) {

	commResponse, _ := a.FollowEntity.DeleteFollowRecord(c.Request.Context(), c.Param("followid"))
	if commResponse.DeletedCount == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "FAILED"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "SUCCESS"})
}