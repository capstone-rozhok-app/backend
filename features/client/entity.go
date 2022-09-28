package client

type Core struct {
	ID       int
	Nama     string
	Email    string
	Password string
	Username string
	Telp     string
	Role     string
	AlamatId int
}

type UsecaseInterface interface {
	CreateClient(data Core) (row int, err error)
}

type DataInterface interface {
	InsertClient(data Core) (row int, err error)
}
