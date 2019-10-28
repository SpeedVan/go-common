package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SpeedVan/go-common/app"
	"github.com/SpeedVan/go-common/config"
)

// Webapp todo
type Webapp struct {
	app.App
	Router  *mux.Router
	Address string
}

// New todo
func New(config config.Config) *Webapp {
	fmt.Println("address:" + config.Get("WEBAPP_LISTEN_ADDRESS"))
	return &Webapp{
		Router:  mux.NewRouter(),
		Address: config.Get("WEBAPP_LISTEN_ADDRESS"),
	}
}

// Handle todo
func (s *Webapp) Handle(p string, h http.Handler) *Webapp {
	s.Router.Handle(p, h)
	return s
}

// HandleFunc todo
func (s *Webapp) HandleFunc(p string, f func(http.ResponseWriter, *http.Request)) *Webapp {
	s.Router.HandleFunc(p, f)
	return s
}

// HandleController todo
func (s *Webapp) HandleController(c Controller) *Webapp {
	for k, v := range c.GetRoute() {
		route := s.Router.HandleFunc(k, v.HandleFunc)
		if len(v.Methods) > 0 {
			route.Methods(v.Methods...)
		}
	}
	return s
}

// Run todo
func (s *Webapp) Run() error {
	return http.ListenAndServe(s.Address, s.Router)
}
