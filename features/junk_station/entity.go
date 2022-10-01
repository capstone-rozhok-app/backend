package junkstation

type Core struct{
	JunkStationID		int
	JunkStationName		string
	JunkStationOwner	string
	Provinsi			string
	Kota				string
	Kecamatan			string
	Telp				string
	Jalan				string
}

type UsecaseInterface interface{
	CreateJunkStation (data Core) (row int, err error)
	GetJunkStationAll () (data []Core, err error)
	GetJunkStationById (id int) (data Core, err error)
	PutJunkStation (id int) (data Core, err error)
}

type DataInterface interface{
	InsertJunkStation (data Core) (row int, err error)
	FindJunkStationAll () (data []Core, err error)
	FindJunkStationById (id int) (data Core, err error)
	UpdateJunkStation (id int) (data Core, err error)
}