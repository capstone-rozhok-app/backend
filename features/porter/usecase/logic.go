package usecase

import (
	"rozhok/features/porter"
	"time"

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

func (usecase *porterUsecase) GetPendapatan(porterCore porter.Core) (row porter.Core, err error) {

	switch porterCore.PeriodicFilter {
	case "harian":
		porterCore.StartDate = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		porterCore.EndDate = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	case "mingguan":
		porterCore.StartDate = time.Now().AddDate(0, 0, -7).Format("2006-01-02")
		porterCore.EndDate = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	case "bulanan":
		porterCore.StartDate = time.Now().AddDate(0, -1, 0).Format("2006-01-02")
		porterCore.EndDate = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	case "tahunan":
		porterCore.StartDate = time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
		porterCore.EndDate = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	}

	return usecase.porterData.GetPendapatan(porterCore)
}

func (usecase *porterUsecase) Get(id uint) (row porter.Core, err error) {
	return usecase.porterData.Get(id)
}
