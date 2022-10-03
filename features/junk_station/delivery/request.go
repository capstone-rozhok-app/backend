package delivery

import (
	js "rozhok/features/junk_station"
)

type JsReq struct {
	Email				string	`json:"email" form:"email" validate:"required"`
	Password			string	`json:"password" form:"password" validate:"required"`
	JunkStationName		string	`json:"junk_station_name" form:"junk_station_name" validate:"required"`
	Username			string	`json:"junk_station_owner" form:"junk_station_owner" validate:"required"`
	Provinsi			string	`json:"provinsi" form:"provinsi" validate:"required"`
	Kota				string	`json:"kota" form:"kota" validate:"required"`
	Kecamatan			string	`json:"kecamatan" form:"kecamatan" validate:"required"`
	Telp				string	`json:"telp" form:"telp" validate:"required"`
	Jalan				string	`json:"jalan" form:"jalan" validate:"required"`
}

func FromCoreReq(req JsReq) js.Core{
	return js.Core{
		Email: 				req.Email,
		Password: 			req.Password,	
		JunkStationName: 	req.JunkStationName,
		JunkStationOwner: 	req.Username,
		Provinsi: 			req.Provinsi,
		Kota: 				req.Kota,
		Kecamatan: 			req.Kecamatan,	
		Telp: 				req.Telp,
		Jalan:				req.Jalan,
	}
}

func ToCoreReq(req JsReq) js.Core{
	return js.Core{
		Email: 				req.Email,
		Password: 			req.Password,	
		JunkStationName: 	req.JunkStationName,
		JunkStationOwner: 	req.Username,
		Provinsi: 			req.Provinsi,
		Kota: 				req.Kota,
		Kecamatan: 			req.Kecamatan,	
		Telp: 				req.Telp,
		Jalan:				req.Jalan,
	}
}