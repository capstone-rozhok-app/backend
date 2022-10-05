package usecase

import (
	"errors"
	transaksiclient "rozhok/features/transaksi_client"
	"rozhok/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	TransaksiClientCore := transaksiclient.Core{}
	responseData := []transaksiclient.Core{
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
		repo := new(mocks.TransaksiClient)

		repo.On("GetAll", mock.Anything).Return(responseData, nil).Once()
		usecase := New(repo)

		rows, err := usecase.GetAll(TransaksiClientCore)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(rows))

		repo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		repo := new(mocks.TransaksiClient)

		repo.On("GetAll", mock.Anything).Return([]transaksiclient.Core{}, errors.New("error")).Once()
		usecase := New(repo)

		rows, err := usecase.GetAll(TransaksiClientCore)

		assert.Equal(t, "error", err.Error())
		assert.Equal(t, 0, len(rows))

		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	TransaksiClientCore := transaksiclient.Core{}
	responseData := transaksiclient.Core{
		ID:            1,
		TipeTransaksi: "penjualan",
		Status:        "terjual",
	}

	t.Run("success", func(t *testing.T) {
		repo := new(mocks.TransaksiClient)

		repo.On("Get", mock.Anything).Return(responseData, nil).Once()
		usecase := New(repo)

		row, err := usecase.Get(TransaksiClientCore)

		assert.NoError(t, err)
		assert.Equal(t, 1, int(row.ID))

		repo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		repo := new(mocks.TransaksiClient)

		repo.On("Get", mock.Anything).Return(transaksiclient.Core{}, errors.New("error")).Once()
		usecase := New(repo)

		row, err := usecase.Get(TransaksiClientCore)

		assert.Equal(t, "error", err.Error())
		assert.Equal(t, 0, int(row.ID))

		repo.AssertExpectations(t)
	})
}

func TestPost(t *testing.T) {
	TransaksiClientCore := transaksiclient.Core{}

	t.Run("success", func(t *testing.T) {
		repo := new(mocks.TransaksiClient)

		repo.On("Insert", mock.Anything).Return(1, nil).Once()
		usecase := New(repo)

		row, err := usecase.Create(TransaksiClientCore)

		assert.NoError(t, err)
		assert.Equal(t, 1, row)

		repo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		repo := new(mocks.TransaksiClient)

		repo.On("Insert", mock.Anything).Return(0, errors.New("error")).Once()
		usecase := New(repo)

		row, err := usecase.Create(TransaksiClientCore)

		assert.Equal(t, "error", err.Error())
		assert.Equal(t, 0, row)

		repo.AssertExpectations(t)
	})
}

func TestPut(t *testing.T) {
	TransaksiClientCore := transaksiclient.Core{}

	t.Run("success", func(t *testing.T) {
		repo := new(mocks.TransaksiClient)

		repo.On("Update", mock.Anything).Return(1, nil).Once()
		usecase := New(repo)

		row, err := usecase.Update(TransaksiClientCore)

		assert.NoError(t, err)
		assert.Equal(t, 1, row)

		repo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		repo := new(mocks.TransaksiClient)

		repo.On("Update", mock.Anything).Return(0, errors.New("error")).Once()
		usecase := New(repo)

		row, err := usecase.Update(TransaksiClientCore)

		assert.Equal(t, "error", err.Error())
		assert.Equal(t, 0, row)

		repo.AssertExpectations(t)
	})
}
