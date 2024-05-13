package generator

import (
	_ "embed"

	"github.com/jsightapi/cli/internal/format"

	"encoding/json"

	"github.com/jsightapi/jsight-api-core/catalog/ser/openapi"

	"github.com/jsightapi/jsight-api-core/kit"

	"io"
)

func newOpenapiJSON() common {
	o := common{}
	o.gen = o.convertJSON
	return o
}

func (common) convertJSON(japi kit.JApi, out io.Writer) error {
	var err error
	var oa *openapi.OpenAPI
	oa, err = openapi.NewOpenAPI(japi.Catalog())
	if err == nil {
		var resp []byte
		resp, err = json.MarshalIndent(oa, "", "  ")
		if err == nil {
			_, err = out.Write(addLastLineBreak(resp))
			if err == nil {
				return nil
			}
		}
	}
	return FormatError(err, format.FormatJSON)
}

func addLastLineBreak(input []byte) []byte {
	if len(input) > 0 && input[len(input)-1] != '\n' {
		return append(input, '\n')
	}
	return input
}
