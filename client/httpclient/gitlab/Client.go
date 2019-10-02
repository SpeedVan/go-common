package gitlab

import (
	"net/http"

	"github.com/SpeedVan/go-common/client/httpclient"
	"github.com/SpeedVan/go-common/config"
)

// Client todo
type Client struct {
	HTTPClient   *http.Client
	PrimaryToken string // sF7us_xdFTBseuKeyvNo
}

// New todo
func New(config config.Config) (*Client, error) {
	primaryToken := config.Get("GITLAB_PRIVATE_TOKEN")
	httpClient, err := httpclient.New(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		HTTPClient:   httpClient,
		PrimaryToken: primaryToken,
	}, nil
}

// Get todo
func (s *Client) Get() error {

	return nil
}
