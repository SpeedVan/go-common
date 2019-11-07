package web

import (
	"net/http"

	"github.com/SpeedVan/go-common/log"
	"github.com/SpeedVan/go-common/log/common"

	"github.com/gorilla/mux"

	"github.com/SpeedVan/go-common/app"
	"github.com/SpeedVan/go-common/app/web/handler"
	"github.com/SpeedVan/go-common/config"
)

// Webapp todo
type Webapp struct {
	Logger log.Logger
	app.App
	Router  *mux.Router
	Address string
}

// New todo
func New(config config.Config, logger log.Logger) *Webapp {
	if logger == nil {
		logger = common.NewCommon(log.Debug)
	}

	return &Webapp{
		Logger:  logger,
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
		s.Logger.DebugF("registed request path:%v %v", k, v)
		s.Router.Handle(k, &handler.DebugHandler{Logger: s.Logger, OrginalHandler: v})
	}
	return s
}

// Run todo
func (s *Webapp) Run(level log.Level) error {
	s.Logger.InfoF("start with address: %v", s.Address)
	s.Logger.SetLevel(level)
	return http.ListenAndServe(s.Address, s.Router)
}

// SimpleRun todo
func (s *Webapp) SimpleRun() error {
	return s.Run(log.Info)
}
