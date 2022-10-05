package client

type Core struct {
	ID       int
	Email    string
	Password string
	Role     string
	Username string
	Foto     string
	Telepon  string
	Bonus    int64
	TotalJual int64
}

type UsecaseInterface interface {
	CreateClient(data Core) (row int, err error)
	PutClient(data Core, id int) (row int, err error)
	DeleteClient(id int) (row int, err error)
	GetClient(id int) (data Core, err error)
}

type DataInterface interface {
	InsertClient(data Core) (row int, err error)
	UpdateClient(data Core, id int) (row int, err error)
	DeleteDataClient(id int) (row int, err error)
	GetClient(id int) (data Core, err error)
}
