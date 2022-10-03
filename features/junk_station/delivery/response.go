package delivery

import (
	js "rozhok/features/junk_station"
)

type JSRes struct{
	JunkStationID	int			`json:"js_id"`
	JunkStationName	string		`json:"js_name"`
	JunkStationOwner string		`json:"js_owner"`
	Status			string 		`json:"status"`
	Provinsi		string		`json:"provinsi"`
	Kota			string 		`json:"kota"`
	Kecamatan		string 		`json:"kecamatan"`
	Telp			string		`json:"telp"`
	Jalan			string		`json:"jalan"`
}

func FromCore(data js.Core) JSRes {
	return JSRes{
		JunkStationID: data.JunkStationID,
		JunkStationName: data.JunkStationName,
		JunkStationOwner: data.JunkStationOwner,
		Status: data.Status,
		Provinsi: data.Provinsi,
		Kota: data.Kota,
		Kecamatan: data.Kecamatan,
		Telp: data.Telp,
		Jalan: data.Jalan,
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
		JunkStationName: data.JunkStationName,
		JunkStationOwner: data.JunkStationOwner,
		Status: data.Status,
		Provinsi: data.Provinsi,
		Kota: data.Kota,
		Kecamatan: data.Kecamatan,
		Telp: data.Telp,
		Jalan: data.Jalan,
	}
	return dataResponse
}