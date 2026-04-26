package user

// UpdateUserRequest is the payload for PUT /users/:id.
type UpdateUserRequest struct {
	Name string `json:"name"`
}

// UserResponse is returned in user endpoints.
type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}

// ListUsersRequest holds query params for listing users.
type ListUsersRequest struct {
	Page    int `query:"page"`
	PerPage int `query:"per_page"`
}
