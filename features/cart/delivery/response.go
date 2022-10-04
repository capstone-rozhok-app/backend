package delivery

import "rozhok/features/cart"

type CartResponse struct {
	ID         uint   `json:"id_cart" form:"id_cart"`
	ProdukName string `json:"product_name" form:"product_name"`
	ProdFile   string `json:"image_url" form:"image_url"`
	Subtotal   int64  `json:"price" form:"price"`
	Qty        uint   `json:"qty" form:"qty"`
	Checklist  bool   `json:"checklist" form:"checklist"`
}

func fromCore(dataCore cart.Core) CartResponse {
	return CartResponse{
		ID:         dataCore.ID,
		ProdukName: dataCore.ProdukName,
		ProdFile:   dataCore.ImageUrl,
		Subtotal:   dataCore.Subtotal,
		Qty:        dataCore.Qty,
		Checklist:  dataCore.ChecklistBool,
	}
}

func fromCoreList(dataCore []cart.Core) []CartResponse {
	var dataResponse []CartResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
