package webserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sergionunezgo/goservice/internal/logger"
	"github.com/urfave/negroni"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Testing")
}

type logFunc func()

func (lf logFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	lf()
	next(rw, r)
}

// HTTPService is the struct that will hold references to all necessary data for
// running an http server.
type HTTPService struct {
	srv *http.Server
}

// NewHTTPService will run the setup process and create a HTTPService that can be
// used to run a http api.
func NewHTTPService(port int) *HTTPService {
	r := mux.NewRouter().StrictSlash(true)
	setupRoutes(r)

	n := negroni.New(logFunc(func() {
		logger.Log.Info("middleware log")
	}))
	n.UseHandler(r)

	srv := &http.Server{
		Handler:      n,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	return &HTTPService{
		srv: srv,
	}
}

// Start will begin listening on the host:port for requests.
// Blocking call.
func (as *HTTPService) Start() error {
	logger.Log.Infof("service listening on address: %v", as.srv.Addr)
	return as.srv.ListenAndServe()
}

// Close will teardown/close any resources used by this http service.
func (as *HTTPService) Close() error {
	logger.Log.Info("ApiService Close method executed")
	return nil
}

func setupRoutes(r *mux.Router) {
	r.HandleFunc("/", handler)
}
