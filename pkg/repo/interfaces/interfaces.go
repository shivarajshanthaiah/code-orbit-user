package interfaces

import "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"

type UserRepoInter interface {
	CreateUser(user *model.User) (string, error)
	FindUserByEmail(email string) (*model.User, error)
	FindUserByPhone(phone string) (*model.User, error)
	FindUserByID(userID string) (*model.User, error)
	UpdateUser(user *model.User) error
	FindAllUsers() (*[]model.User, error)

	CreateInvoice(invoice *model.Invoice) error
	FindInvoiceByID(invoiceID uint) (*model.Invoice, error)
	UpdateInvoice(invoice *model.Invoice) error
	CreateRazorPayment(payment *model.RazorPay) error

	FindAllPrimeUsers() ([]model.User, error)
	UpdateUserMembership(user model.User) error
}
