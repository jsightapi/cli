package generator

import (
	"bytes"

	"github.com/jsightapi/cli/internal/format"

	_ "embed"

	"encoding/json"

	"github.com/itchyny/json2yaml"

	"github.com/jsightapi/jsight-api-core/catalog/ser/openapi"

	"github.com/jsightapi/jsight-api-core/kit"

	"io"
)

func newOpenapiYAML() common {
	o := common{}
	o.gen = o.convertYAML
	return o
}

func (common) convertYAML(japi kit.JApi, out io.Writer) error {
	var err error
	var oa *openapi.OpenAPI
	oa, err = openapi.NewOpenAPI(japi.Catalog())
	if err == nil {
		var respm []byte
		respm, err = json.MarshalIndent(oa, "", "  ")
		if err == nil {
			var resp []byte
			resp, err = jsonToYAML(respm)
			if err == nil {
				_, err = out.Write(resp)
				if err == nil {
					return nil
				}
			}
		}
	}
	return FormatError(err, format.FormatYAML)
}

func jsonToYAML(jsonData []byte) ([]byte, error) {
	from := bytes.NewReader(jsonData)
	to := bytes.NewBuffer(make([]byte, 0, len(jsonData)))

	err := json2yaml.Convert(to, from)
	if err != nil {
		return nil, err
	}

	return removeLastLineBreak(to), nil
}

func removeLastLineBreak(input *bytes.Buffer) []byte {
	if bs := input.Bytes(); len(bs) > 0 && bs[len(bs)-1] == '\n' {
		return bs[:len(bs)-1]
	}
	return input.Bytes()
}
