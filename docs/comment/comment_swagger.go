package docs

// ---------------------------------CreateComment-----------------------

// swagger:route POST /comments comment-CRUD CreateCommentRequest
// Create a comment.
// responses:
//   200: CreatePostResponse

// This text will appear as description of your response body.
// swagger:response CreatePostResponse
type createCommentResponseWrapper struct {
	// in:body
	Status StatusResponse
}

// swagger:parameters CreateCommentRequest
type createCommentRequestWrapper struct {
	// Required body parameter to create a comment
	// in:body
	Body CreateCommentRequest
}

// ---------------------------------GetComments-----------------------

// swagger:route GET /comments comment-CRUD GetCommentRequest
// Fetch Comments.
// responses:
//   200: getCommentsResponse

// Fetch comments response according to supplied query params.
// swagger:response getCommentsResponse
type getCommentsResponseWrapper struct {
	// in:body
	Comments []Comment
}

// swagger:parameters GetCommentRequest
type getCommentsRequestWrapper struct {
	// Fetch Comments by authorid or commentid or both
	// in: query
	// required: true
	AuthorID string `json:"authorid"`
	CommentID string `json:"commentid"`
}

// ---------------------------------UpdateComments-----------------------

// swagger:route PUT /comments/{commentid} comment-CRUD updateCommentRequest
// Update Comment
// responses:
//   200: updateCommentResponse

// Update Comment
// swagger:response updateCommentResponse
type updateCommentsResponseWrapper struct {
	// in:body
	Status StatusResponse
}

// swagger:parameters updateCommentRequest
type updateCommentsRequestWrapper struct {
	// Comment to be updated
	// in: path
	// required: true
	CommentID string `json:"commentid"`
	// Updated data of comment
	// in: body
	// required: true
	Comment Comment
}

// ---------------------------------DeleteComment-----------------------

// swagger:route DELETE /comments/{commentid} comment-CRUD deleteCommentRequest
// Delete Comment
// responses:
//   200: deleteCommentResponse

// Delete Comment
// swagger:response deleteCommentResponse
type deleteCommentsResponseWrapper struct {
	// in:body
	Status StatusResponse
}

// swagger:parameters deleteCommentRequest
type deleteCommentsRequestWrapper struct {
	// Comment to be deleted
	// in: path
	// required: true
	CommentID string `json:"commentid"`
}