package delivery

import "rozhok/features/client"

type Response struct {
	TotaJual int `json:"total_penjualan"`
	Bonus    int `json:"bonus_terakhir"`
}

func fromCore(dataCore client.Core) Response {
	return Response{
		TotaJual: dataCore.TotaJual,
		Bonus:    int(dataCore.Bonus),
	}
}
