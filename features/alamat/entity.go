package alamat

type Core struct {
	ID        uint
	Provinsi  string
	Kota      string
	Kecamatan string
	Jalan     string
	Status    string
	UserId    uint
}

type ResponseCore struct {
	ID        uint
	Provinsi  string
	Kota      string
	Kecamatan string
	Jalan     string
	Status    string
	User      string
}

type UsecaseInterface interface {
	CreateAddress(data Core) (row int, err error)
	PutAddress(data Core, id, userId int) (row int, err error)
	DeleteAddress(id, userId int) (row int, err error)
	GetAllAddress(userId int) (data []ResponseCore, err error)
	GetAddress(id int) (data ResponseCore, err error)
}

type DataInterface interface {
	InsertAddress(data Core) (row int, err error)
	UpdateAdress(data Core, id, userId int) (row int, err error)
	DeleteDataAddress(id, userId int) (row int, err error)
	GetAllAddress(userId int) (data []ResponseCore, err error)
	GetAddress(id int) (data ResponseCore, err error)
}
