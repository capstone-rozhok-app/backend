package usecase

import (
	"rozhok/features/cart"
)

type cartUsecase struct {
	cartData cart.DataInterface
}

func New(dataCart cart.DataInterface) cart.UsecaseInterface {

	return &cartUsecase{
		cartData: dataCart,
	}
}

func (usecase *cartUsecase) CreateCart(cart cart.Core) (int, error) {

	row, err := usecase.cartData.CreateCart(cart)
	return row, err
}

func (usecase *cartUsecase) GetAllCart(userId int) (data []cart.Core, err error) {
	results, err := usecase.cartData.GetAllCart(userId)
	return results, err
}

func (usecase *cartUsecase) UpdateCart(newData cart.Core, id, userId int) (int, error) {

	row, err := usecase.cartData.UpdateCart(newData, id, userId)
	return row, err
}

func (usecase *cartUsecase) DeleteCart(id, userId int) (int, error) {
	row, err := usecase.cartData.DeleteCart(id, userId)
	return row, err
}
