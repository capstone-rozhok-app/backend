package usecase

import (
	"rozhok/features/porter"

	"golang.org/x/crypto/bcrypt"
)

type porterUsecase struct {
	porterData porter.DataInterface
}

func New(dataPorter porter.DataInterface) porter.UsecaseInterface {
	return &porterUsecase{
		porterData: dataPorter,
	}
}

func (usecase *porterUsecase) CreatePorter(porter porter.Core) (int, error) {
	passwordHased, err := bcrypt.GenerateFromPassword([]byte(porter.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	porter.Password = string(passwordHased)
	return usecase.porterData.InsertPorter(porter)
}

func (usecase *porterUsecase) UpdatePorter(porter porter.Core, id int) (row int, err error) {
	return usecase.porterData.UpdatePorter(porter, uint(id))
}

func (usecase *porterUsecase) DeletePorter(id uint) (row int, err error) {
	return usecase.porterData.DeletePorter(uint(id))
}

func (usecase *porterUsecase) GetAll() (rows []porter.Core, err error) {
	return usecase.porterData.GetAll()
}

func (usecase *porterUsecase) GetPendapatan(id uint) (row porter.Core, err error) {
	return usecase.porterData.Get(id)
}

func (usecase *porterUsecase) Get(id uint) (row porter.Core, err error) {
	return usecase.porterData.Get(id)
}
