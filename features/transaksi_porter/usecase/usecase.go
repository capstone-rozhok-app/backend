package usecase

import transaksiporter "rozhok/features/transaksi_porter"

type TransaksiPorter struct {
	Repo transaksiporter.TransaksiPorterData
}

func New(repo transaksiporter.TransaksiPorterData) *TransaksiPorter {
	return &TransaksiPorter{
		Repo: repo,
	}
}

func (usecase *TransaksiPorter) GetAll(TransaksiCore transaksiporter.Core) (rows []transaksiporter.Core, err error) {
	return usecase.Repo.GetAll(TransaksiCore)
}

func (usecase *TransaksiPorter) Get(TransaksiCore transaksiporter.Core) (row transaksiporter.Core, err error) {
	return usecase.Repo.Get(TransaksiCore)
}

func (usecase *TransaksiPorter) PostTransaksiPenjualan(TransaksiCore transaksiporter.Core) (row int, err error) {
	return usecase.Repo.CreateTransaksiPenjualan(TransaksiCore)
}

func (usecase *TransaksiPorter) PutTransaksiPembelian(TransaksiCore transaksiporter.Core) (row int, err error) {
	return usecase.Repo.UpdateTransaksiPembelian(TransaksiCore)
}
