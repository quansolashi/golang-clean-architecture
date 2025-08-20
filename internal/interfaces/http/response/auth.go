package response

type Auth struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type AuthResponse struct {
	Auth *Auth `json:"auth"`
}
