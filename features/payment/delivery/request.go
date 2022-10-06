package delivery

type Request struct {
	Bank  string `json:"bank" form:"bank" validate:"required"`
	Kurir string `json:"kurir" form:"kurir" validate:"required"`
}
