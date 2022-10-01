package usecase

import (
	"errors"
	"rozhok/features/porter"
	"rozhok/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsert(t *testing.T) {
	requestedData := porter.Core{
		Username:  "ucup",
		Email:     "ucup@gmail.com",
		Password:  "thomasSlebew",
		Telp:      "08910237014",
		Provinsi:  "Jawa Timur",
		Kota:      "Surabaya",
		Kecamatan: "Rungkut",
		Jalan:     "jl.Aja sendirian",
	}

	t.Run("success insert", func(t *testing.T) {
		repo := new(mocks.PorterRepo)
		repo.On("InsertPorter", mock.Anything).Return(1, nil).Once()

		usecase := New(repo)
		row, err := usecase.CreatePorter(requestedData)

		assert.NoError(t, err)
		assert.NotEqual(t, 0, row)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	requestedData := porter.Core{
		Username:  "ucup",
		Email:     "ucup@gmail.com",
		Password:  "thomasSlebew",
		Telp:      "08910237014",
		Provinsi:  "Jawa Timur",
		Kota:      "Surabaya",
		Kecamatan: "Rungkut",
		Jalan:     "jl.Aja sendirian",
	}

	t.Run("success update", func(t *testing.T) {
		repo := new(mocks.PorterRepo)
		repo.On("UpdatePorter", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)
		row, err := usecase.UpdatePorter(requestedData, 1)

		assert.NoError(t, err)
		assert.NotEqual(t, 0, row)
		repo.AssertExpectations(t)
	})

	t.Run("error update", func(t *testing.T) {
		repo := new(mocks.PorterRepo)
		repo.On("UpdatePorter", mock.Anything, mock.Anything).Return(0, errors.New("internal server error")).Once()

		usecase := New(repo)
		row, err := usecase.UpdatePorter(requestedData, 1)

		assert.Error(t, err)
		assert.Equal(t, 0, row)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success delete", func(t *testing.T) {
		repo := new(mocks.PorterRepo)
		repo.On("DeletePorter",mock.Anything).Return(1, nil).Once()

		usecase := New(repo)
		row, err := usecase.DeletePorter(1)

		assert.NoError(t, err)
		assert.NotEqual(t, 0, row)
		repo.AssertExpectations(t)
	})

	t.Run("error delete", func(t *testing.T) {
		repo := new(mocks.PorterRepo)
		repo.On("DeletePorter", mock.Anything).Return(0, errors.New("internal server error")).Once()

		usecase := New(repo)
		row, err := usecase.DeletePorter(1)

		assert.Error(t, err)
		assert.Equal(t, 0, row)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	responseDataList := []porter.Core{
		{
			ID: 1,
			Username:  "ucup",
			Email:     "ucup@gmail.com",
			Password:  "thomasSlebew",
			Telp:      "08910237014",
			Provinsi:  "Jawa Timur",
			Kota:      "Surabaya",
			Kecamatan: "Rungkut",
			Jalan:     "jl.Aja sendirian",
		},
	}

	t.Run("success get data", func(t *testing.T) {
		repo := new(mocks.PorterRepo)
		repo.On("GetAll").Return(responseDataList, nil).Once()

		usecase := New(repo)
		rows, err := usecase.GetAll()

		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(rows))
		repo.AssertExpectations(t)
	})

	t.Run("error get all", func(t *testing.T) {
		repo := new(mocks.PorterRepo)
		repo.On("GetAll").Return([]porter.Core{}, errors.New("internal server error")).Once()

		usecase := New(repo)
		rows, err := usecase.GetAll()

		assert.Error(t, err)
		assert.Equal(t, 0, len(rows))
		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	responseData := porter.Core{
		ID: 1,
		Username:  "ucup",
		Email:     "ucup@gmail.com",
		Password:  "thomasSlebew",
		Telp:      "08910237014",
		Provinsi:  "Jawa Timur",
		Kota:      "Surabaya",
		Kecamatan: "Rungkut",
		Jalan:     "jl.Aja sendirian",
	}

	t.Run("success get data", func(t *testing.T) {
		repo := new(mocks.PorterRepo)
		repo.On("Get", mock.Anything).Return(responseData, nil).Once()

		usecase := New(repo)
		row, err := usecase.Get(uint(1))

		assert.NoError(t, err)
		assert.NotEqual(t, 1, row.ID)
		repo.AssertExpectations(t)
	})

	t.Run("error get data", func(t *testing.T) {
		repo := new(mocks.PorterRepo)
		repo.On("Get", mock.Anything).Return(porter.Core{}, errors.New("internal server error")).Once()

		usecase := New(repo)
		row, err := usecase.Get(uint(1))

		assert.Error(t, err)
		assert.Equal(t, uint(0), row.ID)
		repo.AssertExpectations(t)
	})
}
