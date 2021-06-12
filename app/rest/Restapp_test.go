package rest

import (
	"net/http"
	"testing"

	"github.com/alpha-supsys/go-common/config/mock"
	"github.com/alpha-supsys/go-common/log"
)

func Test(t *testing.T) {
	logger := log.NewCommon(log.Debug) // this level control webapp init log level

	m := map[string]string{
		"WEBAPP_LISTEN_ADDRESS": ":9999",
	}
	cfg := mock.New(m)

	app := New(cfg, logger)

	app.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {

	})

	go app.Run(log.Debug) // this level control webapp runtime log level

}
