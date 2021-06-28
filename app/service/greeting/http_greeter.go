package greeting

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sergionunezgo/go-service/internal/logger"
)

// GreeterHandler struct will handle greet requests to greet back at callers.
type GreeterHandler struct {
}

// Handle will handle HTTP requests.
func (h *GreeterHandler) Handle(w http.ResponseWriter, r *http.Request) {
	greet := "Hello %s!"
	name, ok := mux.Vars(r)["name"]
	if !ok {
		logger.Log.Error("failed to extract name value")
		w.WriteHeader(http.StatusBadRequest)
		greet = "invalid parameter%s"
	}
	_, err := w.Write([]byte(fmt.Sprintf(greet, name)))
	if err != nil {
		logger.Log.Errorf("failed to write response: %s\n", err)
	}
}

// AddRoute allows to configure itself accepting a router.
func (h *GreeterHandler) AddRoute(r *mux.Router) {
	h.route(r.NewRoute().HandlerFunc(h.Handle))
}

// route receives a mux.Route to modify, like adding path, methods, etc.
func (h *GreeterHandler) route(r *mux.Route) {
	r.Path("/greet/{name}").Methods("GET")
}

// NewGreeter creates a new GreeterHandler and returns a pointer.
func NewGreeter() *GreeterHandler {
	return &GreeterHandler{}
}
