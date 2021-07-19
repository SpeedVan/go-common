package web

import (
	"testing"

	"github.com/alpha-supsys/go-common/log"
	"github.com/alpha-supsys/go-common/mock/config"
)

func Test(t *testing.T) {
	logger := log.NewCommon(log.Debug) // this level control webapp init log level

	m := map[string]string{
		"WEBAPP_LISTEN_ADDRESS": ":9999",
	}
	cfg := config.New(m)

	app := New(cfg, logger)
	app.Run(log.Debug) // this level control webapp runtime log level

}
