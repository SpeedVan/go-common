package httpclient

import (
	"net/http"

	"github.com/SpeedVan/go-common/config"
)

// New todo
func New(config config.Config) (*http.Client, error) {
	cl := &http.Client{}
	return cl, nil
}
