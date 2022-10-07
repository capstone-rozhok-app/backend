package junkstation

type Core struct {
	Email            string ``
	Password         string
	JunkStationID    int
	JunkStationName  string
	JunkStationOwner string
	StatusKemitraan  string
	Provinsi         string
	Kota             string
	Kecamatan        string
	Telp             string
	Jalan            string
	Image_url        string
	FilterPeriodic   string
	StartDate        string
	EndDate          string
}

type UsecaseInterface interface {
	Dashboard(data Core) (int64, error)
	CreateJunkStation(data Core) (row int, err error)
	GetJunkStationAll(dataCore Core) (row []Core, err error)
	GetJunkStationById(id int) (data Core, err error)
	PutJunkStation(id int, data Core) (row int, err error)
	PutKemitraan(id int) (row int, err error)
}

type DataInterface interface {
	Dashboard(data Core) (int64, error)
	InsertJunkStation(data Core) (row int, err error)
	FindJunkStationAll(dataCore Core) (row []Core, err error)
	FindJunkStationById(id int) (data Core, err error)
	UpdateJunkStation(id int, data Core) (row int, err error)
	UpdateKemitraan(id int) (row int, err error)
}
