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

type ResponseCore struct {
	ID       int
	Nama     string
	Email    string
	Password string
	Username string
	Telp     string
	Role     string
	Alamat   string
}

type UsecaseInterface interface {
	CreateClient(data Core) (row int, err error)
	PutClient(data Core, id int) (row int, err error)
	DeleteClient(id int) (row int, err error)
}

type DataInterface interface {
	InsertClient(data Core) (row int, err error)
	UpdateClient(data Core, id int) (row int, err error)
	DeleteDataClient(id int) (row int, err error)
}
