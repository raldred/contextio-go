// Package cio2 is the Golang client library for the 2.0 Context.IO API
package cio2

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/contextio/contextio-go/cioutil"
)

const (
	// DefaultHost is the default host of CIO 2.0 API
	DefaultHost = "https://api.context.io/2.0"

	// DefaultRequestTimeout is the default timeout duration used on HTTP requests
	DefaultRequestTimeout = 120 * time.Second
)

// Cio2 struct contains the api key and secret, along with an optional logger,
// and provides convenience functions for accessing all CIO 2.0 endpoints.
type Cio2 struct {
	cioutil.Cio
}

// NewCio2 returns a CIO 2.0 struct (without a logger) for accessing the CIO 2.0 API.
func NewCio2(key string, secret string) Cio2 {
	return NewCio2WithLogger(key, secret, nil)
}

// NewCio2WithLogger returns a CIO 2.0 struct (with a logger) for accessing the CIO 2.0 API.
func NewCio2WithLogger(key string, secret string, logger cioutil.Logger) Cio2 {
	return Cio2{Cio: cioutil.NewCio(key, secret, logger, DefaultHost, DefaultRequestTimeout)}
}

// NewTestCio2Server is a convenience function that returns a Cio2 object
// and a *httptest.Server (which must be closed when done being used).
// The Cio2 instance will hit the test server for all requests.
func NewTestCio2Server(key string, secret string, logger cioutil.Logger, handler http.Handler) (Cio2, *httptest.Server) {
	testServer := httptest.NewServer(handler)
	testCio2 := Cio2{Cio: cioutil.NewCio(key, secret, logger, testServer.URL, 5*time.Second)}
	return testCio2, testServer
}
