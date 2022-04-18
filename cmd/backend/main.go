package main

import (
	"flag"
	"github.com/KittyBot-Org/KittyBotGo/internal/backend"
	"github.com/KittyBot-Org/KittyBotGo/internal/modules"
	"github.com/KittyBot-Org/KittyBotGo/internal/routes"
	"os"
	"os/signal"
	"syscall"

	"github.com/KittyBot-Org/KittyBotGo/internal/config"
	"github.com/KittyBot-Org/KittyBotGo/internal/db"
	"github.com/disgoorg/log"
)

var (
	shouldSyncDBTables *bool
	exitAfterSync      *bool
	version            = "dev"
)

func init() {
	shouldSyncDBTables = flag.Bool("sync-db", false, "Whether to sync the database tables")
	exitAfterSync = flag.Bool("exit-after-sync", false, "Whether to exit after syncing commands and database tables")
	flag.Parse()
}

func main() {
	var err error
	logger := log.New(log.Ldate | log.Ltime | log.Lshortfile)

	logger.Infof("Starting b version: %s", version)
	logger.Infof("Syncing DB tables? %v", *shouldSyncDBTables)
	logger.Infof("Exiting after syncing? %v", *exitAfterSync)

	var cfg backend.Config
	if err = config.LoadConfig(&cfg); err != nil {
		logger.Fatal("Failed to load config: ", err)
	}
	logger.SetLevel(cfg.LogLevel)

	b := &backend.Backend{
		Logger:  logger,
		Version: version,
	}

	if b.DB, err = db.SetupDatabase(b.Config.Database); err != nil {
		b.Logger.Fatal("Failed to setup database: ", err)
	}
	defer b.DB.Close()

	if *exitAfterSync {
		b.Logger.Infof("Exiting after syncing database tables")
		os.Exit(0)
	}

	b.LoadCommands(modules.Modules)
	b.SetupRestServices()
	if err = b.SetupPrometheusAPI(); err != nil {
		b.Logger.Fatal("Failed to setup prometheus api: ", err)
	}
	if err = b.SetupScheduler(); err != nil {
		b.Logger.Fatal("Failed to setup scheduler: ", err)
	}
	defer b.Scheduler.Shutdown()
	b.SetupServer(routes.Handler(b))

	b.Logger.Info("Backend is running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
