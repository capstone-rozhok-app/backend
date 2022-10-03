package usecase

import (
	"errors"
	pengambilanrosok "rozhok/features/pengambilan_rosok"
	"rozhok/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	PengambilanRosok := pengambilanrosok.Core{}
	responseData := []pengambilanrosok.Core{
		{
			ID: 1,
			Client: pengambilanrosok.User{
				Provinsi:  "Jawa Timur",
				Kota:      "Banyuwangi",
				Kecamatan: "Bondo",
			},
		},
		{
			ID: 2,
			Client: pengambilanrosok.User{
				Provinsi:  "Jawa Timur",
				Kota:      "Banyuwangi",
				Kecamatan: "Bondo",
			},
		},
	}

	t.Run("success", func(t *testing.T) {
		repo := new(mocks.PengambilanRosokStruct)

		repo.On("GetAll", mock.Anything).Return(responseData, nil).Once()
		usecase := New(repo)

		rows, err := usecase.GetAll(PengambilanRosok)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(rows))

		repo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		repo := new(mocks.PengambilanRosokStruct)

		repo.On("GetAll", mock.Anything).Return([]pengambilanrosok.Core{}, errors.New("error")).Once()
		usecase := New(repo)

		rows, err := usecase.GetAll(PengambilanRosok)

		assert.Equal(t, "error", err.Error())
		assert.Equal(t, 0, len(rows))

		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	PengambilanRosok := pengambilanrosok.Core{}
	responseData := pengambilanrosok.Core{
		ID: 1,
	}

	t.Run("success", func(t *testing.T) {
		repo := new(mocks.PengambilanRosokStruct)

		repo.On("Get", mock.Anything).Return(responseData, nil).Once()
		usecase := New(repo)

		row, err := usecase.Get(PengambilanRosok)

		assert.NoError(t, err)
		assert.Equal(t, 1, int(row.ID))

		repo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		repo := new(mocks.PengambilanRosokStruct)

		repo.On("Get", mock.Anything).Return(pengambilanrosok.Core{}, errors.New("error")).Once()
		usecase := New(repo)

		row, err := usecase.Get(PengambilanRosok)

		assert.Equal(t, "error", err.Error())
		assert.Equal(t, 0, int(row.ID))

		repo.AssertExpectations(t)
	})
}

func TestPost(t *testing.T) {
	PengambilanRosok := pengambilanrosok.Core{}

	t.Run("success", func(t *testing.T) {
		repo := new(mocks.PengambilanRosokStruct)

		repo.On("CreatePengambilanRosok", mock.Anything).Return(1, nil).Once()
		usecase := New(repo)

		row, err := usecase.CreatePengambilanRosok(PengambilanRosok)

		assert.NoError(t, err)
		assert.Equal(t, 1, row)

		repo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		repo := new(mocks.PengambilanRosokStruct)

		repo.On("CreatePengambilanRosok", mock.Anything).Return(0, errors.New("error")).Once()
		usecase := New(repo)

		row, err := usecase.CreatePengambilanRosok(PengambilanRosok)

		assert.Equal(t, "error", err.Error())
		assert.Equal(t, 0, row)

		repo.AssertExpectations(t)
	})
}
