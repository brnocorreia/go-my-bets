package request

import "time"

type UserRequest struct {
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Name      string    `json:"name" binding:"required,min=4,max=100"`
	Nickname  string    `json:"nickname" binding:"required,min=4,max=30"`
	CreatedAt time.Time `json:"created_at"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"omitempty,min=4,max=100"`
	NickName string `json:"nickname" binding:"omitempty,min=4,max=30"`
}
