package web

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/SpeedVan/go-common/log"

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

	doneChan chan error
	server   *http.Server
	ctx      context.Context
}

// New todo
func New(config config.Config, logger log.Logger) *Webapp {
	if logger == nil {
		logger = log.NewCommon(log.Debug)
	}

	doneChan := make(chan error)
	osSignalChan := make(chan os.Signal, 1)
	osKillChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(osKillChan, syscall.SIGKILL)

	ctx := context.Background()

	addr := config.Get("WEBAPP_LISTEN_ADDRESS")
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    addr,
		Handler: router,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		sig := <-osSignalChan
		logger.DebugF("received sig: %v", sig)
		// Shutdown会让监听断开，即协程里的server.ListenAndServe()将往后执行。
		// Shutdown按协议说的是graceful，Close是immediately（强杀）。
		if err := server.Shutdown(ctx); err == nil || err == http.ErrServerClosed {
			logger.Debug("Shutdown ok")
			doneChan <- nil
		} else {
			logger.ErrorF("Shutdown err: %v", err)
			doneChan <- err
		}
	}()

	go func() {
		sig := <-osKillChan
		logger.DebugF("received sig: %v", sig)
		if err := server.Close(); err == nil || err == http.ErrServerClosed {
			logger.Debug("Close ok")
			doneChan <- nil
		} else {
			logger.ErrorF("Close err: %v", err)
			doneChan <- err
		}
	}()

	return &Webapp{
		Logger:  logger,
		Router:  router,
		Address: addr,

		doneChan: doneChan,
		server:   server,
		ctx:      ctx,
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
		s.Logger.DebugF("registed request path: %v %v", k, v)
		s.Router.Handle(k, &handler.DebugHandler{Logger: s.Logger, OrginalHandler: v})
	}
	return s
}

// RegisterOnShutdown todo
func (s *Webapp) RegisterOnShutdown(f func()) {
	s.server.RegisterOnShutdown(f)
}

// Run todo
func (s *Webapp) Run(level log.Level) error {
	s.Logger.InfoF("start with address: %v", s.Address)
	s.Logger.SetLevel(level)

	// http 端口监听
	go func() {
		if err := s.server.ListenAndServe(); err == nil || err == http.ErrServerClosed {
			s.Logger.Debug("Listen and serve close ok")
		} else {
			s.Logger.ErrorF("Listen and serve close err: %v", err)
		}
	}()

	// 可以考虑其他端口监听
	//

	// 退出信号
	err := <-s.doneChan
	if err != nil {
		s.Logger.ErrorF("server shutdown err:", err.Error())
	} else {
		s.Logger.InfoF("server shutdown graceful")
	}
	return err
}

// SimpleRun todo
func (s *Webapp) SimpleRun() error {
	return s.Run(log.Info)
}
