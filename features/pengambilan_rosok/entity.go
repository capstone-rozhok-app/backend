package pengambilanrosok

type Core struct {
	ID                    uint
	PorterID              uint
	TipeTransaksi         string
	Status                string
	DetailTransaksiClient []DetailTransaksiClient
	Client                User
}

type User struct {
	ID        uint
	Username  string
	Provinsi  string
	Kota      string
	Kecamatan string
	Jalan     string
	Telepon   string
}

type DetailTransaksiClient struct {
	Id         uint
	IdKategori uint
	Nama       string
}

type TransaksiPorterData interface {
	GetAll() (rows []Core, err error)
	Get(TransaksiCore Core) (row Core, err error)
	CreateTransaksiPenjualan(TransaksiCore Core) (row int, err error)
}

type TransaksiPorterUsecase interface {
	GetAll() (rows []Core, err error)
	Get(TransaksiCore Core) (row Core, err error)
	PostTransaksiPenjualan(TransaksiCore Core) (row int, err error)
}
