package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sergionunezgo/go-service/app/service/greeting"
	"github.com/sergionunezgo/go-service/internal/logger"
	"github.com/urfave/negroni"
)

// Handler is responsible for defining a HTTP request route and corresponding handler.
type Handler interface {
	// Handle should handle HTTP requests.
	Handle(w http.ResponseWriter, r *http.Request)

	// AddRoute should allow the handler to configure itself accepting a router.
	AddRoute(r *mux.Router)
}

// Service is the struct that will hold references to all necessary data for
// running an http server.
type Service struct {
	srv *http.Server
}

// Start will begin listening on the host:port for requests.
// Blocking call.
func (s *Service) Start() error {
	logger.Log.Infof("service listening on address: %v", s.srv.Addr)
	return s.srv.ListenAndServe()
}

// Close will teardown/close any resources used by this http service.
func (s *Service) Close() error {
	logger.Log.Info("closing http service")
	return s.srv.Close()
}

// NewService will run the setup process and create a Service that can be
// used to run a http api.
func NewService(port int) *Service {
	r := mux.NewRouter().StrictSlash(true)
	api := r.PathPrefix("/api/v1").Subrouter()

	// Instantiate and setup all handlers.
	NewNotFoundHandler().AddRoute(r)

	greeting.NewGreeter().AddRoute(api)

	n := negroni.New()
	n.Use(negroni.HandlerFunc(logMiddleware))
	n.UseHandler(r)

	srv := &http.Server{
		Handler:      n,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	return &Service{
		srv: srv,
	}
}

func logMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger.Log.Infof("middleware log: %s\n", r.RequestURI)
	next(rw, r)
}
