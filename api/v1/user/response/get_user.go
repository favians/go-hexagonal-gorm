package response

import (
	"go-hexagonal/business/user"
	"time"
)

//GetUserResponse Get user by ID response payload
type GetUserResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Version    int       `json:"version"`
}

//NewGetUserResponse construct GetUserResponse
func NewGetUserResponse(user user.User) *GetUserResponse {
	var getUserResponse GetUserResponse

	getUserResponse.ID = user.ID
	getUserResponse.Name = user.Name
	getUserResponse.Username = user.Username
	getUserResponse.ModifiedAt = user.ModifiedAt
	getUserResponse.Version = user.Version

	return &getUserResponse
}
