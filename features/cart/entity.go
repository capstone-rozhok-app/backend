package cart

type Core struct {
	ID        uint
	Subtotal  int64
	Qty       uint
	Checklist bool
	UserId    uint
	ProdukId  uint
}

type ResponseCore struct {
	ID         uint
	Subtotal   int64
	Qty        uint
	Checklist  bool
	ProdFile   string
	ProdukName string
}

type UsecaseInterface interface {
	CreateCart(data Core) (row int, err error)
	UpdateCart(data Core, id, userId int) (row int, err error)
	GetAllCart(userId int) (data []ResponseCore, err error)
	DeleteCart(id, userId int) (row int, err error)
}

type DataInterface interface {
	CreateCart(data Core) (row int, err error)
	UpdateCart(data Core, id, userId int) (row int, err error)
	GetAllCart(userId int) (data []ResponseCore, err error)
	DeleteCart(id, userId int) (row int, err error)
}
