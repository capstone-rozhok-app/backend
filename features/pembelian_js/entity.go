package pembelianjs

type PembelianCore struct{
	IDPembelian		int
	JunkStationID	int
	Kategori		int
	NamaKategori	string
	Berat			int
	Harga			int
}

type UsecaseInterface interface{
	CreatePembelian (data PembelianCore) (row int, err error)
	GetPembelian (JunkStationID int) (data []PembelianCore, err error)
	PutPembelian (id int, data PembelianCore) (row int, err error)
	DeletePembelian (id int, data PembelianCore) (row int, err error)
}

type DataInterface interface{
	InsertPembelian (data PembelianCore) (row int, err error)
	FindPembelian (JunkStationID int) (data []PembelianCore, err error)
	UpdatePembelian (id int, data PembelianCore) (row int, err error)
	DeletePembelian (id int, data PembelianCore) (row int, err error)
}