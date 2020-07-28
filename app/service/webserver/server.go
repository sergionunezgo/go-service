package webserver

import (
	"net/http"

	"github.com/sergionunezgo/gorest/internal/logger"
)

type ApiService struct {
	srv *http.Server
}

func New(srv *http.Server) *ApiService {
	return &ApiService{
		srv: srv,
	}
}

func (as *ApiService) Start() error {
	logger.Log.Infof("service listening on address: %v", as.srv.Addr)
	return as.srv.ListenAndServe()
}

func (as *ApiService) Close() error {
	logger.Log.Info("ApiService Close method executed")
	return nil
}
