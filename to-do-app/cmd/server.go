package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"to-do-app/config"
	"to-do-app/internal/app"
	"to-do-app/internal/app/delivery/rest"
	"to-do-app/internal/pkg/datasource"

	"github.com/madevara24/go-common/server"
	ginCommon "github.com/madevara24/go-common/server/gin"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srvConfig := server.Config{
		Address:        config.Get().Port,
		Env:            config.Get().ENV,
		ReadTimeout:    time.Duration(config.Get().ServerReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Get().ServerWriteTimeout) * time.Second,
		AllowedOrigins: config.Get().AllowedOrigins,
	}

	ginHttpServer, err := ginCommon.NewGinHttpServer(srvConfig)
	if err != nil {
		panic(err)
	}

	srv := server.NewServer(ginHttpServer.GetRouter(), srvConfig)

	datasource := datasource.NewDataSource()

	container := app.NewContainer(datasource)

	router := rest.NewRouter(ctx, ginHttpServer.GetRouter(), datasource, container)

	router.RegisterRouter()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	srv.StartServer()

	<-quit
}
