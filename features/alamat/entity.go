package alamat

type Core struct {
	ID       int
	Email    string
	Password string
	Role     string
	Username string
	Foto     string
	Telepon  string
	Bonus    int64
}

type UsecaseInterface interface {
	CreateAddress(data Core) (row int, err error)
	PutAddress(data Core, id int) (row int, err error)
	DeleteAddress(id int) (row int, err error)
}

type DataInterface interface {
	InsertAddress(data Core) (row int, err error)
	UpdateAdress(data Core, id int) (row int, err error)
	DeleteDataAddress(id int) (row int, err error)
}
