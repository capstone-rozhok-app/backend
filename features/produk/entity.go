package produk

type Core struct {
	ID        uint
	Nama      string
	Image_url string
	Stok      int
	Harga     int64
	Desc      string
}

type UsecaseInterface interface {
	CreateProduk(data Core) (row int, err error)
	UpdateProduk(data Core, id int) (row int, err error)
	GetAllProduk() (data []Core, err error)
	GetProduk(id int) (data Core, err error)
	DeleteProduk(id int) (row int, err error)
	GetFavorite() (data []Core, err error)
}

type DataInterface interface {
	CreateProduk(data Core) (row int, err error)
	UpdateProduk(data Core, id int) (row int, err error)
	GetAllProduk() (data []Core, err error)
	GetProduk(id int) (data Core, err error)
	DeleteProduk(id int) (row int, err error)
	GetFavorite() (data []Core, err error)
}
