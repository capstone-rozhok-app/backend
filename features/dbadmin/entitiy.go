package dbadmin

type GrafikData struct {
	Bulan    int `json:"bulan"`
	JumlahCL int `json:"jumlah_client"`
	JumlahJS int `json:"jumlah_junk_station"`
}

type ResponseCore struct {
	ID      int
	TotalJS int
	TotalCL int
	Grafik  []GrafikData
}

type TransaksiCore struct {
	ID            uint
	GrandTotal    int64
	TipeTransaksi string
	Status        string
	StartDate     string
	EndDate       string
	KodeTransaksi string
	Kurir         string

	Client  User
	Porter  User
	Product []Product
}

type User struct {
	ID        uint
	Name      string
	NoTelp    string
	Provinsi  string
	Kota      string
	Kecamatan string
}

type Product struct {
	ImageUrl    string
	ProductName string
	Qty         uint
	Subtotal    int64
}

type UsecaseInterface interface {
	GetUsers() (data ResponseCore, err error)
	GetTransaksiDetail(TransaksiCore TransaksiCore) (TransaksiCore, error)
	GetTransaksi(TransaksiCore TransaksiCore) ([]TransaksiCore, error)
	UpdateTransaksi(TransaksiCore TransaksiCore) error
}

type DataInterface interface {
	GetUsers() (data ResponseCore, err error)
	GetTransaksiDetail(TransaksiCore TransaksiCore) (TransaksiCore, error)
	GetTransaksi(TransaksiCore TransaksiCore) ([]TransaksiCore, error)
	UpdateTransaksi(TransaksiCore TransaksiCore) error
}
