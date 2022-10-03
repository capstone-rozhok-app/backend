package delivery

type Request struct {
	IdKategori uint `json:"id_kategori" form:"id_kategori" validate:"required"`
}
