package data

import (
	"errors"
	"rozhok/features/cart"
	produkModel "rozhok/features/produk/data"
	"strconv"

	"gorm.io/gorm"
)

type cartData struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.DataInterface {
	return &cartData{
		db: db,
	}
}

func (repo *cartData) CreateCart(cart cart.Core) (int, error) {
	var dbCek produkModel.Produk
	repo.db.First(&dbCek, "id = ?", cart.ProdukId)

	cartModel := fromCore(cart)
	cartModel.Qty = 1
	cartModel.Subtotal = dbCek.Harga * int64(cartModel.Qty)

	tx := repo.db.Create(&cartModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *cartData) GetAllCart(userId int) ([]cart.Core, error) {
	var allCartData []Cart
	tx := repo.db.Where("user_id = ?", userId).Joins("User").Joins("Produk").Find(&allCartData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return toCoreList(allCartData), nil
}

func (repo *cartData) UpdateCart(data cart.Core, id, userId int) (row int, err error) {
	var dbCek Cart
	txCek := repo.db.Preload("Produk").First(&dbCek, "id = ?", id)
	if txCek.Error != nil {
		return 0, txCek.Error
	}

	if data.Counter == "increment" {
		dbCek.Qty += 1
	} else if dbCek.Qty > 1 {
		dbCek.Qty -= 1
	}

	var dataModel Cart
	dataModel.Qty = dbCek.Qty
	dataModel.Subtotal = dbCek.Produk.Harga * int64(dbCek.Qty)

	if data.Checklist != "" {
		checklist, err := strconv.ParseBool(data.Checklist)
		if err != nil {
			return 0, err
		}
		dataModel.Checklist = checklist
	}

	tx := repo.db.Model(&Cart{}).Where("id = ?", id).Select("qty", "checklist").Updates(&dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal meperbarui data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *cartData) DeleteCart(id, userId int) (row int, err error) {
	tx := repo.db.Where("id = ?", id).Where("user_id = ?", userId).Delete(&Cart{})
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal menghapus akun")
	}
	return int(tx.RowsAffected), nil
}
