package delivery

import "rozhok/features/dbadmin"

type DboardResponse struct {
	TotalJS int `json:"total_junk_station" form:"total_junk_station"`
	TotalCL int `json:"total_client" form:"total_client"`
	Grafik  []dbadmin.GrafikData
}

type TransaksiResponse struct {
	IdTransaksi   uint     `json:"id_transaksi,omitempty"`
	KodeTransaksi string   `json:"kode_transaksi,omitempty"`
	Status        string   `json:"status,omitempty"`
	Client        Client   `json:"client,omitempty"`
	GrandTotal    int64    `json:"grand_total,omitempty"`
	Produk        []Produk `json:"produk,omitempty"`
}

type Client struct {
	Name      string `json:"name,omitempty"`
	NoTelp    string `json:"no_telp,omitempty"`
	Provinsi  string `json:"provinsi,omitempty"`
	Kota      string `json:"kota,omitempty"`
	Kecamatan string `json:"kecamatan,omitempty"`
}

type Produk struct {
	ImageUrl    string `json:"image_url,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	Qty         int    `json:"qty,omitempty"`
	Subtotal    int64  `json:"subtotal,omitempty"`
}

func FromCoreTransaksi(transaksiCore dbadmin.TransaksiCore) TransaksiResponse {
	transaksiResponse := TransaksiResponse{
		IdTransaksi:   transaksiCore.ID,
		KodeTransaksi: transaksiCore.KodeTransaksi,
		Status:        transaksiCore.Status,
		GrandTotal:    transaksiCore.GrandTotal,
		Client: Client{
			Name:      transaksiCore.Client.Name,
			NoTelp:    transaksiCore.Client.NoTelp,
			Provinsi:  transaksiCore.Client.Provinsi,
			Kota:      transaksiCore.Client.Kota,
			Kecamatan: transaksiCore.Client.Kecamatan,
		},
	}

	produkList := []Produk{}
	for _, produkcore := range transaksiCore.Product {
		produkList = append(produkList, Produk{
			ImageUrl:    produkcore.ImageUrl,
			ProductName: produkcore.ProductName,
			Qty:         int(produkcore.Qty),
			Subtotal:    produkcore.Subtotal,
		})
	}

	transaksiResponse.Produk = produkList
	return transaksiResponse
}

func fromCore(dataCore dbadmin.ResponseCore) DboardResponse {
	return DboardResponse{
		TotalJS: dataCore.TotalJS,
		TotalCL: dataCore.TotalCL,
		Grafik:  dataCore.Grafik,
	}
}
