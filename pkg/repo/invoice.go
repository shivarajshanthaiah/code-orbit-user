package repo

import (
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"
)

func (u *UserRepository) CreateInvoice(invoice *model.Invoice) error {
	if err := u.DB.Create(&invoice).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) FindInvoiceByID(invoiceID uint) (*model.Invoice, error) {
	var invoice model.Invoice
	if err := u.DB.Where("id = ?", invoiceID).First(&invoice).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (u *UserRepository) UpdateInvoice(invoice *model.Invoice) error {
	if err := u.DB.Save(&invoice).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) CreateRazorPayment(payment *model.RazorPay) error {
	if err := u.DB.Create(payment).Error; err != nil {
		return err
	}
	return nil
}
