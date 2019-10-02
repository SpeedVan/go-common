package httpclient

import (
	"net"
	"net/http"
	"time"

	"github.com/SpeedVan/go-common/config"
)

// New todo
func New(config config.Config) (*http.Client, error) {
	transport := &http.Transport{
		MaxIdleConnsPerHost:   256,
		IdleConnTimeout:       90 * time.Second,
		DisableCompression:    true,
		MaxIdleConns:          100,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	transport.DialContext = (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext

	timeout := time.Duration(100 * time.Second)
	cl := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
	return cl, nil
}
