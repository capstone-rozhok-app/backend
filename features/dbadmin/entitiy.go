package dbadmin

type GrafikData struct {
	Bulan    int `json:"bulan"`
	JumlahCL int `json:"jumlah_client"`
	JumlahJS int `json:"jumlah_junk_station"`
}

type ResponseCore struct {
	ID      int
	TotalJS int
	TotalCL int
	Grafik  []GrafikData
}

type UsecaseInterface interface {
	GetUsers() (data ResponseCore, err error)
}

type DataInterface interface {
	GetUsers() (data ResponseCore, err error)
}
