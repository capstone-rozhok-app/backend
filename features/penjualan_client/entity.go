package penjualanclient

type Core struct {
	ID           uint
	KategoriID   uint
	ClientID     uint
	NamaKategori string
}

type PenjualanClientData interface {
	GetAll(PenjualanClientCore Core) ([]Core, error)
	Update(PenjualanClientCore Core) (int, error)
	Insert(PenjualanClientCore Core) (int, error)
	Delete(PenjualanClientCore Core) (int, error)
}

type PenjualanClientUsecase interface {
	GetAll(PenjualanClientCore Core) ([]Core, error)
	Update(PenjualanClientCore Core) (int, error)
	Store(PenjualanClientCore Core) (int, error)
	Delete(PenjualanClientCore Core) (int, error)
}
