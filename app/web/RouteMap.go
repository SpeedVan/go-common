package web

import "net/http"

// RouteMethods todo
type RouteMethods []string

// RouteItem todo
type RouteItem struct {
	Path       string
	Methods    RouteMethods
	HandleFunc func(http.ResponseWriter, *http.Request)
}

// RouteMap todo
type RouteMap map[string]*RouteItem

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
		result[item.Path] = item
	}
	return result
}
