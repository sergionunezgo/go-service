package webserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sergionunezgo/gorest/internal/logger"
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

type HttpService struct {
	srv *http.Server
}

func NewHttpService(host string, port string) *HttpService {
	r := mux.NewRouter().StrictSlash(true)
	setupRoutes(r)

	n := negroni.New(logFunc(func() {
		logger.Log.Info("middleware log")
	}))
	n.UseHandler(r)

	srv := &http.Server{
		Handler:      n,
		Addr:         fmt.Sprintf("%s:%s", host, port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	return &HttpService{
		srv: srv,
	}
}

func (as *HttpService) Start() error {
	logger.Log.Infof("service listening on address: %v", as.srv.Addr)
	return as.srv.ListenAndServe()
}

func (as *HttpService) Close() error {
	logger.Log.Info("ApiService Close method executed")
	return nil
}

func setupRoutes(r *mux.Router) {
	r.HandleFunc("/", handler)
}
