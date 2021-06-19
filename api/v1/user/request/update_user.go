package request

//UpdateUserRequest update User request payload
type UpdateUserRequest struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}
