package transport

import "time"

type User struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserName  string    `json:"userName"`
	FullName  string    `json:"fullName"`
	Password  string    `json:"password"`
}
