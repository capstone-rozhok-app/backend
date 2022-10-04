package usecase

import (
	"rozhok/features/client"

	"golang.org/x/crypto/bcrypt"
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
	passwordHased, err := bcrypt.GenerateFromPassword([]byte(client.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	client.Password = string(passwordHased)

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

func (usecase *clientUsecase) GetClient(id int) (client.Core, error) {
	result, err := usecase.clientData.GetClient(id)
	if err != nil {
		return client.Core{}, err
	}
	return result, nil
}
