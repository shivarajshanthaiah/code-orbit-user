package di

import (
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_user/config"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/clients/admin"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/clients/problem"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/cron"
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

	problemClient, err := problem.ClientDial(*cnfg)
	if err != nil {
		log.Fatal("failed to connect to problem client")
	}

	adminClient, err := admin.ClientDial(*cnfg)
	if err != nil {
		log.Fatal("failed to connect to admin client")
	}

	userRepo := repo.NewUserRepository(db)

	ci := cron.NewCron(userRepo)
	if err := ci.SheduleJob(); err != nil {
		log.Fatalf("Failed to schedule cron job: %v", err)
	}
	ci.InitCron()

	userService := service.NewUserService(userRepo, redis, twilio, problemClient, adminClient)

	userHandler := handler.NewUserHandler(userService)

	err = server.NewGrpcUserServer(cnfg.GrpcPort, userHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}
}
