package delivery

import "rozhok/features/client"

type Response struct {
	TotalJual int64 `json:"total_penjualan"`
	Bonus     int64 `json:"bonus_terakhir"`
}

func fromCore(dataCore client.Core) Response {
	return Response{
		TotalJual: dataCore.TotalJual,
		Bonus:     dataCore.Bonus,
	}
}
