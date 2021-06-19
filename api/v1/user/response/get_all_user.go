package response

import "go-hexagonal/business/user"

type getAllUserResponse struct {
	Meta  meta              `json:"meta"`
	Users []GetUserResponse `json:"users"`
}

type meta struct {
	Page         int  `json:"page"`
	RowPerPage   int  `json:"row_per_page"`
	NextPage     bool `json:"next_page"`
	PreviousPage bool `json:"previous_page"`
}

//NewGetAllUserResponse construct GetAllUserResponse
func NewGetAllUserResponse(users []user.User, page int, rowPerPage int) getAllUserResponse {

	var (
		lenUsers = len(users)
	)

	getAllUserResponse := getAllUserResponse{}
	getAllUserResponse.Meta.Page = page
	getAllUserResponse.Meta.RowPerPage = rowPerPage
	getAllUserResponse.Meta.NextPage = false

	if lenUsers > rowPerPage {
		getAllUserResponse.Meta.NextPage = true
	}

	if (lenUsers-1 <= rowPerPage) && (page != 1) {
		getAllUserResponse.Meta.PreviousPage = true
	}

	for index, value := range users {
		if index == rowPerPage {
			continue
		}

		var getUserResponse GetUserResponse

		getUserResponse.ID = value.ID
		getUserResponse.Name = value.Name
		getUserResponse.Username = value.Username
		getUserResponse.ModifiedAt = value.ModifiedAt
		getUserResponse.Version = value.Version

		getAllUserResponse.Users = append(getAllUserResponse.Users, getUserResponse)
	}

	if getAllUserResponse.Users == nil {
		getAllUserResponse.Users = []GetUserResponse{}
	}

	return getAllUserResponse
}
