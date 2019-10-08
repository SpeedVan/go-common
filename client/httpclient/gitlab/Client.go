package gitlab

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

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

// GetTree todo
func (s *Client) GetTree(group, project, sha, path string) ([]*TreeNode, error) {
	url := "https://" + s.Domain +
		"/api/v4/projects/" + group + "%2F" + project +
		"/repository/tree?ref=" + sha + "&path=" + path + "&per_page=500"
	println(url)
	req, _ := http.NewRequest("GET", url, http.NoBody)
	req.Header.Set("Private-Token", s.PrimaryToken)
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	// 后续考虑使用流处理方式解决json转化问题
	bytes, _ := ioutil.ReadAll(res.Body)
	nodes := []*TreeNode{}
	err = json.Unmarshal(bytes, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

// GetFile todo
func (s *Client) GetFile(group, project, sha, path string) (io.ReadCloser, error) {
	url := "https://" + s.Domain +
		"/api/v4/projects/" + group + "%2F" + project +
		"/repository/files/" + url.QueryEscape(path) + "/raw?ref=" + sha
	println(url)
	req, _ := http.NewRequest("GET", url, http.NoBody)
	req.Header.Set("Private-Token", s.PrimaryToken)
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
