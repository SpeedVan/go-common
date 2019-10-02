package controller

import "net/http"

// Controller todo
type Controller interface {
	GetRoute() map[string]struct {
		Methods    []string
		HandleFunc func(http.ResponseWriter, *http.Request)
	}
}
