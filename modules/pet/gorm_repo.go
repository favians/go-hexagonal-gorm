package pet

import (
	"go-hexagonal/business/pet"
	"go-hexagonal/modules/user"
	"time"

	"gorm.io/gorm"
)

//GormRepository The implementation of pet.Repository object
type GormRepository struct {
	DB *gorm.DB
}

type PetTable struct {
	ID         int            `gorm:"id"`
	UserID     int            `gorm:"user_id"`
	Name       string         `gorm:"name"`
	Kind       string         `gorm:"kind"`
	CreatedAt  time.Time      `gorm:"created_at"`
	CreatedBy  string         `gorm:"created_by"`
	ModifiedAt time.Time      `gorm:"modified_at"`
	ModifiedBy string         `gorm:"modified_by"`
	Version    int            `gorm:"version"`
	User       user.UserTable `gorm:"foreignKey:UserID"`
}

func newPetTable(pet pet.Pet) *PetTable {

	return &PetTable{
		pet.ID,
		pet.UserID,
		pet.Name,
		pet.Kind,
		pet.CreatedAt,
		pet.CreatedBy,
		pet.ModifiedAt,
		pet.ModifiedBy,
		pet.Version,
		user.UserTable{},
	}

}

func (col *PetTable) ToPet() pet.Pet {
	var pet pet.Pet

	pet.ID = col.ID
	pet.UserID = col.UserID
	pet.Name = col.Name
	pet.Kind = col.Kind
	pet.CreatedAt = col.CreatedAt
	pet.CreatedBy = col.CreatedBy
	pet.ModifiedAt = col.ModifiedAt
	pet.ModifiedBy = col.ModifiedBy
	pet.Version = col.Version

	return pet
}

//NewGormDBRepository Generate Gorm DB pet repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

//FindPetByID If data not found will return nil without error
func (repo *GormRepository) FindPetByID(id int, userID int) (*pet.Pet, error) {

	var petData PetTable

	err := repo.DB.Where("id = ?", id).Where("user_id = ?", userID).First(&petData).Error
	if err != nil {
		return nil, err
	}

	pet := petData.ToPet()

	return &pet, nil
}

//FindAllPet find all pet with given specific page and row per page, will return empty slice instead of nil
func (repo *GormRepository) FindAllPet(userID int) ([]pet.Pet, error) {

	var pets []PetTable

	err := repo.DB.Where("user_id = ?", userID).Find(&pets).Error
	if err != nil {
		return nil, err
	}

	var result []pet.Pet
	for _, value := range pets {
		result = append(result, value.ToPet())
	}

	return result, nil
}

//InsertPet Insert new Pet into storage
func (repo *GormRepository) InsertPet(pet pet.Pet) error {

	petData := newPetTable(pet)
	petData.ID = 0

	err := repo.DB.Create(petData).Error
	if err != nil {
		return err
	}

	return nil
}

//UpdateItem Update existing item in database
func (repo *GormRepository) UpdatePet(pet pet.Pet, currentVersion int) error {

	petData := newPetTable(pet)

	err := repo.DB.Model(&petData).Where("user_id = ?", pet.UserID).Updates(PetTable{Name: petData.Name, Version: petData.Version}).Error
	if err != nil {
		return err
	}

	return nil
}
