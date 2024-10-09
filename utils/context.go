package utils

import (
	"context"
	"errors"
	"fmt"
	"log"

	"google.golang.org/grpc/metadata"
)

func GetUserIDFromContext(ctx context.Context) (string, error) {
	// Check gRPC metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata found")
	}

	// Extract the user_id from the metadata
	userIDs := md.Get("user_id")
	if len(userIDs) == 0 {
		return "", errors.New("user_id not found in metadata")
	}

	userID := userIDs[0]
	log.Println("UserID in service: ", userID)
	return userID, nil
}

func ExtractUserEmail(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("could not extract metadata from context")
	}
	userEmail := md["email"]
	if len(userEmail) == 0 {
		return "", fmt.Errorf("user_email not found in metadata")
	}
	return userEmail[0], nil
}

func GetInvoiceIDFromContext(ctx context.Context) (string, error) {
	// Check gRPC metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata found")
	}

	// Extract the invoice_id from the metadata
	invoiceIDs := md.Get("invoice_id")
	if len(invoiceIDs) == 0 {
		return "", errors.New("invoice_id not found in metadata")
	}

	invoiceID := invoiceIDs[0]
	log.Println("UserID in service: ", invoiceID)
	return invoiceID, nil
}
