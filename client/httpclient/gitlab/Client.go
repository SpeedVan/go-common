package gitlab

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/SpeedVan/go-common/client/httpclient"
	"github.com/SpeedVan/go-common/client/httpclient/gitlab/graphql"
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
func (s *Client) GetFile(group, project, sha, path string) (io.ReadCloser, http.Header, error) {
	url := "https://" + s.Domain +
		"/api/v4/projects/" + group + "%2F" + project +
		"/repository/files/" + url.QueryEscape(path) + "/raw?ref=" + sha
	println(url)
	req, _ := http.NewRequest("GET", url, http.NoBody)
	req.Header.Set("Private-Token", s.PrimaryToken)
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	return res.Body, res.Header, nil
}

// HeadFile todo
func (s *Client) HeadFile(group, project, sha, path string) (http.Header, error) {
	url := "https://" + s.Domain +
		"/api/v4/projects/" + group + "%2F" + project +
		"/repository/files/" + url.QueryEscape(path) + "/raw?ref=" + sha
	println(url)
	req, _ := http.NewRequest("HEAD", url, http.NoBody)
	req.Header.Set("Private-Token", s.PrimaryToken)
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Header, nil
}

// Graphql todo
func (s *Client) Graphql(group, project, sha, path string) (*graphql.Graphql, error) {
	url := "https://gitlab.com/api/graphql"
	println(url)
	query :=
		"{\n" +
			"  project(fullPath: \"" + group + "/" + project + "\") {\n" +
			"    repository {\n" +
			"      tree(path:\"" + path + "\", ref:\"" + sha + "\") {\n" +
			"        trees {\n" +
			"          nodes {\n" +
			"            flatPath\n" +
			"            id\n" +
			"            name\n" +
			"            path\n" +
			"            type\n" +
			"            webUrl\n" +
			"          }\n" +
			"        }\n" +
			"        blobs {\n" +
			"          nodes {\n" +
			"            flatPath\n" +
			"            id\n" +
			"            name\n" +
			"            path\n" +
			"            type\n" +
			"            webUrl\n" +
			"          }\n" +
			"        }\n" +
			"      }\n" +
			"    }\n" +
			"  }\n" +
			"}"
	byteData, _ := json.Marshal(&graphql.Payload{
		Query: query,
	})
	println(string(byteData))
	req, _ := http.NewRequest("POST", url, bytes.NewReader(byteData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Private-Token", "sF7us_xdFTBseuKeyvNo")
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	println(string(resBytes))
	result := &graphql.Graphql{}
	err = json.Unmarshal(resBytes, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
