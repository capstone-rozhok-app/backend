package delivery

import (
	js "rozhok/features/junk_station"
)

type JsReq struct {
	JunkStationID		int	`json:"js_id" form:"js_id" validate:"required"`
	Email				string	`json:"email" form:"email" validate:"required"`
	Password			string	`json:"password" form:"password" validate:"required"`
	Status				string	`json:"status" form:"status" validate:"required"`
	JunkStationName		string	`json:"js_name" form:"js_name" validate:"required"`
	Username			string	`json:"js_owner" form:"js_owner" validate:"required"`
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
		Status: 			req.Status,
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
		Status: 			req.Status,
		Provinsi: 			req.Provinsi,
		Kota: 				req.Kota,
		Kecamatan: 			req.Kecamatan,	
		Telp: 				req.Telp,
		Jalan:				req.Jalan,
	}
}