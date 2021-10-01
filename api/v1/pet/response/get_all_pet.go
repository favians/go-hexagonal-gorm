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

	for index, value := range pets {
		if index == getAllPetResponse.Meta.RowPerPage {
			continue
		}

		var getPetResponse GetPetResponse

		getPetResponse.ID = value.ID
		getPetResponse.UserID = value.UserID
		getPetResponse.Name = value.Name
		getPetResponse.Kind = value.Kind
		getPetResponse.ModifiedAt = value.ModifiedAt
		getPetResponse.Version = value.Version

		getAllPetResponse.Pets = append(getAllPetResponse.Pets, getPetResponse)
	}

	if getAllPetResponse.Pets == nil {
		getAllPetResponse.Pets = []GetPetResponse{}
	}

	return getAllPetResponse
}
