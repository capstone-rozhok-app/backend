package login

type Core struct {
	ID       int
	Email    string
	Password string
	Role     string
	Username string
	Status   string
}

type ResponseCore struct {
	ID      int
	TotalJS int
	TotalCL int
	Grafik  []Core
}

type UsecaseInterface interface {
	LoginAuthorized(email, password string) (string, string, string, string)
	GetUsers() (data ResponseCore, err error)
}

type DataInterface interface {
	LoginUser(email string) (Core, error)
	GetUsers() (data ResponseCore, err error)
}
