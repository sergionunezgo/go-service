package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/sergionunezgo/go-reuse/v2/pkg/logger"
	"github.com/sergionunezgo/go-service/app/service"
	"github.com/urfave/cli"
)

var (
	// Reference to the api service, it has to implement io.Closer interface for clean-up.
	serviceRef service.Service
)

func main() {
	log.Print("running app")
	if err := createApp().Run(os.Args); err != nil {
		log.Fatalf("app failed: %+v\n", err)
		os.Exit(1)
	}
}

// createApp loads env variables and performs setup for the service.
func createApp() *cli.App {
	setupInterruptCloseHandler()

	// May be replaced with actual config struct for the service.
	config := &service.Config{}
	app := cli.NewApp()
	app.Version = "0.0.0"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "api_port",
			EnvVar:      "API_PORT",
			Value:       80,
			Usage:       "port for the web service",
			Destination: &config.Port,
		},
		cli.StringFlag{
			Name:        "log_level",
			EnvVar:      "LOG_LEVEL",
			Value:       "debug",
			Usage:       "log level for the logger",
			Destination: &config.LogLevel,
		},
	}

	app.Action = func(ctx *cli.Context) error {
		err := logger.UseZapLogger(config.LogLevel)
		if err != nil {
			return errors.Wrap(err, "logger UseZapLogger")
		}

		logger.Log.Info("starting service")
		serviceRef = service.New(config)
		err = serviceRef.Start()
		return err
	}

	return app
}

// setupInterruptCloseHandler run a goroutine to listen for interruption signals to perform clean-up.
func setupInterruptCloseHandler() {
	interruptions := make(chan os.Signal, 2)
	signal.Notify(interruptions, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interruptions
		logger.Log.Warn("interruption signal received, starting clean-up")
		if serviceRef != nil {
			serviceRef.Close()
		}
		logger.CloseLogger()
		os.Exit(0)
	}()
}
