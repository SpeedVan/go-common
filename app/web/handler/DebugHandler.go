package handler

import (
	"net/http"

	"github.com/SpeedVan/go-common/log"
)

// DebugHandler todo
type DebugHandler struct {
	Logger log.Logger
	http.Handler
	OrginalHandler http.Handler
}

func (s *DebugHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Logger.DebugF("%v %v %v %v", r.Method, r.URL.Path, r.URL.RawQuery, r.Header)
	s.OrginalHandler.ServeHTTP(w, r)
}
