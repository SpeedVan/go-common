package gitlab

import (
	"net/http"

	"github.com/SpeedVan/go-common/config"
)

type GitlabClient struct {
	HttpClient   *http.Client
	PrimaryToken string
}

func New(config config.Config) (*GitlabClient, error) {
	primaryToken := config.Get("primaryToken")
}
