package repo

import "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"

func (u *UserRepository) CreateInvoice(invoice *model.Invoice) error {
	if err := u.DB.Create(&invoice).Error; err != nil {
		return err
	}
	return nil
}
