package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
)

func (u *UserHandler) GenerateInvoice(ctx context.Context, p *pb.InvoiceRequest) (*pb.Response, error) {
	response, err := u.SVC.GenerateInvoiceService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) MakePayment(ctx context.Context, p *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	response, err := u.SVC.MakePaymentService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) PaymentSuccess(ctx context.Context, p *pb.ConfirmRequest) (*pb.ConfirmResponse, error) {
	response, err := u.SVC.PaymentSuccessService(ctx, p)
	if err != nil {
		return response, err
	}
	return nil, err
}
