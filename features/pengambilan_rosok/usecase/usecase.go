package usecase

import (
	pengambilanrosok "rozhok/features/pengambilan_rosok"
)

type PengambilanRosok struct {
	Repo pengambilanrosok.PengambilanRosokData
}

func New(repo pengambilanrosok.PengambilanRosokData) *PengambilanRosok {
	return &PengambilanRosok{
		Repo: repo,
	}
}

func (u *PengambilanRosok) GetAll(TransaksiCore pengambilanrosok.Core) (rows []pengambilanrosok.Core, err error) {
	return u.Repo.GetAll(TransaksiCore)
}

func (u *PengambilanRosok) Get(TransaksiCore pengambilanrosok.Core) (row pengambilanrosok.Core, err error) {
	return u.Repo.Get(TransaksiCore)
}

func (u *PengambilanRosok) CreatePengambilanRosok(TransaksiCore pengambilanrosok.Core) (row int, err error) {
	return u.Repo.CreatePengambilanRosok(TransaksiCore)
}
