package usecase

import (
	"rozhok/features/login"
	"rozhok/middlewares"
)

type authUsecase struct {
	authData login.DataInterface
}

func New(data login.DataInterface) login.UsecaseInterface {
	return &authUsecase{
		authData: data,
	}
}

func (usecase *authUsecase) LoginAuthorized(email, password string) (string, string, string, string) {

	if email == "" || password == "" {
		return "please input email and password", "", "", ""
	}

	results, errEmail := usecase.authData.LoginUser(email)
	if errEmail != nil {
		return "email not found", "", "", ""
	}

	token, errToken := middlewares.CreateToken(int(results.ID), results.Role, results.Status)

	if errToken != nil {
		return "error to created token", "", "", ""
	}

	return token, results.Role, results.Username, results.Status

}

func (usecase *authUsecase) GetUsers() (data login.ResponseCore, err error) {
	return usecase.authData.GetUsers()
}
