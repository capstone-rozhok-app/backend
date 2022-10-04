package delivery

import mitra "rozhok/features/kemitraan"

type MitraRequest struct{
	StatusKemitraan		string	`json:"status_kemitraan" form:"status_kemitraan" validate:"required"`
}

func FromCoreReq(r MitraRequest) mitra.MitraCore {
	return mitra.MitraCore{
		StatusKemitraan: r.StatusKemitraan,
	}
}

func ToCore(r MitraRequest) mitra.MitraCore {
	return mitra.MitraCore{
		StatusKemitraan: r.StatusKemitraan,
	}
}