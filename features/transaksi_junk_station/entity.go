package transaksijunkstation

type Core struct {
	ID          uint
	GrandTotal  uint
	KodeTF      string
	CreatedAt   string
	BarangRosok []BarangRosok
}

type BarangRosok struct {
	ID       uint
	Kategori string
	Berat    uint
	Harga    int64
}

type TransaksiJunkStationData interface {
	GetAll(TransaksiJunkStationCore Core) ([]Core, error)
	Get(TransaksiJunkStationCore Core) (Core, error)
	Insert(TransaksiJunkStationCore Core) (int, error)
}

type TransaksiJunkStationUsecase interface {
	GetAll(TransaksiJunkStationCore Core) ([]Core, error)
	Get(TransaksiJunkStationCore Core) (Core, error)
	Create(TransaksiJunkStationCore Core) (int, error)
}
