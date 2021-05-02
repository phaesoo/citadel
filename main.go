package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/phaesoo/keybox/apps/keybox"
	"github.com/phaesoo/keybox/configs"
	gw "github.com/phaesoo/keybox/gen/gw/proto"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

type App interface {
	Listen() error
	Shutdown()
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := gw.RegisterAdminHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":10080", mux)
}

func runApp(app App, onDone func()) {
	done := make(chan struct{})
	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdown

		app.Shutdown()
		close(done)
	}()

	if err := app.Listen(); err != nil {
		panic(err)
	}

	<-done
	onDone()
}

func main() {
	log.Print("Run")
	flag.Parse()
	defer glog.Flush()

	// if err := run(); err != nil {
	// 	glog.Fatal(err)
	// }

	wg := sync.WaitGroup{}

	app := keybox.NewApp(configs.Get())

	wg.Add(1)
	go runApp(app, wg.Done)

	log.Print("Wait")
	wg.Wait()
	log.Print("Finished")

	log.Print("End")
}
