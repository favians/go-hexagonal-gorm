package response

import (
	"go-hexagonal/api/paginator"
	"go-hexagonal/business/user"
)

type getAllUserResponse struct {
	Meta  paginator.Meta    `json:"meta"`
	Users []GetUserResponse `json:"users"`
}

//NewGetAllUserResponse construct GetAllUserResponse
func NewGetAllUserResponse(users []user.User, page int, rowPerPage int) getAllUserResponse {

	var (
		lenUsers = len(users)
	)

	getAllUserResponse := getAllUserResponse{}
	getAllUserResponse.Meta.BuildMeta(lenUsers, page, rowPerPage)

	for i := 0; i < getAllUserResponse.Meta.RowPerPage; i++ {
		var getUserResponse GetUserResponse

		getUserResponse.ID = users[i].ID
		getUserResponse.Name = users[i].Name
		getUserResponse.Username = users[i].Username
		getUserResponse.ModifiedAt = users[i].ModifiedAt
		getUserResponse.Version = users[i].Version

		getAllUserResponse.Users = append(getAllUserResponse.Users, getUserResponse)
	}

	if getAllUserResponse.Users == nil {
		getAllUserResponse.Users = []GetUserResponse{}
	}

	return getAllUserResponse
}
