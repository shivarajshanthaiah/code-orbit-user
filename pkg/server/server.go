package server

import (
	"fmt"
	"log"
	"net"

	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/handler"
	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/proto"
	"google.golang.org/grpc"
)

func NewGrpcUserServer(port string, handler *handler.UserHandler) error {

	log.Println("connecting to gRPC server")

	addr := fmt.Sprintf(":%s", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("error creating listener on %v", addr)
		return err
	}

	grpc := grpc.NewServer()
	pb.RegisterUserServiceServer(grpc, handler)

	log.Printf("listening on gRPC server %v", port)
	err = grpc.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}