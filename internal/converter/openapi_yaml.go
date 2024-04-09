package converter

import (
	"bytes"

	_ "embed"

	"encoding/json"

	"github.com/itchyny/json2yaml"

	"github.com/jsightapi/jsight-api-core/catalog/ser/openapi"

	"github.com/jsightapi/jsight-api-core/kit"
	
	"io"
)

func newOpenapiYAML() OpenapiFmt {
	o := OpenapiFmt{}
	o.cnv = o.convertYAML
	return o
}

func (OpenapiFmt) convertYAML(japi kit.JApi, out io.Writer) error {
	oa, err := openapi.NewOpenAPI(japi.Catalog())
	if err != nil {
		return FormatError(err, FormatYAML)
	}
	respm, err := json.MarshalIndent(oa, "", "  ")
	if err != nil {
		return FormatError(err, FormatYAML)
	}

	resp, err := jsonToYAML(respm)
	if err != nil {
		return FormatError(err, FormatYAML)
	}

	_, err = out.Write(resp)
	if err != nil {
		return FormatError(err, FormatYAML)
	}
	return nil
}

func jsonToYAML(jsonData []byte) ([]byte, error) {
	from := bytes.NewReader(jsonData)
	to := bytes.NewBuffer(make([]byte, 0, len(jsonData)))

	err := json2yaml.Convert(to, from)
	if err != nil {
		return nil, err
	}

	return to.Bytes(), nil
}
