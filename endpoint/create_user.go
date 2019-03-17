package endpoint

// CreateUserRequest for holding create user request payload
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
