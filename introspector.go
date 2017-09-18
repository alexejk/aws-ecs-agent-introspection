package ecsagent

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// DefaultEcsAgentEndpoint is the standard endpoint that ECS Agent is available on
const DefaultEcsAgentEndpoint = "http://localhost:51678/"

// Introspector is a client towards ECS Agent introspection API
type Introspector struct {
	client   *http.Client
	endpoint string
}

// NewIntrospector creates an new instance of ECS Agent Introspector
func NewIntrospector() *Introspector {

	return &Introspector{
		client: &http.Client{
			Timeout: time.Second * 2,
		},
		endpoint: DefaultEcsAgentEndpoint,
	}
}

// NewIntrospectorWithClient allows customization of HTTP Client used to create ECS Agent Introspector
func NewIntrospectorWithClient(client *http.Client) *Introspector {
	return &Introspector{
		client:   client,
		endpoint: DefaultEcsAgentEndpoint,
	}
}

// SetEndpoint allows customizing API Endpoint that client will use.
// If provided endpoint does not end with a trailing slash, it will be added
func (i *Introspector) SetEndpoint(endpoint string) {

	if !strings.HasSuffix(endpoint, "/") {
		i.endpoint = endpoint + "/"
	} else {
		i.endpoint = endpoint
	}
}

func (i *Introspector) get(path string, responseHolder interface{}) error {

	url := fmt.Sprintf("%s%s", i.endpoint, path)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	res, getErr := i.client.Do(req)
	if getErr != nil {
		return getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}

	jsonErr := json.Unmarshal(body, responseHolder)
	return jsonErr
}
