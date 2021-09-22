package migration

import (
	"go-hexagonal/modules/pet"
	"go-hexagonal/modules/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.UserTable{}, &pet.PetTable{})
}
