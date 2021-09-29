package auth

//Service outgoing port for user
type Service interface {
	//Login If data not found will return nil without error
	Login(username string, password string) (string, error)
}
