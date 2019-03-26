package main

import (
	"net/http"
	"os"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/maulidihsan/rujuk/pkg/database"
	"github.com/maulidihsan/rujuk/cmd/rest-client/routers"
)

type Config struct {
	DiscoveryURI string
	HTTPPort string

	DatastoreDBHost string
	DatastoreDBUser string
	DatastoreDBPassword string
	DatastoreDBName string
	DatastoreDBPort string
}

func RunServer() error {

	var cfg Config
	flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP port to bind")
	flag.StringVar(&cfg.DatastoreDBPort, "db-port", "", "Database port")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBName, "db-name", "", "Database name")
	flag.StringVar(&cfg.DiscoveryURI, "uri", "", "Discovery Server URI")
	flag.Parse()

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP server: '%s'", cfg.HTTPPort)
	}

	if len(cfg.DiscoveryURI) == 0 {
		return fmt.Errorf("invalid URI: '%s'", cfg.DiscoveryURI)
	}

	if len(cfg.DatastoreDBPort) == 0 {
		cfg.DatastoreDBPort = "3306"
	}
	if len(cfg.DatastoreDBHost) == 0 {
		cfg.DatastoreDBHost = "127.0.0.1"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBPort,
		cfg.DatastoreDBName)

	conn, err := database.Init(dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer conn.DB.Close()
	rest, err := routers.Init(conn.GetDB(), cfg.DiscoveryURI)
	if err != nil {
		return fmt.Errorf("failed to connect to discovery server: %v", err)
	}

	r := gin.Default()
	r.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "hello"})
	})

	r.GET("/register", rest.Register)
	r.GET("/pasien", rest.GetMyPasiens)
	r.GET("/room", rest.GetMyRooms)
	r.POST("/pasien", rest.AddPasien)
	r.POST("/room", rest.AddRoom)
	r.GET("/network", rest.ListOtherRS)
	r.GET("/network/rooms", rest.GetRoomList)
	r.POST("/network/transfer", rest.TransferPasien)

	r.Run(":"+cfg.HTTPPort)
	return nil
}

func main() {
	if err := RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}