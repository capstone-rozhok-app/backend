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
