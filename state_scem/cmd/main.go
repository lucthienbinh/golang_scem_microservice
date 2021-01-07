package main

import (
	"database/sql"
	"fmt"
	"net"

	transport "github.com/junereycasuga/gokit-grpc-demo/transports"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"os"
	"os/signal"
	"syscall"

	"github.com/lucthienbinh/golang_scem_microservice/state_scem/endpoint"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/service"
)

const dbsource = "postgresql://postgres:postgres@postgres:5432/state_scem_database?sslmode=disable"

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

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error
		db, err = sql.Open("postgres", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

	}

	addRepository := service.NewRepo(db, logger)
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
