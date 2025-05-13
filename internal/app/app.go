package app

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"open-api-client/config"
	"open-api-client/internal/client"
	"open-api-client/internal/constants"
	"open-api-client/internal/handler"
	"open-api-client/internal/logger"
	"open-api-client/internal/models"
	"open-api-client/internal/repo"
	"open-api-client/internal/service"
	"open-api-client/internal/utils"
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

	appCli(ctx, Service)

	go func() {
		Logger.Infof("starting server on http://0.0.0.0:%s", config.HttpPort)
		http.ListenAndServe(fmt.Sprintf(":%s", config.HttpPort), Router)
	}()

	<-ctx.Done()

	Logger.Info("shutting down server")
}

func appCli(ctx context.Context, service *service.Service) {
	// flags, include for all the commands.
	cli := flag.Bool("cli", false, "cli")
	createAdmin := flag.Bool("create-admin", false, "Create a new admin")
	email := flag.String("email", "", "Admin email")
	password := flag.String("password", "", "Admin password")

	flag.Parse()

	if !(*cli) {
		return
	}

	if *createAdmin {
		if *email == "" || *password == "" {
			fmt.Println("email and password are required to create admin")
			os.Exit(1)
		}

		admin := &models.Admin{Email: *email, Password: *password}
		err := service.CreateAdmin(ctx, admin)
		if err != nil {
			fmt.Printf("Failed to create admin: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Admin created successfully")
	}

	// exit cli
	os.Exit(0)

}
