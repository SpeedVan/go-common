package gitlabclient

import (
	"net/http"

	"github.com/SpeedVan/go-common/client/httpclient"
	"github.com/SpeedVan/go-common/config"
)

// GitlabClient todo
type GitlabClient struct {
	HTTPClient   *http.Client
	Domain       string // https://www.gitlab.com/api/v4
	Username     string // SpeedVan
	PrivateToken string // sF7us_xdFTBseuKeyvNo
}

// New todo
func New(config config.Config) (*GitlabClient, error) {
	privateToken := config.Get("GITLAB_PRIVATE_TOKEN")
	username := config.Get("GITLAB_USERNAME")

	cl, err := httpclient.New(config)
	if err != nil {
		return nil, err
	}
	return &GitlabClient{
		HTTPClient:   cl,
		PrivateToken: privateToken,
		Username:     username,
	}, nil
}

// Get todo
func (s *GitlabClient) Get() error {

	return nil
}
