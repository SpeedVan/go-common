package rest

import (
	"net/http"
	"reflect"
	"runtime"
	"sort"
	"strings"
)

// RouteMethod todo
type RouteMethod string

// RouteItem todo
type RouteItem struct {
	Path       string
	Method     string
	HandleFunc func(http.ResponseWriter, *http.Request)
}

// RouteHandler todo
type RouteHandler map[string]func(w http.ResponseWriter, r *http.Request)

func (s RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := s[r.Method]; ok {
		handler(w, r)
	} else {
		allow := []string{}
		for k := range s {
			allow = append(allow, k)
		}
		sort.Strings(allow)
		w.Header().Set("Allow", strings.Join(allow, ", "))
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (s RouteHandler) String() string {
	arr := []string{}
	for k, v := range s {
		arr = append(arr, "\""+k+"\": \""+runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()+"\"")
	}
	return "{" + strings.Join(arr, ",") + "}"
}

// AllMethodHandler todo
type AllMethodHandler func(http.ResponseWriter, *http.Request)

func (s AllMethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s(w, r)
}

func (s AllMethodHandler) String() string {
	return "{" + "\"ALL\": \"" + runtime.FuncForPC(reflect.ValueOf(s).Pointer()).Name() + "\"" + "}"
}

// RouteMap todo
type RouteMap map[string]http.Handler

// Merge todo
func (s RouteMap) Merge(rm RouteMap) RouteMap {
	return MergeRouteMap(s, rm)
}

// MergeRouteMap todo
func MergeRouteMap(arr ...RouteMap) RouteMap {
	result := make(RouteMap)
	for _, item := range arr {
		for k, v := range item {
			result[k] = v
		}
	}
	return result
}

// NewRouteMap todo
func NewRouteMap(arr ...*RouteItem) RouteMap {
	result := make(RouteMap)
	for _, item := range arr {
		if item.Method != "" {
			if _, ok := result[item.Path]; !ok {
				result[item.Path] = make(RouteHandler)
			}
			if h, ok := result[item.Path].(RouteHandler); ok {
				h[item.Method] = item.HandleFunc
			}
		} else {
			result[item.Path] = AllMethodHandler(item.HandleFunc)
		}
	}
	return result
}
