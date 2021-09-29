package user

//Service outgoing port for user
type Service interface {
	//FindUserByID If data not found will return nil without error
	FindUserByID(id int) (*User, error)

	//FindUserByUsernameAndPassword If data not found will return nil
	FindUserByUsernameAndPassword(username string, password string) (*User, error)

	//FindAllUser find all user with given specific page and row per page, will return empty slice instead of nil
	FindAllUser(skip int, rowPerPage int) ([]User, error)

	//InsertUser Insert new User into storage
	InsertUser(insertUserSpec InsertUserSpec, createdBy string) error

	//UpdateUser if data not found will return error
	UpdateUser(id int, name string, modifiedBy string, currentVersion int) error
}

//Repository ingoing port for user
type Repository interface {
	//FindUserByID If data not found will return nil without error
	FindUserByID(id int) (*User, error)

	//FindUserByUsernameAndPassword If data not found will return nil
	FindUserByUsernameAndPassword(username string, password string) (*User, error)

	//FindAllUser find all user with given specific page and row per page, will return empty slice instead of nil
	FindAllUser(skip int, rowPerPage int) ([]User, error)

	//InsertUser Insert new User into storage
	InsertUser(user User) error

	//UpdateUser if data not found will return error
	UpdateUser(user User, currentVersion int) error
}
