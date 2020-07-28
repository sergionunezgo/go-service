package service

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sergionunezgo/gorest/app/service/webserver"
	"github.com/sergionunezgo/gorest/internal/logger"
	"github.com/urfave/negroni"
)

type Config struct {
	Port     int
	LogLevel string
}

type Service interface {
	Start() error
	Close() error
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Testing")
}

type logFunc func()

func (lf logFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	lf()
	next(rw, r)
}

func New(cfg *Config) Service {
	logger.Log.Info("initializing api service")
	r := mux.NewRouter().StrictSlash(true)
	setupRoutes(r)

	n := negroni.New(logFunc(func() {
		logger.Log.Info("middleware log")
	}))
	n.UseHandler(r)

	sPort := strconv.Itoa(cfg.Port)
	srv := &http.Server{
		Handler: n,
		Addr:    "127.0.0.1:" + sPort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	return webserver.New(srv)
}

func setupRoutes(r *mux.Router) {
	r.HandleFunc("/", handler)
}
