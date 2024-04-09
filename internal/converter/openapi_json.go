package converter

import (
	_ "embed"

	"encoding/json"

	"github.com/jsightapi/jsight-api-core/catalog/ser/openapi"

	"github.com/jsightapi/jsight-api-core/kit"
	
	"io"
)

func newOpenapiJSON() OpenapiFmt {
	o := OpenapiFmt{}
	o.cnv = o.convertJSON
	return o
}

func (OpenapiFmt) convertJSON(japi kit.JApi, out io.Writer) error {
	oa, err := openapi.NewOpenAPI(japi.Catalog())
	if err != nil {
		return FormatError(err, FormatJSON)
	}
	resp, err := json.MarshalIndent(oa, "", "  ")
	if err != nil {
		return FormatError(err, FormatJSON)
	}
	_, err = out.Write(resp)
	if err != nil {
		return FormatError(err, FormatJSON)
	}
	return nil
}
