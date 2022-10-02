package junkstation

type Core struct{
	Email				string ``
	Password			string
	JunkStationID		int
	JunkStationName		string
	JunkStationOwner	string
	Status				string
	Provinsi			string
	Kota				string
	Kecamatan			string
	Telp				string
	Jalan				string
}

type UsecaseInterface interface{
	CreateJunkStation (data Core, id int) (row int, err error)
	GetJunkStationAll () (data []Core, err error)
	GetJunkStationById (id int) (data Core, err error)
	PutJunkStation (id int, data Core) (row int, err error)
}

type DataInterface interface{
	InsertJunkStation (data Core) (row int, err error)
	FindJunkStationAll () (data []Core, err error)
	FindJunkStationById (id int) (data Core, err error)
	UpdateJunkStation (id int, data Core) (row int, err error)
}