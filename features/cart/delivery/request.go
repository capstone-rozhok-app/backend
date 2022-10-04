package delivery

import "rozhok/features/cart"

type CartRequest struct {
	Subtotal int64
	Qty      uint
	UserId   uint
	ProdukId uint `json:"id_barang" form:"id_barang" validate:"required"`
}

func toCore(dataRequest CartRequest) cart.Core {
	return cart.Core{
		Subtotal: dataRequest.Subtotal,
		Qty:      dataRequest.Qty,
		UserId:   dataRequest.UserId,
		ProdukId: dataRequest.ProdukId,
	}
}
