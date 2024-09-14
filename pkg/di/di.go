package di

import (
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_user/config"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/db"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/handler"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/repo"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/server"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/service"
)

func Init() {
	cnfg := config.LoadConfig()

	redis, err := config.SetupRedis(cnfg)
	if err != nil {
		log.Fatalf("Failed to connect to redis")
	}

	twilio := config.SetupTwilio(cnfg)
	db := db.ConnectDB(cnfg)

	userRepo := repo.NewUserRepository(db)

	userService := service.NewUserService(userRepo, redis, twilio)

	userHandler := handler.NewUserHandler(userService)

	err = server.NewGrpcUserServer(cnfg.GrpcPort, userHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}
}
