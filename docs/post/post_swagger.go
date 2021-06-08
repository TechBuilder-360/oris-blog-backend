package docs

// ---------------------------------GetAllPostsByQueryParam-----------------------

// swagger:route GET /posts post-CRUD getPostsByQueryParam
// Endpoint to fetch posts.
// responses:
//   200: postsResponse

// consumes:
//   - application/json
//
// produces:
//   - application/json
//
// schemes: https,
//

// Returns all posts based on supplied queries
// swagger:response postsResponse
type getAllPostsResponse struct {
	// Posts Data
	// in:body
	Posts []GetPostResponse
}

// swagger:parameters getPostsByQueryParam
type getAllPostsRequest struct {
	// Query param to filter post response.
	// No param returns all posts!
	// in: query
	// required: false
	Category string `json:"category"`
	AuthorID string `json:"authorid"`
	PostID string `json:"postid"`
	Status string `json:"status"`
	Slug string `json:"slug"`
}

// ---------------------------------CreatePost-----------------------

// swagger:route POST /posts/{author_id} post-CRUD CreatePostRequest
// Create a post.
// responses:
//   200: CreatePostResponse

// This text will appear as description of your response body.
// swagger:response CreatePostResponse
type createPostResponseWrapper struct {
	// in:body
	Body CreatePostResponse
}

// swagger:parameters CreatePostRequest
type createPostRequestWrapper struct {
	// Author ID
	// in: path
	// required: true
	Author_id string `json:"author_id"`
	// Required body parameter to create a post
	// in:body
	Body CreatePostRequest
}

// ---------------------------------UpdatePost-----------------------

// swagger:route PUT /posts/{author_id} post-CRUD UpdatePostRequest
// Update a post data.
// responses:
//   200: UpdatePostResponse

// This text will appear as description of your response body.
// swagger:response UpdatePostResponse
type updatePostResponse struct {
	// in:body
	Status StatusReponse
}

// swagger:parameters UpdatePostRequest
type updatePostRequest struct {
	// Author ID
	// in: path
	// required: true
	Author_id string `json:"author_id"`
	// Required body parameter to create a post
	// in:body
	Body GetPostResponse
}


// ---------------------------------DeletePost-----------------------

// swagger:route DELETE /posts/{author_id}/{post_id} post-CRUD DeletePostRequest
// Delete a post data.
// responses:
//   200: DeletePostResponse

// This text will appear as description of your response body.
// swagger:response DeletePostResponse
type deletePostResponse struct {
	// in:body
	Status StatusReponse
}

// swagger:parameters DeletePostRequest
type deletePostRequest struct {
	// Author ID
	// in: path
	// required: true
	Author_id string `json:"author_id"`
	// Post ID
	// in: path
	// required: true
	Post_id string `json:"post_id"`
}


// ---------------------------------LikePost-----------------------

// swagger:route PUT /like/{user_id}/{post_id}/{operation} post-CRUD LikePostRequest
// Like a post data.
// responses:
//   200: LikePostResponse

// This text will appear as description of your response body.
// swagger:response LikePostResponse
type LikePostResponse struct {
	// in:body
	Status StatusReponse
}

// swagger:parameters LikePostRequest
type LikePostRequest struct {
	// User ID
	// in: path
	// required: true
	User_id string `json:"user_id"`
	// Post ID
	// in: path
	// required: true
	Post_id string `json:"post_id"`
	// operation value is either 'add' or 'remove'
	// add: to like a post
	// remove: to unlike a post
	// in: path
	// required: true
	Operation string `json:"operation"`
}