package model

type User struct {
	Id   int
	Name string
}

type CreateUserRequest struct {
	Name string `json:"name"`
}
