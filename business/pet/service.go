package pet

import (
	"go-hexagonal/business"
	"go-hexagonal/business/user"
	"go-hexagonal/util/validator"
	"time"
)

//InsertPetSpec create pet spec
type InsertPetSpec struct {
	UserID int    `validate:"required"`
	Name   string `validate:"required"`
	Kind   string `validate:"required"`
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository  Repository
	userService user.Service
}

//NewService Construct pet service object
func NewService(repository Repository, userService user.Service) Service {
	return &service{
		repository,
		userService,
	}
}

//FindPetByID Get pet by given ID, return nil if not exist
func (s *service) FindPetByID(id int) (*Pet, error) {
	return s.repository.FindPetByID(id)
}

//FindAllPet Get all pets , will be return empty array if no data or error occured
func (s *service) FindAllPet(userID int, skip int, rowPerPage int) ([]Pet, error) {

	pet, err := s.repository.FindAllPet(userID, skip, rowPerPage)
	if err != nil {
		return []Pet{}, err
	}

	return pet, err
}

//InsertPet Create new pet and store into database
func (s *service) InsertPet(insertPetSpec InsertPetSpec, createdBy string) error {
	err := validator.GetValidator().Struct(insertPetSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	pet := NewPet(
		0,
		insertPetSpec.UserID,
		insertPetSpec.Name,
		"",
		"",
		insertPetSpec.Kind,
		createdBy,
		time.Now(),
	)

	err = s.repository.InsertPet(pet)
	if err != nil {
		return err
	}

	return nil
}

//UpdatePet will update found pet, if not exists will be return error
func (s *service) UpdatePet(id int, name string, modifiedBy string, currentVersion int) error {

	pet, err := s.repository.FindPetByID(id)
	if err != nil {
		return err
	} else if pet == nil {
		return business.ErrNotFound
	} else if pet.Version != currentVersion {
		return business.ErrHasBeenModified
	}

	modifiedPet := pet.ModifyPet(name, time.Now(), modifiedBy)

	return s.repository.UpdatePet(modifiedPet, currentVersion)
}

//FindPetByIDWithUserDataJoinInAPP Get pet data with user data inside, join in app service
func (s *service) FindPetByIDWithUserDataJoinInAPP(id int) (*Pet, error) {

	pet, err := s.repository.FindPetByID(id)
	if err != nil {
		return nil, business.ErrNotFound
	}

	userData, err := s.userService.FindUserByID(pet.UserID)
	if err != nil {
		return nil, business.ErrNotFound
	}

	petJoinWithUser := NewPet(
		pet.ID,
		pet.UserID,
		pet.Name,
		userData.Name,
		userData.Username,
		pet.Kind,
		pet.CreatedBy,
		pet.CreatedAt,
	)

	return &petJoinWithUser, nil
}

//FindPetByIDWithUserDataJoinInAPP Get pet data with user data inside, join in app service
func (s *service) FindPetByIDWithUserDataJoinInDB(id int) (*Pet, error) {

	pet, err := s.repository.FindPetByID(id)
	if err != nil {
		return nil, business.ErrNotFound
	}

	return s.repository.FindPetByIDWithUserDataJoinInDB(pet.ID, pet.UserID)
}
