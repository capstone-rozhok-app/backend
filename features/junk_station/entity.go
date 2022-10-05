package junkstation

type Core struct{
	Email				string ``
	Password			string
	JunkStationID		int
	JunkStationName		string
	JunkStationOwner	string
	StatusKemitraan		string
	Provinsi			string
	Kota				string
	Kecamatan			string
	Telp				string
	Jalan				string
	Image_url			string
}

type UsecaseInterface interface{
	CreateJunkStation (data Core) (row int, err error)
	GetJunkStationAll (dataCore Core) (row []Core, err error)
	GetJunkStationById (id int) (data Core, err error)
	PutJunkStation (id int, data Core) (row int, err error)
	PutKemitraan (id int)(row int, err error)
}

type DataInterface interface{
	InsertJunkStation (data Core) (row int, err error)
	FindJunkStationAll (dataCore Core) (row []Core, err error)
	FindJunkStationById (id int) (data Core, err error)
	UpdateJunkStation (id int, data Core) (row int, err error)
	UpdateKemitraan (id int)(row int, err error)
}