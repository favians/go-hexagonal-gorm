package request

//UpdatePetRequest update Pet request payload
type UpdatePetRequest struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}
