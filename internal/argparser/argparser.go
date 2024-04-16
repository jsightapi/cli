package argparser

import (
	"bytes"

	_ "embed"

	"encoding/json"

	"errors"

	"fmt"

	"github.com/itchyny/json2yaml"

	"github.com/jsightapi/jsight-api-core/catalog/ser/openapi"

	"github.com/jsightapi/jsight-api-core/jerr"

	"github.com/jsightapi/jsight-api-core/kit"

	"io"

	"github.com/jsightapi/cli/internal/statistics"

	"github.com/jsightapi/cli/internal/format"

	"github.com/urfave/cli/v2"

	"os"

	"text/template"
)

var noInputFileErrorString = "no input file"
var ErrUnsupportedFormat = errors.New("unsupported format")

//go:embed template.html
var templateContent string

func FormatError(err error, f format.Format) error {
	return fmt.Errorf("convert %s: %w", f, err)
}

func DoWork(f format.Format, ctx *cli.Context) error {
	var err error
	var specPath *string
	var sendStatFlag bool
	var fileSize int64
	var japi *kit.JApi
	var je *jerr.JApiError

	specPath, sendStatFlag, fileSize, err = parseArgs(ctx)
	switch f {
	case format.FormatHTML, format.FormatJSON, format.FormatYAML:
		if err == nil {
			japi, je, err = makeOut(f, *specPath, ctx.App.Writer)
			statistics.SendStat(japi, je, sendStatFlag, fileSize, err)
			return err
		}
		statistics.SendStat(nil, nil, sendStatFlag, fileSize, err)
		if err.Error() == noInputFileErrorString {
			err = cli.ShowCommandHelp(ctx, string(f))
		}
	default:
		err = ErrUnsupportedFormat
		statistics.SendStat(nil, nil, sendStatFlag, fileSize, err)
	}
	return err
}

func makeOut(f format.Format, filepath string, out io.Writer) (japi *kit.JApi, je *jerr.JApiError, err error) {
	var _japi kit.JApi
	_japi, je = kit.NewJapi(filepath)
	japi = &_japi
	if je != nil {
		filepath := "no file"
		if je.File != nil {
			filepath = je.File.Name()
		}
		err = fmt.Errorf("JSight API parsing error at `%v` [%v, %v]\nError message: %v\nQuote:\n----> %v",
			filepath,
			je.Location.Line,
			je.Location.Column,
			je.Msg,
			je.Quote)
		return japi, je, err
	}

	var jsonbytes []byte
	jsonbytes, err = japi.ToJson()
	if err != nil {
		return japi, je, fmt.Errorf("generate JSON: %w", err)
	}

	if f == format.FormatHTML {
		tmpl, err1 := template.New("").Parse(templateContent)
		if err1 == nil {
			return japi, je, tmpl.Execute(out, map[string]string{"jdoc": string(jsonbytes)})
		}
		return japi, je, fmt.Errorf("parse template: %w", err1)
	} else if f == format.FormatJSON || f == format.FormatYAML {
		var oa *openapi.OpenAPI
		oa, err = openapi.NewOpenAPI(japi.Catalog())
		if err == nil {
			var respm []byte
			respm, err = json.MarshalIndent(oa, "", "  ")
			if err == nil {
				var resp []byte
				if f == format.FormatYAML {
					resp, err = jsonToYAML(respm)
				} else {
					resp = respm
				}
				if err == nil {
					_, err = out.Write(resp)
					return japi, je, err
				}
			}
		}
		return japi, je, FormatError(err, f)
	}
	return japi, je, fmt.Errorf("unhandled error")
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

func parseArgs(ctx *cli.Context) (specPath *string, sendStatFlag bool, fileSize int64, e error) {
	var err error = nil
	sendStatFlag = !Contains(ctx.FlagNames(), "s")
	aa := ctx.Args()
	if aa.Len() == 1 {
		specPath = StringRef(aa.First())
		if specPath != nil {
			var fileInfo os.FileInfo
			fileInfo, err = os.Stat(*specPath)
			if err == nil {
				fileSize = fileInfo.Size()
				_, err = os.Open(*specPath)
				return specPath, sendStatFlag, fileSize, err
			}
		}
	}
	e = err
	if e == nil {
		e = errors.New(noInputFileErrorString)
	}
	return nil, sendStatFlag, 0, e
}

func Contains[T comparable](arr []T, x T) bool {
	for _, v := range arr {
		if v == x {
			return true
		}
	}
	return false
}

func StringRef(s string) *string {
	return &s
}
