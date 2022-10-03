package usecase

import penjualanclient "rozhok/features/penjualan_client"

type PenjualanClient struct {
	Repo penjualanclient.PenjualanClientData
}

func New(repo penjualanclient.PenjualanClientData) *PenjualanClient {
	return &PenjualanClient{
		Repo: repo,
	}
}

func (u *PenjualanClient) GetAll(PenjualanClientCore penjualanclient.Core) ([]penjualanclient.Core, error) {
	return u.Repo.GetAll(PenjualanClientCore)
}

func (u *PenjualanClient) Update(PenjualanClientCore penjualanclient.Core) (int, error) {
	return u.Repo.Update(PenjualanClientCore)
}

func (u *PenjualanClient) Store(PenjualanClientCore penjualanclient.Core) (int, error) {
	return u.Repo.Insert(PenjualanClientCore)
}

func (u *PenjualanClient) Delete(PenjualanClientCore penjualanclient.Core) (int, error) {
	return u.Repo.Delete(PenjualanClientCore)
}
