package ecsagent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntrospector_SetEndpoint(t *testing.T) {

	expectation := "http://example.com/"

	api := NewIntrospector()

	// Setting endpoint with trailing slash
	api.SetEndpoint("http://example.com/")
	assert.Equal(t, expectation, api.endpoint)

	// Setting endpoint with trailing slash
	api.SetEndpoint("http://example.com")
	assert.Equal(t, expectation, api.endpoint)
}

//----- Below are support functions

func loadTestData(t *testing.T, name string) []byte {
	path := filepath.Join("testData", name) // relative path
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}

func loadTestDataString(t *testing.T, name string) string {
	return string(loadTestData(t, name))
}

func testServerAndClient(t *testing.T, testFile string) (*httptest.Server, *Introspector) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, loadTestDataString(t, testFile))
	}))

	api := NewIntrospector()
	api.SetEndpoint(ts.URL)

	return ts, api
}

func failingServerAndClient(t *testing.T, errCode int, errMessage string) (*httptest.Server, *Introspector) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, errMessage, errCode)
	}))

	api := NewIntrospector()
	api.SetEndpoint(ts.URL)

	return ts, api
}
