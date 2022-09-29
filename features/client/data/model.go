package data

import (
	"rozhok/features/client"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Telp     string
	Alamat   []Alamat
}

type Alamat struct {
	gorm.Model
	ClientID uint
	Provinsi  string
	Kota      string
	Kecamatan string
	Jalan     string
	Status    string
	Client    Client
}

func fromCore(dataCore client.Core) Client {
	return Client{
		Telp:     dataCore.Telp,
	}
}

func toCore(data Client) client.Core {
	return client.Core{
		ID:       int(data.ID),
	}
}
