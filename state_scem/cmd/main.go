package main

import (
	"fmt"
	"net"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"os"
	"os/signal"
	"syscall"

	"github.com/lucthienbinh/golang_scem_microservice/state_scem/endpoint"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/pb"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/repo"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/service"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/transport"
)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"microservice", "state_scem",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	if os.Getenv("RUNENV") != "docker" {
		err := godotenv.Load()
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *gorm.DB
	{
		dsn := os.Getenv("SQLITE_DSN")
		var err error
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		level.Info(logger).Log("msg", "database connected")
	}

	addRepository := repo.NewSQLRepo(db, logger)
	addService := service.NewService(addRepository, logger)
	addEndpoints := endpoint.MakeEndpoints(addService)
	grpcServer := transport.NewGRPCServer(addEndpoints, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
		pb.RegisterStateScemServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
