package request

import "go-hexagonal/business/user"

//InsertUserRequest create User request payload
type InsertUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//ToUpsertUserSpec convert into User.UpsertUserSpec object
func (req *InsertUserRequest) ToUpsertUserSpec() *user.InsertUserSpec {

	var insertUserSpec user.InsertUserSpec

	insertUserSpec.Name = req.Name
	insertUserSpec.Username = req.Username
	insertUserSpec.Password = req.Password

	return &insertUserSpec
}
