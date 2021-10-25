package response

import (
	"go-hexagonal/business/pet"
	"time"
)

//GetPetWithUserResponse Get pet by ID response payload
type GetPetWithUserResponse struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	User       User      `json:"user"`
	Name       string    `json:"name"`
	Kind       string    `json:"kind"`
	ModifiedAt time.Time `json:"modified_at"`
	Version    int       `json:"version"`
}

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

//NewGetPetWithUserResponse construct GetPetWithUserResponse
func NewGetPetWithUserResponse(pet pet.Pet) *GetPetWithUserResponse {
	var getPetWithUserResponse GetPetWithUserResponse

	getPetWithUserResponse.ID = pet.ID
	getPetWithUserResponse.UserID = pet.UserID
	getPetWithUserResponse.User.Name = pet.User.Name
	getPetWithUserResponse.User.Username = pet.User.Username
	getPetWithUserResponse.Name = pet.Name
	getPetWithUserResponse.Kind = pet.Kind
	getPetWithUserResponse.ModifiedAt = pet.ModifiedAt
	getPetWithUserResponse.Version = pet.Version

	return &getPetWithUserResponse
}
