package main

import (
	"fmt"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"os"
	"os/signal"
	"syscall"

	"github.com/lucthienbinh/golang_scem_microservice/state_scem/endpoint"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/repo"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/service"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/transport"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
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

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		// pb.RegisterMathServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
