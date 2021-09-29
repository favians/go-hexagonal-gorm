package pet

//Service outgoing port for pet
type Service interface {
	//FindPetByID If data not found will return nil without error
	FindPetByID(id int) (*Pet, error)

	//FindAllPet find all pet with given specific page and row per page, will return empty slice instead of nil
	FindAllPet() ([]Pet, error)

	//InsertPet Insert new Pet into storage
	InsertPet(insertPetSpec InsertPetSpec, createdBy string) error

	//UpdatePet if data not found will return error
	UpdatePet(id int, name string, modifiedBy string, currentVersion int) error
}

//Repository ingoing port for pet
type Repository interface {
	//FindPetByID If data not found will return nil without error
	FindPetByID(id int) (*Pet, error)

	//FindAllPet find all pet with given specific page and row per page, will return empty slice instead of nil
	FindAllPet() ([]Pet, error)

	//InsertPet Insert new Pet into storage
	InsertPet(pet Pet) error

	//UpdatePet if data not found will return error
	UpdatePet(pet Pet, currentVersion int) error
}
