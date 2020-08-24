package rest

import (
	"errors"
	"net/http"
	"testing"

	"github.com/alpha-ss/go-common/config/mock"
	"github.com/alpha-ss/go-common/log"
)

func TestRestappForTesting(t *testing.T) {
	logger := log.NewCommon(log.Debug) // this level control webapp init log level

	m := map[string]string{
		"WEBAPP_LISTEN_ADDRESS": ":9999",
	}
	cfg := mock.New(m)

	restapp := New(cfg, logger)
	restapp.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
		w.Write([]byte("abc"))
	})
	app := NewForTesting(restapp)
	app.URLTest("/check").Do("GET", map[string]string{}, nil).Assert(func(res *ResExpect) error {
		return res.Err
	})

	app.URLTest("/check").Do("GET", map[string]string{}, nil).AssertRes(func(res *http.Response) error {
		if res.StatusCode != 202 {
			return errors.New("/check is not 202")
		}
		return nil
	})
}
