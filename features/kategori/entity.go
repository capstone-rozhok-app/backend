package kategori

type Core struct {
	ID           uint
	Nama         string
	Harga_mitra  int
	Harga_client int
	Desc         string
}

type UsecaseInterface interface {
	CreateKategori(data Core) (row int, err error)
	UpdateKategori(data Core, id int) (row int, err error)
	GetAllKategori() (data []Core, err error)
	DeleteKategori(id int) (row int, err error)
}

type DataInterface interface {
	CreateKategori(data Core) (row int, err error)
	UpdateKategori(data Core, id int) (row int, err error)
	GetAllKategori() (data []Core, err error)
	DeleteKategori(id int) (row int, err error)
}
