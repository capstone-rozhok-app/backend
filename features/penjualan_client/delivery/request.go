package delivery

type Request struct {
	IdKategori uint `json:"id_kategori" validate:"required"`
}
