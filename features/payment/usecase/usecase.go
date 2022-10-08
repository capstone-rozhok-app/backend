package usecase

import (
	"rozhok/features/payment"
	"rozhok/utils/helper"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Payment struct {
	Repo payment.PaymentData
}

func New(repo payment.PaymentData) *Payment {
	return &Payment{
		Repo: repo,
	}
}

func (r *Payment) Create(PaymentData payment.Core) (payment.Core, error) {

	// get user data to check that user has inserted alamat with status utama
	_, err := r.Repo.GetUserData(PaymentData)
	if err != nil {
		return payment.Core{}, err
	}

	// get grandtotal
	PaymentData.GrandTotal, err = r.Repo.GetGrandTotal(PaymentData)
	if err != nil {
		return payment.Core{}, err
	}

	// create payment midtrans
	PaymentData.KodeTransaksi = helper.GenerateTF(int(PaymentData.Client.ID))
	midtransCore := payment.ToMidtransCore(PaymentData)
	switch PaymentData.Bank {
	case "bca":
		midtransCore.PaymentType = coreapi.PaymentTypeBankTransfer
		midtransCore.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		}
	case "bni":
		midtransCore.PaymentType = coreapi.PaymentTypeBankTransfer
		midtransCore.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBni,
		}
	case "bri":
		midtransCore.PaymentType = coreapi.PaymentTypeBankTransfer
		midtransCore.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBri,
		}
	case "permata":
		midtransCore.PaymentType = coreapi.PaymentTypeBankTransfer
		midtransCore.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankPermata,
		}
	}

	midtransInvoice, errChargeMidtrans := coreapi.ChargeTransaction(midtransCore)
	if errChargeMidtrans != nil {
		return payment.Core{}, err
	}

	// update stok produk
	errUpdateStok := r.Repo.UpdateStokProduct(PaymentData)
	if errUpdateStok != nil {
		return payment.Core{}, errUpdateStok
	}

	// create tagihan
	if PaymentData.Bank != "permata" {
		PaymentData.NoVA = midtransInvoice.VaNumbers[0].VANumber
	} else {
		PaymentData.NoVA = midtransInvoice.PermataVaNumber
	}
	PaymentData.TipePembayaran = midtransInvoice.PaymentType
	PaymentData.IdTagihan, err = r.Repo.InsertTagihan(PaymentData)
	if err != nil {
		return payment.Core{}, err
	}

	// create transaksi
	err = r.Repo.InsertTransaksi(PaymentData)
	if err != nil {
		return payment.Core{}, err
	}

	return PaymentData, nil
}

func (r *Payment) PaymentWebHook(OrderID, status string) error {
	var PaymentCore payment.Core

	PaymentCore.KodeTransaksi = OrderID
	switch status {
	case "settlement":
		PaymentCore.StatusTransaksi = "dibayar"
	case "pending":
		PaymentCore.StatusTransaksi = "belum_dibayar"
	default:
		PaymentCore.StatusTransaksi = "dibatalkan"
	}

	err := r.Repo.UpdateTransaksi(PaymentCore)
	if err != nil {
		return err
	}

	return nil
}
