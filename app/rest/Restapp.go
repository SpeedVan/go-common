package rest

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gorilla/mux"

	"github.com/alpha-ss/go-common/app"
	"github.com/alpha-ss/go-common/app/rest/handler"
	"github.com/alpha-ss/go-common/config"
	"github.com/alpha-ss/go-common/log"
)

// Restapp todo
type Restapp struct {
	Logger log.Logger
	app.App
	Router  *mux.Router
	Address string

	doneChan chan error
	server   *http.Server
	ctx      context.Context

	onSignalFunc []func(os.Signal)
}

// New todo
func New(config config.Config, logger log.Logger) *Restapp {
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
	// router.StrictSlash(true)
	// file, _ := exec.LookPath(os.Args[0])
	// path, _ := filepath.Abs(file)
	// currPath := filepath.Dir(path)
	// logger.DebugF("currPath:%v", currPath)
	// router.Handle("/", &handler.DebugHandler{Logger: logger, OrginalHandler: FileResource(currPath + "/static/index.html")})
	// router.Handle("/favicon.ico", &handler.DebugHandler{Logger: logger, OrginalHandler: FileResource(currPath + "/static/favicon.ico")})
	// router.Handle("/static/{_dummy:.*}", &handler.DebugHandler{Logger: logger, OrginalHandler: http.StripPrefix("/static/", http.FileServer(http.Dir(currPath+"/static/")))})
	server := &http.Server{
		Addr:    addr,
		Handler: router,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	app := &Restapp{
		Logger:  logger,
		Router:  router,
		Address: addr,

		doneChan:     doneChan,
		server:       server,
		ctx:          ctx,
		onSignalFunc: []func(os.Signal){},
	}

	go func() {
		sig := <-osSignalChan
		for _, f := range app.onSignalFunc {
			f(sig)
		}
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
		for _, f := range app.onSignalFunc {
			f(sig)
		}
		logger.DebugF("received sig: %v", sig)
		if err := server.Close(); err == nil || err == http.ErrServerClosed {
			logger.Debug("Close ok")
			doneChan <- nil
		} else {
			logger.ErrorF("Close err: %v", err)
			doneChan <- err
		}
	}()

	return app
}

// Handle todo
func (s *Restapp) Handle(p string, h http.Handler) *Restapp {
	s.Router.Name(p).Path(p).Handler(h)
	return s
}

// HandleFunc todo
func (s *Restapp) HandleFunc(p string, f func(http.ResponseWriter, *http.Request)) *Restapp {
	s.Router.Name(p).Path(p).HandlerFunc(f)
	return s
}

// HandleController todo
func (s *Restapp) HandleController(c Controller) *Restapp {
	for k, v := range c.GetRoute() {
		s.Logger.DebugF("registed request path: %v %v", k, v)
		s.Handle(k, &handler.DebugHandler{Logger: s.Logger, OrginalHandler: v})
	}
	return s
}

// RegisterOnShutdown todo
func (s *Restapp) RegisterOnShutdown(f func()) {
	s.server.RegisterOnShutdown(f)
}

// RegisterOnSignal todo
func (s *Restapp) RegisterOnSignal(f func(os.Signal)) {
	s.onSignalFunc = append(s.onSignalFunc, f)
}

// RegisterSignalChan todo
func (s *Restapp) RegisterSignalChan(c chan os.Signal) {
	s.RegisterOnSignal(func(s os.Signal) {
		c <- s
	})
}

// Run todo
func (s *Restapp) Run(level log.Level) error {
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
func (s *Restapp) SimpleRun() error {
	return s.Run(log.Info)
}

// StaticResource todo
func StaticResource(prefixs []string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, prefix := range prefixs {
			if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
				r2 := new(http.Request)
				*r2 = *r
				r2.URL = new(url.URL)
				*r2.URL = *r.URL
				r2.URL.Path = p
				h.ServeHTTP(w, r2)
			}
		}
		http.NotFound(w, r)
	})
}

// FileResource todo
func FileResource(filepath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath)
	})
}
