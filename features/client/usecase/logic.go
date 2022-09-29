package usecase

import (
	"errors"
	"rozhok/features/client"
	"rozhok/middlewares"
	// "rozhok/middlewares"
)

type clientUsecase struct {
	clientData client.DataInterface
}

func New(dataClient client.DataInterface) client.UsecaseInterface {
	return &clientUsecase{
		clientData: dataClient,
	}
}

func (usecase *clientUsecase) CreateClient(client client.Core) (int, error) {
	if client.Email == "" || client.Password == "" {
		return -1, errors.New("email dan password tidak boleh kosong")
	}

	row, err := usecase.clientData.InsertClient(client)
	return row, err
}

func (usecase *clientUsecase) PutClient(newData client.Core, id int) (int, error) {

	row, err := usecase.clientData.UpdateClient(newData, id)
	return row, err
}

func (usecase *clientUsecase) LoginAuthorized(email, password string) (string, string, string) {

	if email == "" || password == "" {
		return "email dan password tidak boleh kosong", "", ""
	}

	results, errEmail := usecase.clientData.LoginClient(email)
	if errEmail != nil {
		return "email tidak ditemukan", "", ""
	}

	token, errToken := middlewares.CreateToken(int(results.ID), results.Role, "") // coba

	if errToken != nil {
		return "gagal membuat token", "", ""
	}

	return token, results.Role, results.Username

}

func (usecase *clientUsecase) DeleteClient(id int) (int, error) {
	row, err := usecase.clientData.DeleteDataClient(id)
	return row, err
}
