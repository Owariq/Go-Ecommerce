package app

import (
	"log"
	"net"
	"os"

	"github.com/Owariq/Go-Ecommerce/product-service/internal/utils/closer"
	"github.com/Owariq/Go-Ecommerce/product-service/pkg/product_v1"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer     *grpc.Server
}

func NewApp() (*App, error) {
	a := &App{}

	err := a.initDeps()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func(){
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initDeps() error {
	inits := []func() error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer, 	
	}

	for _, init := range inits {
		if err := init(); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	return nil
}


func (a *App) initServiceProvider() error {
	a.serviceProvider = NewServiceProvider()
	return nil
}

func (a *App) initGRPCServer() error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcServer)

	product_v1.RegisterProductV1Server(a.grpcServer, a.serviceProvider.ProductImpl())

	return nil
}

func (a *App) runGRPCServer() error {
log.Printf("starting gRPC server on: %s", os.Getenv("GRPC_ADDRESS"))

list, err := net.Listen("tcp", os.Getenv("GRPC_ADDRESS"))
if err != nil {
	return err
}
err = a.grpcServer.Serve(list)
if err != nil {
	return err
}

return nil
}