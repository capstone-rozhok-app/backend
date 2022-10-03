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

type PengambilanRosokData interface {
	GetAll(TransaksiCore Core) (rows []Core, err error)
	Get(TransaksiCore Core) (row Core, err error)
	CreatePengambilanRosok(TransaksiCore Core) (row int, err error)
}

type PengambilanRosokUsecase interface {
	GetAll(TransaksiCore Core) (rows []Core, err error)
	Get(TransaksiCore Core) (row Core, err error)
	CreatePengambilanRosok(TransaksiCore Core) (row int, err error)
}
