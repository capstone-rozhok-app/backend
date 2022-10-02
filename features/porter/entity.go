package porter

type Core struct {
	ID             uint
	Username       string
	Email          string
	Password       string
	Telp           string
	Provinsi       string
	Kota           string
	Kecamatan      string
	Jalan          string
	Laba           int64
	TotalPenjualan int64
	TotalPembelian int64
	PeriodicFilter string
	StartDate      string
	EndDate        string
}

type UsecaseInterface interface {
	CreatePorter(porter Core) (row int, err error)
	UpdatePorter(porter Core, id int) (row int, err error)
	DeletePorter(id uint) (row int, err error)
	GetPendapatan(porter Core) (row Core, err error)
	GetAll() (rows []Core, err error)
	Get(id uint) (row Core, err error)
}

type DataInterface interface {
	InsertPorter(porter Core) (row int, err error)
	UpdatePorter(porter Core, id uint) (row int, err error)
	DeletePorter(id uint) (row int, err error)
	GetPendapatan(porter Core) (row Core, err error)
	GetAll() (rows []Core, err error)
	Get(id uint) (row Core, err error)
}
