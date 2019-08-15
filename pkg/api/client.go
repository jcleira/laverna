package api

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Client its api main struct, it defines the information needed for the api to
// generate API endpoints from YAML files.
//
// You might use a YAML file like the following to initializa an api.Client
//
// client := *api.Client {
// YAML: `
// ---
// - endpoint:
//   url: /steps
//   status_code: 200
//   response:
//   - step:
//     description: Good or Evil?
//     options:
//     - option:
//       value: Good
//     - option: Good
//       value: Evil
//`
//}
type Client struct {
	YAML string
}

// Endpoint defines the attributes available to configure an endpoint from a
// YAML file, a valid YML file is an array of endpoints.
type Endpoint struct {
	URL        string      `yaml:"url"`
	StatusCode int         `yaml:"status_code"`
	Response   interface{} `yaml:"response"`
}

// Endpoints provides a shortcut to work with an array of Endpoint.
type Endpoints []Endpoint

// Parse will parse a previously configured YAML file (on the client), and
// create an array of Endpoint.
//
// Returns an array of Endpoints, error otherwise.
func (c *Client) Parse() (*Endpoints, error) {
	endpoints := &Endpoints{}

	if err := yaml.Unmarshal([]byte(c.YAML), endpoints); err != nil {
		return nil, errors.Wrap(err, "error parsing the YAML configuration")
	}

	return endpoints, nil
}
