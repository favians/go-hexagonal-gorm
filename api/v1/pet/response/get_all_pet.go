package response

import "go-hexagonal/business/pet"

type getAllPetResponse struct {
	Meta meta             `json:"meta"`
	Pets []GetPetResponse `json:"pets"`
}

type meta struct {
	Page         int  `json:"page"`
	RowPerPage   int  `json:"row_per_page"`
	NextPage     bool `json:"next_page"`
	PreviousPage bool `json:"previous_page"`
}

//NewGetAllPetResponse construct GetAllPetResponse
func NewGetAllPetResponse(pets []pet.Pet, page int, rowPerPage int) getAllPetResponse {

	var (
		lenPets = len(pets)
	)

	getAllPetResponse := getAllPetResponse{}
	getAllPetResponse.Meta.Page = page
	getAllPetResponse.Meta.RowPerPage = rowPerPage
	getAllPetResponse.Meta.NextPage = false

	if lenPets > rowPerPage {
		getAllPetResponse.Meta.NextPage = true
	}

	if (lenPets-1 <= rowPerPage) && (page != 1) {
		getAllPetResponse.Meta.PreviousPage = true
	}

	for index, value := range pets {
		if index == rowPerPage {
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
