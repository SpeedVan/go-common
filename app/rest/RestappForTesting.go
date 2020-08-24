package rest

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
)

var (
	// EmptyParams todo
	EmptyParams = map[string]string{}
)

// TestRequest todo
type TestRequest struct {
	url *url.URL
}

// Do todo
func (s *TestRequest) Do(method string, header map[string]string, body io.Reader) *ResExpect {
	req, _ := http.NewRequest(method, s.url.String(), body)

	res, err := http.DefaultClient.Do(req)
	return &ResExpect{
		Res: res,
		Err: err,
	}
}

// ResExpect todo
type ResExpect struct {
	Res *http.Response
	Err error
}

// Assert todo
func (s *ResExpect) Assert(f func(err *ResExpect) error) {
	if err := f(s); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// AssertRes todo
func (s *ResExpect) AssertRes(f func(*http.Response) error) {
	if err := f(s.Res); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// RestappForTesting todo
type RestappForTesting struct {
	*Restapp

	testServer *httptest.Server
}

// URLTestWithParams todo
func (s *RestappForTesting) URLTestWithParams(urlName string, urlParams map[string]string, queryParams map[string]string) *TestRequest {
	localServerURL, err := url.Parse(s.testServer.URL)
	if err != nil {
		fmt.Println(err)
	}
	paramsArr := []string{}
	for k, v := range urlParams {
		paramsArr = append(paramsArr, k, v)
	}
	urlPtr, err := s.Restapp.Router.Get(urlName).URL(paramsArr...)
	if err != nil {
		fmt.Println(err)
	}
	localServerURL.Path = urlPtr.Path
	query := localServerURL.Query()
	for k, v := range queryParams {
		query.Add(k, v)
	}
	localServerURL.RawQuery = query.Encode()
	return &TestRequest{
		url: localServerURL,
	}
}

// URLTest todo
func (s *RestappForTesting) URLTest(urlName string) *TestRequest {
	return s.URLTestWithParams(urlName, EmptyParams, EmptyParams)
}

// NewForTesting todo
func NewForTesting(restapp *Restapp) *RestappForTesting {
	ts := httptest.NewServer(restapp.Router)
	return &RestappForTesting{
		Restapp:    restapp,
		testServer: ts,
	}
}
