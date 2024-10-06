package service

import (
	"context"
	"fmt"
	"log"

	adminpb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/clients/admin/adminpb"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/clients/rabbitmq"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"
	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
	"github.com/shivaraj-shanthaiah/code_orbit_user/utils"
)

func (u *UserService) GenerateInvoiceService(ctx context.Context, req *pb.InvoiceRequest) (*pb.Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "User not authenticated(id)",
			Payload: &pb.Response_Error{
				Error: "User authentication failed(id)",
			},
		}, err
	}

	user, err := u.Repo.FindUserByID(userID)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "User not found(id)",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}
	log.Println("User found:", user.UserName)

	userEmail, err := utils.ExtractUserEmail(ctx)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "User not found(email)",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	// subscription, err := u.AdminClient.GetSubscriptionByID(req.SubsriptionId)
	// if err != nil {
	// 	return nil, err
	// }

	subscriptionReq := &adminpb.SubscriptionID{
		ID: req.SubsriptionId, // Assuming SubsriptionId is the correct field
	}

	// Call the GetSubscriptionByID method with context, request, and options
	subscription, err := u.AdminClient.GetSubscriptionByID(ctx, subscriptionReq)
	if err != nil {
		return nil, err
	}

	// Create invoic
	invoice := model.Invoice{
		UserID:         req.UserId,
		UserEmail:      req.UserEmail,
		SubscriptionID: int(subscription.ID),
		Plan:           subscription.Plan,
		Amount:         subscription.TotalPrice,
	}

	err = u.Repo.CreateInvoice(&invoice)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error creating invoice in database",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	// Prepare email message
	message := "Invoice Details:\n" +
		"Plan: " + subscription.Plan + "\n" +
		"Duration: " + subscription.Duration + "\n" +
		"Price: " + fmt.Sprintf("%.2f", subscription.Price) + "\n" +
		"GST: " + fmt.Sprintf("%.2f", subscription.Gst) + "\n" +
		"Total Price: " + fmt.Sprintf("%.2f", subscription.TotalPrice)

	subject := "Your Pending Invoice Details"

	// Send email notification via RabbitMQ
	if err := rabbitmq.SendMail(userEmail, message, subject); err != nil {
		return nil, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Invoice generated successfully",
		Payload: &pb.Response_Data{
			Data: string(rune(invoice.ID)),
		},
	}, nil
}
