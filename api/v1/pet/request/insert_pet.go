package request

import "go-hexagonal/business/pet"

//InsertPetRequest create Pet request payload
type InsertPetRequest struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

//ToUpsertPetSpec convert into Pet.UpsertPetSpec object
func (req *InsertPetRequest) ToUpsertPetSpec(userID int) *pet.InsertPetSpec {

	var insertPetSpec pet.InsertPetSpec

	insertPetSpec.UserID = userID
	insertPetSpec.Name = req.Name
	insertPetSpec.Kind = req.Kind

	return &insertPetSpec
}
