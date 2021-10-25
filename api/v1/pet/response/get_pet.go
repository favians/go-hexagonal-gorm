package response

import (
	"go-hexagonal/business/pet"
	"time"
)

//GetPetResponse Get pet by ID response payload
type GetPetResponse struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Name       string    `json:"name"`
	Kind       string    `json:"kind"`
	ModifiedAt time.Time `json:"modified_at"`
	Version    int       `json:"version"`
}

//NewGetPetResponse construct GetPetResponse
func NewGetPetResponse(pet pet.Pet) *GetPetResponse {
	var getPetResponse GetPetResponse

	getPetResponse.ID = pet.ID
	getPetResponse.UserID = pet.UserID
	getPetResponse.Name = pet.Name
	getPetResponse.Kind = pet.Kind
	getPetResponse.ModifiedAt = pet.ModifiedAt
	getPetResponse.Version = pet.Version

	return &getPetResponse
}
