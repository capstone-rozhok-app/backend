package transaksijunkstation

type Core struct {
	ID          uint
	UserID      uint
	GrandTotal  int64
	KodeTF      string
	CreatedAt   string
	StartDate   string
	EndDate     string
	BarangRosok []BarangRosok
}

type BarangRosok struct {
	ID       uint
	Kategori string
	Berat    uint
	Subtotal int64
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
