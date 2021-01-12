package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/vardius/gorouter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/endpoint"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/pb"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/repo"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/service"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/transport"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	customFunc   grpc_recovery.RecoveryHandlerFunc
	loggerGlobal log.Logger
)

func main() {
	// --------------- Define Log template ---------------
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
	loggerGlobal = logger
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	if os.Getenv("RUNENV") != "docker" {
		err := godotenv.Load()
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	// --------------- Connnect database ---------------
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

	// --------------- Migrate database ---------------
	initRepo := repo.NewSQLInitRepo(db, logger)
	// BECAREFUL OF THESE LINES WILL WIPE OUT ALL OF YOUR DATA
	if err := initRepo.DeleteDatabase(context.Background()); err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(1)
	}
	if err := initRepo.MigrationDatabase(context.Background()); err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(1)
	}
	level.Info(logger).Log("msg", "database migrated")

	// --------------- Create new implement instance ---------------
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "scem_group",
		Subsystem: "state_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "scem_group",
		Subsystem: "state_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	addRepository := repo.NewSQLRepo(db, logger)
	addService := service.NewService(addRepository, logger)
	addEndpoints := endpoint.MakeEndpoints(addService, logger, requestCount, requestLatency)
	grpcServer := transport.NewGRPCServer(addEndpoints, logger)

	// --------------- Listen to kill signal ---------------
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// --------------- Define GRPC server ---------------
	grpcListener, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(1)
	}

	// --------------- Handle GRPC panic ---------------
	// Define customfunc to handle panic
	customFunc = func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(customFunc),
	}

	// Register GRPC server and implement instance
	go func() {
		baseServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
		)
		pb.RegisterStateScemServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server GRPC started successfully 🚀 at port 9001")
		baseServer.Serve(grpcListener)
	}()

	// --------------- Define HTTP server ---------------
	go func() {
		router := gorouter.New(recoverMiddleware)
		router.GET("/metrics", promhttp.Handler())
		level.Info(logger).Log("msg", "Server HTTP started successfully 🚁 at port 9002")
		err2 := http.ListenAndServe(os.Getenv("HTTP_PORT"), router)
		if err2 != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(1)
		}
	}()

	level.Error(logger).Log("exit", <-errs)
}

func recoverMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rcv := recover(); rcv != nil {
				level.Error(loggerGlobal).Log("err", "Internal Server Error")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
