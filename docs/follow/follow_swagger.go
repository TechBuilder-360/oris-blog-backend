package docs


// ---------------------------------CreateFollowRelationship-----------------------

// swagger:route POST /follow/{userId}/{followedUserId} follow-CRUD CreateFollowRequest
// Follow a user.
// responses:
//   200: CreateFollowResponse

// User follows FollowedUser.
// swagger:response CreateFollowResponse
type createFollowResponseWrapper struct {
	// in:body
	Status StatusResponse
}

// swagger:parameters CreateFollowRequest
type createFollowRequestWrapper struct {
	// user
	// in:path
	// required:true
	UserID string `json:"userId"`
	// user being followed
	// in:path
	// required:true
	FollowedUserId string `json:"followedUserId"`
}

// ---------------------------------DestroyFollowRelationship-----------------------

// swagger:route PUT /unfollow/{userId}/{unfollowedUserId} follow-CRUD CreateUnFollowRequest
// Unfollow a user.
// responses:
//   200: CreateUnFollowResponse

// User unfollows FollowedUser.
// swagger:response CreateUnFollowResponse
type unFollowResponseWrapper struct {
	// in:body
	Status StatusResponse
}

// swagger:parameters CreateUnFollowRequest
type unFollowRequestWrapper struct {
	// user
	// in:path
	// required:true
	UserID string `json:"userId"`
	// user to unfollow
	// in:path
	// required:true
	UnfollowedUserId string `json:"unfollowedUserId"`
}

// ---------------------------------FetchFollowRelationships-----------------------

// swagger:route GET /follows/ follow-CRUD getUnFollowsRequest
// Unfollow a user.
// responses:
//   200: getFollowsResponse

// Fetch follows
// swagger:response getFollowsResponse
type getFollowResponseWrapper struct {
	// in:body
	Follow []Follow
}

// swagger:parameters getUnFollowsRequest
type getFollowRequestWrapper struct {
	// Get a user follow relationships
	// in:query
	UserID string `json:"userId"`
}

// ---------------------------------DeleteFollowRecord-----------------------

// swagger:route DELETE /follows/{followid} follow-CRUD deleteFollowRequest
// Delete Folow
// responses:
//   200: deleteFollowResponse

// Delete Follow
// swagger:response deleteFollowResponse
type deleteFollowsResponseWrapper struct {
	// in:body
	Status StatusResponse
}

// swagger:parameters deleteFollowRequest
type deleteFollowsRequestWrapper struct {
	// follow record to be deleted
	// in: path
	// required: true
	Followid string `json:"followid"`
}