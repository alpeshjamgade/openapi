package app

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"open-api/config"
	"open-api/internal/client"
	"open-api/internal/constants"
	"open-api/internal/handler"
	"open-api/internal/logger"
	"open-api/internal/repo"
	"open-api/internal/service"
	"open-api/internal/utils"
	"os"
	"os/signal"
	"syscall"
)

func Start() {

	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.WithValue(context.Background(), constants.TraceID, utils.GetUUID())
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	var Logger *zap.SugaredLogger
	if config.LogFile {
		Logger = logger.CreateFileLoggerWithCtx(ctx)
	} else {
		Logger = logger.CreateLoggerWithCtx(ctx)
	}

	Router := GetRouter()
	DbClient, Cache, HttpClient := client.GetClients(ctx)

	Repo := repo.NewRepo(DbClient, Cache, HttpClient)
	Service := service.NewService(Repo)
	Handler := handler.NewHandler(Service)
	Handler.SetupRoutes(Router)

	go func() {
		Logger.Infof("starting server on http://0.0.0.0:%s", config.HttpPort)
		http.ListenAndServe(fmt.Sprintf(":%s", config.HttpPort), Router)
	}()

	<-ctx.Done()

	Logger.Info("shutting down server")
}
