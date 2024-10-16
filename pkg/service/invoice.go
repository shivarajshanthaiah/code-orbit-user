package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/razorpay/razorpay-go"
	"github.com/shivaraj-shanthaiah/code_orbit_user/config"
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

	subscriptionReq := &adminpb.SubscriptionID{
		ID: req.SubsriptionId,
	}

	subscription, err := u.AdminClient.GetSubscriptionByID(ctx, subscriptionReq)
	if err != nil {
		return nil, err
	}

	// Create invoice
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

	userdetails, err := u.Repo.FindUserByID(string(userID))
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error finding user details",
			Payload: &pb.Response_Error{
				Error: err.Error(),
			},
		}, err
	}

	// Prepare email message
	message := fmt.Sprintf(
		"Your Invoice Details:\nName : %s\nInvoiceID: %d\nPlan: %s\nDuration: %s\nPrice: %.2f\nGST: %.2f\nTotal Price: %d",
		userdetails.UserName, invoice.ID, subscription.Plan, subscription.Duration,
		subscription.Price, subscription.Gst, subscription.TotalPrice,
	)

	subject := "Pending Invoice Details"

	// Send email notification via RabbitMQ
	if err := rabbitmq.SendMail(userEmail, message, subject); err != nil {
		return nil, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Invoice generated successfully",
	}, nil
}

func (u *UserService) MakePaymentService(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {

	config := config.LoadConfig()

	invoice, err := u.Repo.FindInvoiceByID(uint(req.InvoiceId))
	if err != nil {
		return nil, err
	}

	if invoice.PaymentStatus == "PAID" {
		return nil, errors.New("this invoice is already paid")
	}

	amountInPaisa := invoice.Amount * 100
	razorpayClient := razorpay.NewClient(config.RAZORPAYKEYID, config.RAZORPAYSECRETKEY)

	orderData := map[string]interface{}{
		"amount":   amountInPaisa,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}

	// Create a Razorpay order
	order, err := razorpayClient.Order.Create(orderData, nil)
	if err != nil {
		log.Println("Error in Razorpay order creation:", err.Error())
		log.Println("Response from Razorpay:", order)
		return nil, errors.New("failed to create Razorpay order")
	}

	// Extract the order ID from Razorpay's response
	orderID, ok := order["id"].(string)
	if !ok {
		return nil, errors.New("failed to extract order ID from Razorpay response")
	}

	response := &pb.PaymentResponse{
		User_ID:    invoice.UserID,
		Invoice_ID: uint32(invoice.ID),
		Amount:     uint32(invoice.Amount),
		Plan:       invoice.Plan,
		Order_ID:   orderID,
	}
	return response, nil
}

func (u *UserService) PaymentSuccessService(ctx context.Context, req *pb.ConfirmRequest) (*pb.ConfirmResponse, error) {

	invoiceID, err := strconv.Atoi(req.Invoice_ID)
	if err != nil {
		return nil, errors.New("invalid invoice ID format")
	}

	paidAmount, err := strconv.Atoi(req.Amount)
	if err != nil {
		return nil, errors.New("invalid Amount format")
	}

	invoice, err := u.Repo.FindInvoiceByID(uint(invoiceID))
	if err != nil {
		return nil, err
	}

	invoice.PaymentStatus = "PAID"
	invoice.PaymentID = req.Payment_ID

	if err := u.Repo.UpdateInvoice(invoice); err != nil {
		return nil, err
	}

	// Create a new record in the RazorPay table
	razorpayRecord := &model.RazorPay{
		InvoiceID:       uint32(invoiceID),
		RazorPaymentID:  req.Payment_ID,
		RazorPayOrderID: req.Order_ID,
		Signature:       req.Signature,
		AmountPaid:      uint32(paidAmount),
	}
	if err := u.Repo.CreateRazorPayment(razorpayRecord); err != nil {
		return nil, err
	}

	user, err := u.Repo.FindUserByID(invoice.UserID)
	if err != nil {
		return nil, err
	}

	user.IsPrimeMember = true
	currentTime := time.Now()

	// If the current membership expiry date is in the future, extend from that date
	if user.MembershipExpiryDate.After(currentTime) {
		currentTime = user.MembershipExpiryDate
	}

	// Extend membership based on the selected plan
	switch invoice.Plan {
	case "Basic":
		user.MembershipExpiryDate = currentTime.AddDate(0, 1, 0) // 1 month
	case "Standard":
		user.MembershipExpiryDate = currentTime.AddDate(0, 3, 0) // 3 months
	case "Premium":
		user.MembershipExpiryDate = currentTime.AddDate(0, 6, 0) // 6 months
	default:
		return nil, errors.New("invalid plan type")
	}

	if err := u.Repo.UpdateUser(user); err != nil {
		return nil, err
	}

	message := fmt.Sprintf("Your Invoice Details: \nName : %s\nInvoiceID: %d\nPayment ID: %s\nTotal Amount Paid: %d\nCurrent Plan: %s\nPrime Member: %t\nMembership Till %s",
		user.UserName, invoiceID, invoice.PaymentID, invoice.Amount, invoice.Plan, user.IsPrimeMember, user.MembershipExpiryDate)

	subject := "Payment Confirmation Notification"

	if err := rabbitmq.SendMail(user.Email, message, subject); err != nil {
		return nil, err
	}

	return &pb.ConfirmResponse{
		Status:     "true",
		Invoice_ID: req.Invoice_ID,
	}, nil
}
