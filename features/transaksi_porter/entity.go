package transaksiporter

type Core struct {
	ID                    uint
	PorterID              uint
	TipeTransaksi         string
	Status                string
	GrandTotal            float64
	StartDate             string
	EndDate               string
	DetailTransaksiPorter []DetailTransaksiPorter
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

type DetailTransaksiPorter struct {
	Id            uint
	Nama          string
	Berat         int64
	Subtotal      int64
	HargaKategori int64
}

type TransaksiPorterData interface {
	GetAll(TransaksiCore Core) (rows []Core, err error)
	Get(TransaksiCore Core) (row Core, err error)
	CreateTransaksiPenjualan(TransaksiCore Core) (row int, err error)
	UpdateTransaksiPembelian(TransaksiCore Core) (row int, err error)
}

type TransaksiPorterUsecase interface {
	GetAll(TransaksiCore Core) (rows []Core, err error)
	Get(TransaksiCore Core) (row Core, err error)
	PostTransaksiPenjualan(TransaksiCore Core) (row int, err error)
	PutTransaksiPembelian(TransaksiCore Core) (row int, err error)
}
