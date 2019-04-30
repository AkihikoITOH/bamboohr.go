package bamboohr_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"gotest.tools/assert"

	bamboohr "github.com/AkihikoITOH/bamboohr.go"
)

const (
	apiDomain = "test"
	apiKey    = "bamboohr-api-secret-key"
)

var (
	mockAPIConfig = bamboohr.NewConfig(apiDomain, apiKey)
)

type MockHTTPClient struct {
	Request  *http.Request
	Response *http.Response
}

func (mockClient *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	mockClient.Request = req
	return mockClient.Response, nil
}

func NewMockClient(body []byte) *MockHTTPClient {
	responseHeader := http.Header{}
	responseHeader.Add("Date", "Fri, 08 Mar 2019 19:36:47 GMT")
	responseHeader.Add("Content-Type", "application/json")
	responseBody := ioutil.NopCloser(bytes.NewReader(body))
	mockResponse := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/2.0",
		ProtoMajor:    2,
		ProtoMinor:    0,
		Header:        responseHeader,
		Body:          responseBody,
		ContentLength: -1,
		Close:         false,
		Uncompressed:  false,
	}
	return &MockHTTPClient{Response: mockResponse}
}

func SharedRequestTests(t *testing.T, req *http.Request) {
	assert.Equal(t, req.Host, bamboohr.BambooHRAPIHost)
	user, pass, ok := req.BasicAuth()
	assert.Equal(t, user, apiKey)
	assert.Equal(t, pass, "x")
	assert.Assert(t, ok)
}

func TestNewAPI(t *testing.T) {
	config := &bamboohr.Config{}
	api := bamboohr.NewAPI(config)

	assert.Equal(t, reflect.TypeOf(api).String(), "*bamboohr.API")
}
