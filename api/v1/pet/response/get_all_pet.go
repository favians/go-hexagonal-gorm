package response

import (
	"go-hexagonal/api/paginator"
	"go-hexagonal/business/pet"
)

type getAllPetResponse struct {
	Meta paginator.Meta   `json:"meta"`
	Pets []GetPetResponse `json:"pets"`
}

//NewGetAllPetResponse construct GetAllPetResponse
func NewGetAllPetResponse(pets []pet.Pet, page int, rowPerPage int) getAllPetResponse {

	var (
		lenPets = len(pets)
	)

	getAllPetResponse := getAllPetResponse{}
	getAllPetResponse.Meta.BuildMeta(lenPets, page, rowPerPage)

	for i := 0; i < getAllPetResponse.Meta.RowPerPage; i++ {
		var getPetResponse GetPetResponse

		getPetResponse.ID = pets[i].ID
		getPetResponse.UserID = pets[i].UserID
		getPetResponse.Name = pets[i].Name
		getPetResponse.Kind = pets[i].Kind
		getPetResponse.ModifiedAt = pets[i].ModifiedAt
		getPetResponse.Version = pets[i].Version

		getAllPetResponse.Pets = append(getAllPetResponse.Pets, getPetResponse)
	}

	if getAllPetResponse.Pets == nil {
		getAllPetResponse.Pets = []GetPetResponse{}
	}

	return getAllPetResponse
}
