package delivery

import (
	js "rozhok/features/junk_station"
)

type JSRes struct {
	JunkStationID    int    `json:"id_junk_station"`
	JunkStationName  string `json:"junk_station_name"`
	JunkStationOwner string `json:"junk_station_owner"`
	StatusKemitraan  string `json:"status_kemitraan"`
	Provinsi         string `json:"provinsi"`
	Kota             string `json:"kota"`
	Kecamatan        string `json:"kecamatan"`
	Telp             string `json:"telp"`
	Jalan            string `json:"jalan,omitempty"`
	Image_url        string `json:"image_url,omitempty"`
}

func FromCore(data js.Core) JSRes {
	return JSRes{
		JunkStationID:    data.JunkStationID,
		JunkStationName:  data.JunkStationName,
		Provinsi:         data.Provinsi,
		Kota:             data.Kota,
		Kecamatan:        data.Kecamatan,
		JunkStationOwner: data.JunkStationOwner,
		Telp:             data.Telp,
		StatusKemitraan:  data.StatusKemitraan,
		Image_url: 		  data.Image_url,
	}
}

func CoreList(data []js.Core) []JSRes {
	var res []JSRes
	for _, v := range data {
		res = append(res, FromCore(v))
	}
	return res
}

func FromCoreToResponse(data js.Core) JSRes {
	dataResponse := JSRes{
		JunkStationID:    data.JunkStationID,
		JunkStationName:  data.JunkStationName,
		Provinsi:         data.Provinsi,
		Kota:             data.Kota,
		Kecamatan:        data.Kecamatan,
		JunkStationOwner: data.JunkStationOwner,
		Telp:             data.Telp,
		StatusKemitraan:  data.StatusKemitraan,
		Jalan:            data.Jalan,
		Image_url:        data.Image_url,
	}
	return dataResponse
}

func FromCoreResMitra(data js.Core) JSRes {
	return JSRes{
		JunkStationOwner: data.JunkStationOwner,
	}
}