package main

import (
	"net/http"
	"time"

	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/log"
	"github.com/SpeedVan/go-common/config/mock"
)

func main() {
	logger := log.NewCommon(log.Debug) // this level control webapp init log level

	m := map[string]string{
		"WEBAPP_LISTEN_ADDRESS": ":9999",
	}
	cfg := mock.New(m)

	app := web.New(cfg, logger)
	app.HandleController(NewTestController(logger))
	app.Run(log.Debug) // this level control webapp runtime log level
}

// TestController todo
type TestController struct {
	web.Controller
	logger log.Logger
}

// NewTestController todo
func NewTestController(logger log.Logger) *TestController {
	return &TestController{
		logger: logger,
	}
}

// GetRoute todo
func (s *TestController) GetRoute() web.RouteMap {
	items := []*web.RouteItem{
		&web.RouteItem{Path: "/{_dummy:.*}", HandleFunc: s.Call},
	}

	return web.NewRouteMap(items...)
}

// Call todo
func (s *TestController) Call(w http.ResponseWriter, r *http.Request) {
	s.logger.DebugF("%v, %v", r.Method, r.URL.Path)
	time.Sleep(time.Duration(20) * time.Second)
	s.logger.Debug("after sleep 20s")
}
