package pembelianjs

type PembelianCore struct{
	ID				int
	Kategori		string
	Berat			int
	Harga			string
}

type UsecaseInterface interface{
	CreatePembelian (data PembelianCore) (row int, err error)
	GetPembelian () (data []PembelianCore, err error)
	PutPembelian (id int, data PembelianCore) (row int, err error)
	DeletePembelian (id int, data PembelianCore) (row int, err error)
}

type DataInterface interface{
	InsertPembelian (data PembelianCore) (row int, err error)
	FindPembelian () (data []PembelianCore, err error)
	UpdatePembelian (id int, data PembelianCore) (row int, err error)
	DeletePembelian (id int, data PembelianCore) (row int, err error)
}