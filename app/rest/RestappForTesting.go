package rest

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
)

type TestRequest struct {
	url *url.URL
}

func (s *TestRequest) Do(method string, header map[string]string, body io.Reader) *ResExpect {
	req, _ := http.NewRequest(method, s.url.String(), body)

	res, err := http.DefaultClient.Do(req)
	return &ResExpect{
		Res: res,
		Err: err,
	}
}

type ResExpect struct {
	Res *http.Response
	Err error
}

func (s *ResExpect) Assert(f func(err *ResExpect) error) {
	if err := f(s); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (s *ResExpect) AssertRes(f func(*http.Response) error) {
	if err := f(s.Res); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type RestappForTesting struct {
	*Restapp

	testServer *httptest.Server
}

// URLTestWithParams todo
func (s *RestappForTesting) URLTestWithParams(urlName string, urlParams []string, queryParams map[string]string) *TestRequest {
	localServerURL, _ := url.Parse(s.testServer.URL)
	urlPtr, _ := s.Restapp.Router.Get(urlName).URL()
	localServerURL.Path = urlPtr.Path
	for k, v := range queryParams {
		localServerURL.Query().Add(k, v)
	}
	return &TestRequest{
		url: localServerURL,
	}
}

// URLTest todo
func (s *RestappForTesting) URLTest(urlName string) *TestRequest {
	return s.URLTestWithParams(urlName, []string{}, map[string]string{})
}

func NewForTesting(restapp *Restapp) *RestappForTesting {
	ts := httptest.NewServer(restapp.Router)
	return &RestappForTesting{
		Restapp:    restapp,
		testServer: ts,
	}
}
