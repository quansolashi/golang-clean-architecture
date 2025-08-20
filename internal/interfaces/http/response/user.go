package response

type User struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}

type UserResponse struct {
	User *User `json:"user"`
}

type UsersResponse struct {
	Users []*User `json:"users"`
}
