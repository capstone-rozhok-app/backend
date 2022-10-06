package usecase

import "rozhok/features/payment"

type Payment struct {
	Repo payment.PaymentData
}

func New(repo payment.PaymentData) *Payment {
	return &Payment{
		Repo: repo,
	}
}

func (r *Payment) Create(PaymentData payment.Core) (payment.Core, error) {
	_, err := r.Repo.GetUserData(PaymentData)
	if err != nil {
		return payment.Core{}, err
	}
	
	idTransaksi, err := r.Repo.Insert(PaymentData)
	if err != nil {
		return payment.Core{}, err
	}

	Tagihan, err := r.Repo.GetTagihan(idTransaksi)
	if err != nil {
		return payment.Core{}, err
	}
	return Tagihan, nil
}
