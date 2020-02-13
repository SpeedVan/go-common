package web

import (
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/SpeedVan/go-common/app/rest"
	"github.com/SpeedVan/go-common/app/rest/handler"
	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/go-common/log"
)

// New todo
func New(config config.Config, logger log.Logger) *rest.Restapp {
	app := rest.New(config, logger)
	// .StrictSlash(true)
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	currPath := filepath.Dir(path)
	logger.DebugF("currPath:%v", currPath)
	app.Router.Handle("/", &handler.DebugHandler{Logger: logger, OrginalHandler: FileResource(currPath + "/static/index.html")})
	app.Router.Handle("/favicon.ico", &handler.DebugHandler{Logger: logger, OrginalHandler: FileResource(currPath + "/static/favicon.ico")})
	app.Router.Handle("/static/{_dummy:.*}", &handler.DebugHandler{Logger: logger, OrginalHandler: http.StripPrefix("/static/", http.FileServer(http.Dir(currPath+"/static/")))})
	return app
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
