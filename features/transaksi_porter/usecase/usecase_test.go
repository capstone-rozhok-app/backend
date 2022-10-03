package usecase

import (
	"errors"
	transaksiporter "rozhok/features/transaksi_porter"
	"rozhok/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	TransaksiPorterCore := transaksiporter.Core{}
	responseData := []transaksiporter.Core{
		{
			ID:            1,
			TipeTransaksi: "pembelian",
			Status:        "belum_bayar",
		},
		{
			ID:            1,
			TipeTransaksi: "penjualan",
			Status:        "terjual",
		},
	}

	t.Run("success", func(t *testing.T) {
		repo := new(mocks.TransaksiPorterRepo)

		repo.On("GetAll", mock.Anything).Return(responseData, nil).Once()
		usecase := New(repo)

		rows, err := usecase.GetAll(TransaksiPorterCore)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(rows))

		repo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		repo := new(mocks.TransaksiPorterRepo)

		repo.On("GetAll", mock.Anything).Return([]transaksiporter.Core{}, errors.New("error")).Once()
		usecase := New(repo)

		rows, err := usecase.GetAll(TransaksiPorterCore)

		assert.Equal(t, "error", err.Error())
		assert.Equal(t, 0, len(rows))

		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	TransaksiPorterCore := transaksiporter.Core{}
	responseData := transaksiporter.Core{
		ID:            1,
		TipeTransaksi: "penjualan",
		Status:        "terjual",
	}

	t.Run("success", func(t *testing.T) {
		repo := new(mocks.TransaksiPorterRepo)

		repo.On("Get", mock.Anything).Return(responseData, nil).Once()
		usecase := New(repo)

		row, err := usecase.Get(TransaksiPorterCore)

		assert.NoError(t, err)
		assert.Equal(t, 1, int(row.ID))

		repo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		repo := new(mocks.TransaksiPorterRepo)

		repo.On("Get", mock.Anything).Return(transaksiporter.Core{}, errors.New("error")).Once()
		usecase := New(repo)

		row, err := usecase.Get(TransaksiPorterCore)

		assert.Equal(t, "error", err.Error())
		assert.Equal(t, 0, int(row.ID))

		repo.AssertExpectations(t)
	})
}

func TestPost(t *testing.T) {
	TransaksiPorterCore := transaksiporter.Core{}

	t.Run("success", func(t *testing.T) {
		repo := new(mocks.TransaksiPorterRepo)

		repo.On("CreateTransaksiPenjualan", mock.Anything).Return(1, nil).Once()
		usecase := New(repo)

		row, err := usecase.PostTransaksiPenjualan(TransaksiPorterCore)

		assert.NoError(t, err)
		assert.Equal(t, 1, row)

		repo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		repo := new(mocks.TransaksiPorterRepo)

		repo.On("CreateTransaksiPenjualan", mock.Anything).Return(0, errors.New("error")).Once()
		usecase := New(repo)

		row, err := usecase.PostTransaksiPenjualan(TransaksiPorterCore)

		assert.Equal(t, "error", err.Error())
		assert.Equal(t, 0, row)

		repo.AssertExpectations(t)
	})
}

func TestPut(t *testing.T) {
	TransaksiPorterCore := transaksiporter.Core{}

	t.Run("success", func(t *testing.T) {
		repo := new(mocks.TransaksiPorterRepo)

		repo.On("UpdateTransaksiPembelian", mock.Anything).Return(1, nil).Once()
		usecase := New(repo)

		row, err := usecase.PutTransaksiPembelian(TransaksiPorterCore)

		assert.NoError(t, err)
		assert.Equal(t, 1, row)

		repo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		repo := new(mocks.TransaksiPorterRepo)

		repo.On("UpdateTransaksiPembelian", mock.Anything).Return(0, errors.New("error")).Once()
		usecase := New(repo)

		row, err := usecase.PutTransaksiPembelian(TransaksiPorterCore)

		assert.Equal(t, "error", err.Error())
		assert.Equal(t, 0, row)

		repo.AssertExpectations(t)
	})
}
