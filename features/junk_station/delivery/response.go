package delivery

type JSRes struct{
	JunkStationID	int	`json:js_id`
	JunkStationName	string	`json:js_name`
	JunkStationOwner string	`json:js_owner`
	Provinsi		string	`json:provinsi`
	Kota			string `json:kota`
	Kecamatan		string `json:kecamatan`
	Telp			string	`json:telp`
	Jalan			string	`json:jalan`
}