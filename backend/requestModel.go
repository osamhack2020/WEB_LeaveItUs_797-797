package main

type (
	// TagDeleteRequest is definition of DELETE tag request's body
	TagDeleteRequest struct {
		UIDs []string `json:"uids" validate:"required,printascii"`
	}

	// PersonDeleteRequest is definition of DELETE tag request's body
	PersonDeleteRequest struct {
		IDs []string `json:"ids" validate:"required,printascii"`
	}
)
