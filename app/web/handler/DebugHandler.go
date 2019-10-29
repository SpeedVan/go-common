package handler

import (
	"fmt"
	"net/http"
)

// DebugHandler todo
type DebugHandler struct {
	http.Handler
	OrginalHandler http.Handler
}

func (s *DebugHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v %v %v\n", r.Method, r.URL.Path, r.URL.RawQuery, r.Header)
	s.OrginalHandler.ServeHTTP(w, r)
}
