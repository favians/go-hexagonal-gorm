package user

import (
	"go-hexagonal/business/user"
	"time"

	"gorm.io/gorm"
)

//GormRepository The implementation of user.Repository object
type GormRepository struct {
	DB *gorm.DB
}

type UserTable struct {
	ID         int       `gorm:"id;primaryKey;autoIncrement"`
	Name       string    `gorm:"name"`
	Username   string    `gorm:"username"`
	Password   string    `gorm:"password"`
	CreatedAt  time.Time `gorm:"created_at"`
	CreatedBy  string    `gorm:"created_by"`
	ModifiedAt time.Time `gorm:"modified_at"`
	ModifiedBy string    `gorm:"modified_by"`
	Version    int       `gorm:"version"`
}

func newUserTable(user user.User) *UserTable {

	return &UserTable{
		user.ID,
		user.Name,
		user.Username,
		user.Password,
		user.CreatedAt,
		user.CreatedBy,
		user.ModifiedAt,
		user.ModifiedBy,
		user.Version,
	}

}

func (col *UserTable) ToUser() user.User {
	var user user.User

	user.ID = col.ID
	user.Name = col.Name
	user.Username = col.Username
	user.Password = col.Password
	user.CreatedAt = col.CreatedAt
	user.CreatedBy = col.CreatedBy
	user.ModifiedAt = col.ModifiedAt
	user.ModifiedBy = col.ModifiedBy
	user.Version = col.Version

	return user
}

//NewGormDBRepository Generate Gorm DB user repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

//FindUserByID If data not found will return nil without error
func (repo *GormRepository) FindUserByID(id int) (*user.User, error) {

	var userData UserTable

	err := repo.DB.First(&userData, id).Error
	if err != nil {
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}

//FindUserByID If data not found will return nil without error
func (repo *GormRepository) FindUserByUsernameAndPassword(username string, password string) (*user.User, error) {

	var userData UserTable

	err := repo.DB.Where("username = ?", username).Where("password = ?", password).First(&userData).Error
	if err != nil {
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}

//FindAllUser find all user with given specific page and row per page, will return empty slice instead of nil
func (repo *GormRepository) FindAllUser() ([]user.User, error) {

	var users []UserTable

	err := repo.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	var result []user.User
	for _, value := range users {
		result = append(result, value.ToUser())
	}

	return result, nil
}

//InsertUser Insert new User into storage
func (repo *GormRepository) InsertUser(user user.User) error {

	userData := newUserTable(user)
	userData.ID = 0

	err := repo.DB.Create(userData).Error
	if err != nil {
		return err
	}

	return nil
}

//UpdateItem Update existing item in database
func (repo *GormRepository) UpdateUser(user user.User, currentVersion int) error {

	userData := newUserTable(user)

	err := repo.DB.Model(&userData).Updates(UserTable{Name: userData.Name, Version: userData.Version}).Error
	if err != nil {
		return err
	}

	return nil
}
