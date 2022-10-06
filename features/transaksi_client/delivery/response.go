package delivery

import transaksiclient "rozhok/features/transaksi_client"

type Response struct {
	IdTransaksi   uint          `json:"id_transaksi,omitempty"`
	KodeTransaksi string        `json:"kode_transaksi,omitempty"`
	TipeTransaksi string        `json:"tipe_transaksi,omitempty"`
	Status        string        `json:"status,omitempty"`
	Client        Client        `json:"client,omitempty"`
	Porter        Porter        `json:"porter,omitempty"`
	GrandTotal    int64         `json:"grand_total,omitempty"`
	Kurir         string        `json:"kurir,omitempty"`
	BarangRosok   []BarangRosok `json:"barang_rosok,omitempty"`
	Produk        []Produk      `json:"produk,omitempty"`
}

type Client struct {
	Name      string `json:"name,omitempty"`
	NoTelp    string `json:"no_telp,omitempty"`
	Provinsi  string `json:"provinsi,omitempty"`
	Kota      string `json:"kota,omitempty"`
	Kecamatan string `json:"kecamatan,omitempty"`
}

type Porter struct {
	Name   string `json:"name,omitempty"`
	NoTelp string `json:"no_telp,omitempty"`
}

type BarangRosok struct {
	Kategori string `json:"kategori,omitempty"`
	Berat    uint   `json:"berat,omitempty"`
	Harga    int64  `json:"harga,omitempty"`
}

type Produk struct {
	ImageUrl    string `json:"image_url,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	Qty         int    `json:"qty,omitempty"`
	Subtotal    int64  `json:"subtotal,omitempty"`
}

type Tagihan struct {
	NoVa           string `json:"no_va"`
	TipePembayaran string `json:"tipe_pembayaran"`
	TotalHarga     int64  `json:"total_harga"`
	Bank           string `json:"bank"`
}

func ToTagihan(TagihanCore transaksiclient.Tagihan) Tagihan {
	return Tagihan{
		NoVa:           TagihanCore.NoVA,
		TipePembayaran: TagihanCore.TipePembayaran,
		TotalHarga:     TagihanCore.GrandTotal,
		Bank:           TagihanCore.Bank,
	}
}

func FromCore(transaksiCore transaksiclient.Core) Response {
	transaksiResponse := Response{
		IdTransaksi:   transaksiCore.ID,
		KodeTransaksi: transaksiCore.KodeTransaksi,
		TipeTransaksi: transaksiCore.TipeTransaksi,
		Status:        transaksiCore.Status,
		Kurir:         transaksiCore.Kurir,
		GrandTotal:    transaksiCore.GrandTotal,
		Client: Client{
			Name:      transaksiCore.Client.Name,
			NoTelp:    transaksiCore.Client.NoTelp,
			Provinsi:  transaksiCore.Client.Provinsi,
			Kota:      transaksiCore.Client.Kota,
			Kecamatan: transaksiCore.Client.Kecamatan,
		},
		Porter: Porter{
			Name:   transaksiCore.Porter.Name,
			NoTelp: transaksiCore.Porter.NoTelp,
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

	barangRosokList := []BarangRosok{}
	for _, barangrosokcore := range transaksiCore.BarangRosok {
		barangRosokList = append(barangRosokList, BarangRosok{
			Kategori: barangrosokcore.Kategori,
			Berat:    uint(barangrosokcore.Berat),
			Harga:    barangrosokcore.Harga,
		})
	}

	transaksiResponse.Produk = produkList
	transaksiResponse.BarangRosok = barangRosokList

	return transaksiResponse
}
