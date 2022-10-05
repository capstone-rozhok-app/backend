package transaksiclient

type Core struct {
	ID            uint
	GrandTotal    int64
	TipeTransaksi string
	Status        string
	StartDate     string
	EndDate       string
	KodeTransaksi string
	Kurir         string

	Client      User
	Porter      User
	Tagihan     Tagihan
	BarangRosok []BarangRosok
	Product     []Product
}

type User struct {
	ID        uint
	Name      string
	NoTelp    string
	Provinsi  string
	Kota      string
	Kecamatan string
}

type BarangRosok struct {
	Kategori string
	Berat    int64
	Harga    int64
}

type Tagihan struct {
	NoVA           string
	TipePembayaran string
	Bank           string
	GrandTotal     int64
}

type Product struct {
	ImageUrl    string
	ProductName string
	Qty         uint
	Subtotal    int64
}

type TransaksiClientData interface {
	GetAll(TransaksiClient Core) ([]Core, error)
	Get(TransaksiClient Core) (Core, error)
	Insert(TransaksiClient Core) (int, error)
	Update(TransaksiClient Core) (int, error)
}

type TransaksiClientUsecase interface {
	GetAll(TransaksiClient Core) ([]Core, error)
	Get(TransaksiClient Core) (Core, error)
	Create(TransaksiClient Core) (int, error)
	Update(TransaksiClient Core) (int, error)
}
