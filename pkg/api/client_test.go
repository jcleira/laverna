package api

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Description   string
		YAML          string
		Endpoints     *Endpoints
		ExpectedError error
	}{
		{
			Description: "when Parse successfuly reads the endpoints to create",
			YAML: `
---
- endpoint:
  url: foo
- endpoint:
  url: bar
`,
			Endpoints: &Endpoints{
				Endpoint{
					URL: "foo",
				},
				Endpoint{
					URL: "bar",
				},
			},
		},
		{
			Description: "when Parse fails due a malformed YAML",
			YAML: `
---
- endpoint:
url: /foo
	- endpoint:
  url: /bar
`,
			ExpectedError: errors.New("error parsing the YAML configuration: yaml: line 3: did not find expected '-' indicator"),
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			client := &Client{
				YAML: test.YAML,
			}

			endpoints, err := client.Parse()

			assert.Equal(test.Endpoints, endpoints)

			if test.ExpectedError != nil {
				assert.Error(err)
				assert.EqualError(err, test.ExpectedError.Error())
			} else {
				assert.NoError(err)
			}
		})
	}
}
