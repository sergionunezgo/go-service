package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sergionunezgo/gorest/app/service"
	"github.com/sergionunezgo/gorest/internal/logger"
	"github.com/sergionunezgo/gorest/internal/logger/zap"
	"github.com/urfave/cli"
)

var (
	// Reference to the api service, it has to implement io.Closer interface for clean-up.
	serviceRef service.Service
)

func main() {
	if err := zap.RegisterLog(); err != nil {
		panic("can't setup zap logger")
	}
	defer zap.CloseLog()
	logger.Log.Info("starting api service")
	if err := createApp().Run(os.Args); err != nil {
		logger.Log.Fatalf("service failed to start: %+v \n", err)
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
			Name:        "http_port",
			EnvVar:      "HTTP_PORT",
			Value:       80,
			Usage:       "port for http service",
			Destination: &config.Port,
		},
		cli.StringFlag{
			Name:        "log_level",
			EnvVar:      "LOG_LEVEL",
			Value:       "debug",
			Usage:       "Log level for the logger",
			Destination: &config.LogLevel,
		},
	}

	app.Action = func(ctx *cli.Context) error {
		logger.Log.Info("app start action")
		serviceRef = service.New(config)
		err := serviceRef.Start()
		return err
	}

	return app
}

// setupInterruptCloseHandler run a goroutine to listen for interruption signals to perform clean-up.
func setupInterruptCloseHandler() {
	channel := make(chan os.Signal, 2)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		logger.Log.Warn("service received interruption signal, clean-up and exit")
		// Call close method to perform all necessary clean-up.
		if serviceRef != nil {
			serviceRef.Close()
		}
		os.Exit(0)
	}()
}
