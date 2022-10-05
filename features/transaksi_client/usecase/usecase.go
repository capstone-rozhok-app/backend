package usecase

import transaksiclient "rozhok/features/transaksi_client"

type TransaksiClient struct {
	Repo transaksiclient.TransaksiClientData
}

func New(repo transaksiclient.TransaksiClientData) *TransaksiClient {
	return &TransaksiClient{
		Repo: repo,
	}
}

func (r *TransaksiClient) GetAll(TransaksiClient transaksiclient.Core) ([]transaksiclient.Core, error) {
	return r.Repo.GetAll(TransaksiClient)
}

func (r *TransaksiClient) Get(TransaksiClient transaksiclient.Core) (transaksiclient.Core, error) {
	return r.Repo.Get(TransaksiClient)
}

func (r *TransaksiClient) Create(TransaksiClient transaksiclient.Core) (int, error) {
	return r.Repo.Insert(TransaksiClient)
}

func (r *TransaksiClient) Update(TransaksiClient transaksiclient.Core) (int, error) {
	return r.Repo.Update(TransaksiClient)
}
