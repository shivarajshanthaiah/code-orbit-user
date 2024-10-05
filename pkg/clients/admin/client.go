package admin

import (
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_user/config"
	pb "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/clients/admin/adminpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.AdminServiceClient, error) {
	grpc, err := grpc.Dial(":"+cfg.AdminPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc admin client: %s, ", cfg.AdminPort)
		return nil, err
	}
	log.Printf("Succesfully connected to admin client at port: %v", cfg.AdminPort)
	return pb.NewAdminServiceClient(grpc), nil
}
