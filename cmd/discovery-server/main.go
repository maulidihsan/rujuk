package cmd

import (
	"context"
	"net"
	"os"
	"os/signal"
	"flag"
	"fmt"

	"google.golang.org/grpc"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	service "github.com/maulidihsan/rujuk/pkg/service/v1"
	api "github.com/maulidihsan/rujuk/pkg/api/v1"
	"github.com/maulidihsan/rujuk/pkg/database"
)

type Config struct {
	GRPCPort string
	HTTPPort string

	DatastoreDBHost string
	DatastoreDBUser string
	DatastoreDBPassword string
	DatastoreDBName string
}

func RunServer() error {
	ctx := context.Background()

	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBPort, "db-port", "", "Database port")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBName, "db-name", "", "Database name")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	dsn := fmt.Sprintf("%s:%s@%s:%s/%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBPort,
		cfg.DatastoreDBName)

	db, err := database.Init(dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	v1API := service.NewDiscoveryServiceServer(db.GetDB())

	listenGRPC, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	api.RegisterDiscoveryServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	fmt.Println("starting grpc server")
	return server.Serve(listenGRPC)
}