package cio2

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/contextio/contextio-go/cioutil"
)

// TestNewCio2 tests the construction of Cio2
func TestNewCio2(t *testing.T) {
	t.Parallel()
	NewTestCio2(t)
}

// TestNewCio2WithLogger tests the construction of Cio2 and *TestLogger objects
func TestNewCio2WithLogger(t *testing.T) {
	t.Parallel()
	NewTestCio2WithLogger(t)
}

// TestNewCio2WithLogger tests the construction of Cio2 and *TestLogger objects
func TestNewTestCio2Server(t *testing.T) {
	t.Parallel()
	_, _, testServer, _ := NewTestCio2WithLoggerAndTestServer(t)
	defer testServer.Close()
}

// NewTestCio2 returns a new Cio2 object
func NewTestCio2(t *testing.T) Cio2 {
	return NewCio2(getEnv(t, "UNSUB_CIO_API_KEY"), getEnv(t, "UNSUB_CIO_API_SECRET"))
}

// NewTestCio2WithLogger returns a new Cio2 object and *TestLogger object
func NewTestCio2WithLogger(t *testing.T) (Cio2, *cioutil.TestLogger) {
	logger := &cioutil.TestLogger{Buffer: &bytes.Buffer{}}
	cio2 := NewCio2WithLogger(getEnv(t, "UNSUB_CIO_API_KEY"), getEnv(t, "UNSUB_CIO_API_SECRET"), logger)
	return cio2, logger
}

// NewTestCio2WithLoggerAndTestServer returns a new Cio2, *TestLogger, and *httptest.Server objects
func NewTestCio2WithLoggerAndTestServer(t *testing.T) (Cio2, *cioutil.TestLogger, *httptest.Server, *http.ServeMux) {
	logger := &cioutil.TestLogger{Buffer: &bytes.Buffer{}}
	mux := http.NewServeMux()
	cio2, server := NewTestCio2Server(getEnv(t, "UNSUB_CIO_API_KEY"), getEnv(t, "UNSUB_CIO_API_SECRET"), logger, mux)
	return cio2, logger, server, mux
}

// getEnv returns the named environment variable, or causes t.Fatal
func getEnv(t *testing.T, envVarName string) string {
	val := os.Getenv(envVarName)
	if len(val) == 0 {
		t.Fatal("Empty Environment Variable for: " + envVarName)
	}
	return val
}

// Must panics if error is not nil
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
