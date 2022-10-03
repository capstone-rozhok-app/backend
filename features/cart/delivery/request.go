package delivery

import "rozhok/features/cart"

type CartRequest struct {
	Subtotal  int64
	Qty       uint `json:"qty" form:"qty"`
	Checklist bool
	UserId    uint
	ProdukId  uint `json:"id_barang" form:"id_barang"`
}

func toCore(dataRequest CartRequest) cart.Core {
	return cart.Core{
		Subtotal:  dataRequest.Subtotal,
		Qty:       dataRequest.Qty,
		Checklist: dataRequest.Checklist,
		UserId:    dataRequest.UserId,
		ProdukId:  dataRequest.ProdukId,
	}
}
