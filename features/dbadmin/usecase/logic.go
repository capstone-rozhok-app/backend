package usecase

import (
	"rozhok/features/dbadmin"
)

type authUsecase struct {
	authData dbadmin.DataInterface
}

func New(data dbadmin.DataInterface) dbadmin.UsecaseInterface {
	return &authUsecase{
		authData: data,
	}
}

func (usecase *authUsecase) GetUsers() (data dbadmin.ResponseCore, err error) {
	return usecase.authData.GetUsers()
}

func (usecase *authUsecase) GetTransaksiDetail(TransaksiCore dbadmin.TransaksiCore) (dbadmin.TransaksiCore, error) {
	return usecase.authData.GetTransaksiDetail(TransaksiCore)
}

func (usecase *authUsecase) GetTransaksi(TransaksiCore dbadmin.TransaksiCore) ([]dbadmin.TransaksiCore, error) {
	return usecase.authData.GetTransaksi(TransaksiCore)

}

func (usecase *authUsecase) UpdateTransaksi(TransaksiCore dbadmin.TransaksiCore) error {
	return usecase.authData.UpdateTransaksi(TransaksiCore)
}
