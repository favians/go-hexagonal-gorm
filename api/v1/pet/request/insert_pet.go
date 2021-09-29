package request

import "go-hexagonal/business/pet"

//InsertPetRequest create Pet request payload
type InsertPetRequest struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Kind   string `json:"kind"`
}

//ToUpsertPetSpec convert into Pet.UpsertPetSpec object
func (req *InsertPetRequest) ToUpsertPetSpec() *pet.InsertPetSpec {

	var insertPetSpec pet.InsertPetSpec

	insertPetSpec.UserID = req.UserID
	insertPetSpec.Name = req.Name
	insertPetSpec.Kind = req.Kind

	return &insertPetSpec
}
