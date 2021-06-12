package web

import (
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/alpha-supsys/go-common/app/rest"
	"github.com/alpha-supsys/go-common/app/rest/handler"
	"github.com/alpha-supsys/go-common/config"
	"github.com/alpha-supsys/go-common/log"
)

// Webapp todo
type Webapp struct {
	*rest.Restapp
	CurrentPath string
}

// RegistStaticFolder todo
func (s *Webapp) RegistStaticFolder(requestpath, dirpath string) {
	s.Router.Handle(requestpath+"/{_dummy:.*}", &handler.DebugHandler{Logger: s.Logger, OrginalHandler: http.StripPrefix("/"+dirpath+"/", http.FileServer(http.Dir(s.CurrentPath+"/"+dirpath+"/")))})
}

// RegistStaticFile todo
func (s *Webapp) RegistStaticFile(requestpath, filepath string) {

	s.Router.Handle(requestpath, &handler.DebugHandler{Logger: s.Logger, OrginalHandler: FileResource(s.CurrentPath + filepath)})
}

// New todo
func New(config config.Config, logger log.Logger) *Webapp {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	currPath := filepath.Dir(path)
	restApp := rest.New(config, logger)
	app := &Webapp{
		Restapp:     restApp,
		CurrentPath: currPath,
	}

	app.RegistStaticFile("/", "/static/index.html")
	app.RegistStaticFile("/favicon.ico", "/static/favicon.ico")
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
