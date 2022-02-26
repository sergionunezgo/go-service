package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sergionunezgo/go-reuse/v2/pkg/logger"
)

// NotFoundHandler struct configures and responds to all invalid paths.
type NotFoundHandler struct {
}

// Handle will handle HTTP requests.
func (h *NotFoundHandler) Handle(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("matched invalid path")
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("not found!!"))
	if err != nil {
		logger.Log.Errorf("failed to write response: %+v\n", err)
	}
}

// AddRoute allows to configure itself accepting a router.
func (h *NotFoundHandler) AddRoute(r *mux.Router) {
	h.route(r.NewRoute().HandlerFunc(h.Handle))
}

// route receives a mux.Route to modify, like adding path, methods, etc.
func (h *NotFoundHandler) route(r *mux.Route) {
	r.PathPrefix("/")
}

// NewNotFoundHandler creates a new NotFoundHandler and returns a pointer.
func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}
