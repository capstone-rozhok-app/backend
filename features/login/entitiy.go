package login

type Core struct {
	ID       int
	Email    string
	Password string
	Role     string
	Username string
	Status   string
}

type UsecaseInterface interface {
	LoginAuthorized(email, password string) (string, string, string, string)
}

type DataInterface interface {
	LoginUser(email string) (Core, error)
}
