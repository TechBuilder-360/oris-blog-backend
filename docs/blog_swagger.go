package docs

// swagger:route POST /api/blog/post
// CreatePostResponse creates a post.
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
	// Required body parameter to create a post
	// in:body
	Body CreatePostRequest
}