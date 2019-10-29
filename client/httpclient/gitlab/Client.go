package gitlab

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/SpeedVan/go-common/client/httpclient"
	"github.com/SpeedVan/go-common/client/httpclient/gitlab/graphql"
	"github.com/SpeedVan/go-common/config"
)

// Client todo
type Client struct {
	HTTPClient   *http.Client
	PrivateToken string // sF7us_xdFTBseuKeyvNo
	Domain       string // gitlab.com
}

// New todo
func New(config config.Config) (*Client, error) {
	privateToken := config.Get("PRIVATE_TOKEN")
	domain := config.Get("DOMAIN")
	httpClient, err := httpclient.New(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		HTTPClient:   httpClient,
		PrivateToken: privateToken,
		Domain:       domain,
	}, nil
}

// GetTree todo
func (s *Client) GetTree(protocol, group, project, sha, path string) ([]*TreeNode, error) {
	urlPath := protocol + "://" + s.Domain +
		"/api/v4/projects/" + group + "%2F" + project +
		"/repository/tree?ref=" + sha + "&path=" + path + "&per_page=500"
	println(urlPath)
	req, _ := http.NewRequest("GET", urlPath, http.NoBody)
	req.Header.Set("Private-Token", s.PrivateToken)
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
func (s *Client) GetFile(protocol, group, project, sha, path string) (io.ReadCloser, http.Header, error) {
	urlPath := protocol + "://" + s.Domain +
		"/api/v4/projects/" + group + "%2F" + project +
		"/repository/files/" + url.PathEscape(path) + "/raw?ref=" + sha
	println(urlPath)
	req, _ := http.NewRequest("GET", urlPath, http.NoBody)
	req.Header.Set("Private-Token", s.PrivateToken)
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	return res.Body, res.Header, nil
}

// HeadFile todo
func (s *Client) HeadFile(protocol, group, project, sha, path string) (http.Header, error) {
	urlPath := protocol + "://" + s.Domain +
		"/api/v4/projects/" + group + "%2F" + project +
		"/repository/files/" + url.PathEscape(path) + "/raw?ref=" + sha
	println(urlPath)
	req, _ := http.NewRequest("HEAD", urlPath, http.NoBody)
	req.Header.Set("Private-Token", s.PrivateToken)
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Header, nil
}

// GetBlobSizeFromBody Because there is no blol size in (all) old Gitlab api header, the blob size only can get from body with blob api
func (s *Client) GetBlobSizeFromBody(protocol, group, project, blobID string) (string, error) {
	urlPath := protocol + "://" + s.Domain +
		"/api/v4/projects/" + group + "%2F" + project +
		"/repository/blobs/" + blobID
	println(urlPath)
	req, _ := http.NewRequest("GET", urlPath, http.NoBody)
	req.Header.Set("Private-Token", s.PrivateToken)
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	reader := bufio.NewReader(res.Body)
	line, err := reader.ReadString(':')
	if err != nil {
		return "", err
	}
	line, err = reader.ReadString(',')
	if err != nil {
		return "", err
	}
	size := line[:len(line)-1]

	return size, nil
}

// GetCommits todo
func (s *Client) GetCommits(protocol, group, project, sha, path string) ([]*Commit, error) {

	urlPath := protocol + "://" + s.Domain +
		"/api/v4/projects/" + group + "%2F" + project +
		"/repository/commits"

	querys := []string{}
	if sha != "" {
		querys = append(querys, "sha="+sha)
	}
	if path != "" {
		querys = append(querys, "path="+path)
	}
	rawQuery := strings.Join(querys, "&")
	if rawQuery != "" {
		urlPath = urlPath + "?" + rawQuery
	}
	println(urlPath)
	req, _ := http.NewRequest("GET", urlPath, http.NoBody)
	req.Header.Set("Private-Token", s.PrivateToken)
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := []*Commit{}
	err = json.Unmarshal(resBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetProjects todo
func (s *Client) GetProjects(protocol, group string) ([]*Project, error) {

	urlPath := protocol + "://" + s.Domain +
		"/api/v4/groups/" + group + "/projects"

	// querys := []string{}
	// if sha != "" {
	// 	querys = append(querys, "sha="+sha)
	// }
	// if path != "" {
	// 	querys = append(querys, "path="+path)
	// }
	// rawQuery := strings.Join(querys, "&")
	// if rawQuery != "" {
	// 	urlPath = urlPath + "?" + rawQuery
	// }
	println(urlPath)
	req, _ := http.NewRequest("GET", urlPath, http.NoBody)
	req.Header.Set("Private-Token", s.PrivateToken)
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := []*Project{}
	err = json.Unmarshal(resBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetGroupProjects todo
func (s *Client) GetGroupProjects(protocol string) ([]*Project, error) {
	urlPath := protocol + "://" + s.Domain +
		"/api/v4/projects?membership=true&simple=true"

	// querys := []string{}
	// if sha != "" {
	// 	querys = append(querys, "sha="+sha)
	// }
	// if path != "" {
	// 	querys = append(querys, "path="+path)
	// }
	// rawQuery := strings.Join(querys, "&")
	// if rawQuery != "" {
	// 	urlPath = urlPath + "?" + rawQuery
	// }
	println(urlPath)
	req, _ := http.NewRequest("GET", urlPath, http.NoBody)
	req.Header.Set("Private-Token", s.PrivateToken)
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := []*Project{}
	err = json.Unmarshal(resBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Graphql todo
func (s *Client) Graphql(protocol, group, project, sha, path string) (*graphql.Graphql, error) {
	urlPath := protocol + "://" + s.Domain + "/api/graphql"
	println(urlPath)
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
	req, _ := http.NewRequest("POST", urlPath, bytes.NewReader(byteData))
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
