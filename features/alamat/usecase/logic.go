package usecase

import (
	"errors"
	"rozhok/features/client"
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

func (usecase *clientUsecase) DeleteClient(id int) (int, error) {
	row, err := usecase.clientData.DeleteDataClient(id)
	return row, err
}
