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
	Domain       string // gitlab.com
}

// New todo
func New(config config.Config) (*Client, error) {
	primaryToken := config.Get("PRIVATE_TOKEN")
	domain := config.Get("DOMAIN")
	httpClient, err := httpclient.New(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		HTTPClient:   httpClient,
		PrimaryToken: primaryToken,
		Domain:       domain,
	}, nil
}

// Get todo
func (s *Client) Get(group, project, sha, path string) error {
	s.HTTPClient.Do()
	return nil
}

// GetTree todo
func (s *Client) GetTree(group, project, sha, path string) error {
	s.HTTPClient.Do()
	return nil
}
