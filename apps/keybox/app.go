package keybox

import (
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/phaesoo/keybox/configs"
	handler "github.com/phaesoo/keybox/gen/go/proto"
	"github.com/phaesoo/keybox/internal/repo"
	"github.com/phaesoo/keybox/internal/services/admin"
	"github.com/phaesoo/keybox/internal/services/encrypt"
	"github.com/phaesoo/keybox/pkg/db"
	"github.com/phaesoo/keybox/pkg/memdb"
	"google.golang.org/grpc"
)

type App struct {
	handler.AdminServer
	handler.EncryptServer

	grpcServer *grpc.Server
	config     configs.Config
}

func NewApp(config configs.Config) *App {
	app := App{
		grpcServer: grpc.NewServer(
			grpc_middleware.WithUnaryServerChain(
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
		config: config,
	}
	app.setupServices()
	return &app
}

func (app *App) setupServices() {
	connString, err := app.config.Mysql.ConnString()
	if err != nil {
		panic(err)
	}

	db, err := db.NewDB("mysql", connString)
	if err != nil {
		panic(err)
	}

	rc := app.config.Redis
	memdbConf := memdb.Config{
		Address:      rc.Address(),
		DB:           rc.Database,
		TLSRequired:  rc.TLSRequired,
		AuthRequired: rc.AuthRequired,
		Password:     rc.Password,
		CACert:       rc.CACert,
	}
	pool := memdb.NewPool(memdbConf)

	repo := repo.NewRepo(db, pool, app.config.App.SecretKey)
	adminService := admin.NewService(repo)
	adminServer := admin.NewServer(adminService)
	handler.RegisterAdminServer(app.grpcServer, adminServer)

	encryptService := encrypt.NewService(repo)
	encryptServer := encrypt.NewServer(encryptService)
	handler.RegisterEncryptServer(app.grpcServer, encryptServer)
}

// Listen starts server on the address
func (app *App) Listen() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", app.config.App.Host, app.config.App.Port))
	if err != nil {
		return err
	}
	return app.grpcServer.Serve(lis)
}

// Shutdown server
func (app *App) Shutdown() {
	app.grpcServer.GracefulStop()
}
