package transport

import "time"

// User struct represents the user entity and it used for json un/marshalling.
type User struct {
	ID        uint64    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	UserName  string    `json:"userName,omitempty"`
	FullName  string    `json:"fullName,omitempty"`
	Password  string    `json:"password,omitempty"`
}

// GenericResponse represents a common response of APIs.
type GenericResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
