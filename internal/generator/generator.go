package generator

import (
	"github.com/jsightapi/cli/internal/statistics"

	"github.com/urfave/cli/v2"

	"github.com/jsightapi/cli/internal/format"

	"errors"

	"fmt"

	"io"

	"github.com/jsightapi/jsight-api-core/kit"
)

// Generator an abstraction for generating documentation from the specification.
type Generator interface {
	// Generate generates documentation from the specification.
	// Generate convert jsight to openapi json or yaml schema
	// Specification will be read from in and print to out.
	Generate(ctx *cli.Context, filepath string, in io.Reader, out io.Writer, sendStatFlag bool, fileSize int64) error
}

var ErrUnsupportedFormat = errors.New("unsupported format")

func New(f format.Format) (Generator, error) {
	switch f {
	case format.FormatHTML:
		return newHTML(), nil
	case format.FormatJSON:
		return newOpenapiJSON(), nil
	case format.FormatYAML:
		return newOpenapiYAML(), nil
	default:
		return nil, ErrUnsupportedFormat
	}
}

type common struct {
	gen func(kit.JApi, io.Writer) error
}

func (c common) Generate(ctx *cli.Context, filepath string, in io.Reader, out io.Writer, sendStatFlag bool, fileSize int64) error {
	j, err := kit.NewJapi(filepath)
	statistics.SendStat(&j, err, sendStatFlag, fileSize, nil)
	if err != nil {
		filepath := "no file"
		if err.File != nil {
			filepath = err.File.Name()
		}
		return fmt.Errorf("JSight API parsing error at `%v` [%v, %v]\nError message: %v\nQuote:\n----> %v",
			filepath,
			err.Location.Line,
			err.Location.Column,
			err.Msg,
			err.Quote)
	}
	return c.gen(j, out)
}

func FormatError(err error, f format.Format) error {
	return fmt.Errorf("convert %s: %w", f, err)
}
