package data

import (
	"rozhok/features/cart"
	userModel "rozhok/features/login/data"
	produkModel "rozhok/features/produk/data"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Subtotal  int64
	Qty       uint
	Checklist bool
	UserId    uint
	ProdukId  uint
	User      userModel.User     `gorm:"foreignKey:UserId"`
	Produk    produkModel.Produk `gorm:"foreignKey:ProdukId"`
}

func fromCore(dataCore cart.Core) Cart {
	return Cart{
		Subtotal: dataCore.Subtotal,
		Qty:      dataCore.Qty,
		UserId:   dataCore.UserId,
		ProdukId: dataCore.ProdukId,
	}
}

func (dataCart *Cart) toCore() cart.Core {
	return cart.Core{
		ID:            dataCart.ID,
		Subtotal:      dataCart.Subtotal,
		Qty:           dataCart.Qty,
		ImageUrl:      dataCart.Produk.Image_url,
		ProdukName:    dataCart.Produk.Nama,
		ChecklistBool: dataCart.Checklist,
	}
}

func toCoreList(dataCart []Cart) []cart.Core {
	var dataCore []cart.Core

	for key := range dataCart {
		dataCore = append(dataCore, dataCart[key].toCore())

	}

	return dataCore

}
