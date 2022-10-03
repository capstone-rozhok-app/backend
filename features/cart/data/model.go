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
		Subtotal:  dataCore.Subtotal,
		Qty:       dataCore.Qty,
		Checklist: dataCore.Checklist,
		UserId:    dataCore.UserId,
		ProdukId:  dataCore.ProdukId,
	}
}

func (dataCart *Cart) toCore() cart.ResponseCore {
	return cart.ResponseCore{
		ID:         dataCart.ID,
		Subtotal:   dataCart.Subtotal,
		Qty:        dataCart.Qty,
		Checklist:  dataCart.Checklist,
		ProdFile:   dataCart.Produk.Image_url,
		ProdukName: dataCart.Produk.Nama,
	}
}

func toCoreList(dataCart []Cart) []cart.ResponseCore {
	var dataCore []cart.ResponseCore

	for key := range dataCart {
		dataCore = append(dataCore, dataCart[key].toCore())

	}

	return dataCore

}
