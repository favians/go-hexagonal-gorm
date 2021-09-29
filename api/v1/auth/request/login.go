package request

//Login Request payload
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
