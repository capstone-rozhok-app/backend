package usecase

import transaksijunkstation "rozhok/features/transaksi_junk_station"

type TransaksiJunkStation struct {
	Repo transaksijunkstation.TransaksiJunkStationData
}

func New(repo transaksijunkstation.TransaksiJunkStationData) *TransaksiJunkStation {
	return &TransaksiJunkStation{
		Repo: repo,
	}
}

func (u *TransaksiJunkStation) GetAll(TransaksiJunkStationCore transaksijunkstation.Core) ([]transaksijunkstation.Core, error) {
	return u.Repo.GetAll(TransaksiJunkStationCore)
}

func (u *TransaksiJunkStation) Get(TransaksiJunkStationCore transaksijunkstation.Core) (transaksijunkstation.Core, error) {
	return u.Repo.Get(TransaksiJunkStationCore)
}

func (u *TransaksiJunkStation) Create(TransaksiJunkStationCore transaksijunkstation.Core) (int, error) {
	return u.Repo.Insert(TransaksiJunkStationCore)
}
