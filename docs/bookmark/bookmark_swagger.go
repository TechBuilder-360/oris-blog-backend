package docs

// r.PUT("/bookmarks/:userId/:postId/:operation", handler.UpdateBookmark)

// ---------------------------------GetUserBookmark-----------------------

// swagger:route GET /bookmarks/{userId} bookmark-CRUD GetBookmarkRequest
// Fetch User Bookmarks.
// responses:
//   200: getBookmarkResponse

// Fetch a user's bookmark record.
// swagger:response getBookmarkResponse
type getBookmarkResponseWrapper struct {
	// in:body
	Bookmark Bookmark
}

// swagger:parameters GetBookmarkRequest
type getBookmarkRequestWrapper struct {
	// User's ID
	// in: path
	// required: true
	UserID string `json:"userId"`
}

// ---------------------------------UpdateUserBookmark-----------------------

// swagger:route PUT /bookmarks/{userId}/{postId}/{operation} bookmark-CRUD UpdateBookmarkRequest
// Update User Bookmarks: Add or remove posts
// responses:
//   200: updateBookmarkResponse

// Update a user's bookmark record.
// swagger:response updateBookmarkResponse
type updateBookmarkResponseWrapper struct {
	// in:body
	Status StatusResponse
}

// swagger:parameters UpdateBookmarkRequest
type updateBookmarkRequestWrapper struct {
	// Users ID
	// in: path
	// required: true
	UserID string `json:"userId"`
	// Post ID
	// in: path
	// required: true
	PostID string `json:"postId"`
	// Operation: 'add' or 'remove'
	// in: path
	// required: true
	Operation string `json:"operation"`
}