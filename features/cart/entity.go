package cart

type Core struct {
	ID            uint
	Subtotal      int64
	Qty           uint
	Counter       string
	Checklist     string
	ChecklistBool bool
	UserId        uint
	ProdukId      uint
	ImageUrl      string
	ProdukName    string
}

type UsecaseInterface interface {
	CreateCart(data Core) (row int, err error)
	UpdateCart(data Core, id, userId int) (row int, err error)
	GetAllCart(userId int) (data []Core, err error)
	DeleteCart(id, userId int) (row int, err error)
}

type DataInterface interface {
	CreateCart(data Core) (row int, err error)
	UpdateCart(data Core, id, userId int) (row int, err error)
	GetAllCart(userId int) (data []Core, err error)
	DeleteCart(id, userId int) (row int, err error)
}
