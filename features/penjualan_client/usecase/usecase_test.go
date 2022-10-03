package usecase

import (
	"errors"
	penjualanclient "rozhok/features/penjualan_client"
	"rozhok/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	request := penjualanclient.Core{}
	response := []penjualanclient.Core{
		{
			ID:           1,
			KategoriID:   1,
			NamaKategori: "Plastik Ringan",
		},
		{
			ID:           2,
			KategoriID:   2,
			NamaKategori: "Besi",
		},
	}
	t.Run("success get all data", func(t *testing.T) {
		repo := new(mocks.PenjualanClient)
		repo.On("GetAll", mock.Anything).Return(response, nil).Once()

		usecase := New(repo)

		results, err := usecase.GetAll(request)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(results))

		repo.AssertExpectations(t)
	})

	t.Run("failed get all data", func(t *testing.T) {
		repo := new(mocks.PenjualanClient)
		repo.On("GetAll", mock.Anything).Return([]penjualanclient.Core{}, errors.New("internal server error")).Once()

		usecase := New(repo)

		results, err := usecase.GetAll(request)

		assert.Error(t, err)
		assert.Equal(t, 0, len(results))

		repo.AssertExpectations(t)
	})
}

func TestStore(t *testing.T) {
	request := penjualanclient.Core{
		KategoriID: 1,
	}

	t.Run("success store", func(t *testing.T) {
		repo := new(mocks.PenjualanClient)
		repo.On("Insert", mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		results, err := usecase.Store(request)

		assert.NoError(t, err)
		assert.Equal(t, 1, results)

		repo.AssertExpectations(t)
	})

	t.Run("failed store", func(t *testing.T) {
		repo := new(mocks.PenjualanClient)
		repo.On("Insert", mock.Anything).Return(0, errors.New("internal server error")).Once()

		usecase := New(repo)

		results, err := usecase.Store(request)

		assert.Error(t, err)
		assert.Equal(t, 0, results)

		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	request := penjualanclient.Core{
		ID: 1,
	}

	t.Run("success update", func(t *testing.T) {
		repo := new(mocks.PenjualanClient)
		repo.On("Update", mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		results, err := usecase.Update(request)

		assert.NoError(t, err)
		assert.Equal(t, 1, results)

		repo.AssertExpectations(t)
	})

	t.Run("failed update", func(t *testing.T) {
		repo := new(mocks.PenjualanClient)
		repo.On("Update", mock.Anything).Return(0, errors.New("internal server error")).Once()

		usecase := New(repo)

		results, err := usecase.Update(request)

		assert.Error(t, err)
		assert.Equal(t, 0, results)

		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	request := penjualanclient.Core{
		ID: 1,
	}

	t.Run("success delete", func(t *testing.T) {
		repo := new(mocks.PenjualanClient)
		repo.On("Delete", mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		results, err := usecase.Delete(request)

		assert.NoError(t, err)
		assert.Equal(t, 1, results)

		repo.AssertExpectations(t)
	})

	t.Run("failed Delete", func(t *testing.T) {
		repo := new(mocks.PenjualanClient)
		repo.On("Delete", mock.Anything).Return(0, errors.New("internal server error")).Once()

		usecase := New(repo)

		results, err := usecase.Delete(request)

		assert.Error(t, err)
		assert.Equal(t, 0, results)

		repo.AssertExpectations(t)
	})
}
