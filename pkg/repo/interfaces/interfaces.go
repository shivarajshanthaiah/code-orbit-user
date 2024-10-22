package interfaces

import (
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"
	"gorm.io/gorm"
)

type UserRepoInter interface {
	CreateUser(user *model.User) (string, error)
	FindUserByEmail(email string) (*model.User, error)
	FindUserByPhone(phone string) (*model.User, error)
	FindUserByID(userID string) (*model.User, error)
	UpdateUser(user *model.User) error
	FindAllUsers() (*[]model.User, error)

	CreateSubscription(plan *model.Subscription) error
	GetAllPlans() (*[]model.Subscription, error)
	GetPlanByID(planID uint) (*model.Subscription, error)
	UpdatePlan(plan *model.Subscription) error

	CreateInvoice(invoice *model.Invoice) error
	FindInvoiceByID(invoiceID uint) (*model.Invoice, error)
	UpdateInvoice(invoice *model.Invoice) error
	CreateRazorPayment(payment *model.RazorPay) error

	FindAllPrimeUsers() ([]model.User, error)
	UpdateUserMembership(user model.User) error

	FindAllUsersCount() *gorm.DB
	FindUsersByCondition(condition string) *gorm.DB
	CountByCondition(condition string) int32
	SumAmountsByCondition(condition string) uint32
}
