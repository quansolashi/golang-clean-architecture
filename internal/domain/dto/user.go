package dto

type UserListParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
