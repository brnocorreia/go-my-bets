package response

import "time"

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}
