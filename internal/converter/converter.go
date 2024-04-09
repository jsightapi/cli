package converter

import (
	"context"

	"errors"

	"fmt"
	
	"io"

	"github.com/jsightapi/jsight-api-core/kit"
)

type Format string

const (
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

// Converter an abstraction for generating openapi schema from the jsight specification.
type Converter interface {
	// Convert makes openapi schema from the jsight specification.
	// Specification will be read from in and print to out.
	Convert(ctx context.Context, filepath string, in io.Reader, out io.Writer) error
}

type OpenapiFmt struct {
	common
}

var _ Converter = OpenapiFmt{}

var ErrUnsupportedFormat = errors.New("unsupported format")

func New(f Format) (Converter, error) {
	switch f {
	case FormatJSON:
		return newOpenapiJSON(), nil
	case FormatYAML:
		return newOpenapiYAML(), nil
	default:
		return nil, ErrUnsupportedFormat
	}
}

type common struct {
	cnv func(kit.JApi, io.Writer) error
}

func (c common) Convert(_ context.Context, filepath string, in io.Reader, out io.Writer) error {
	j, err := kit.NewJapi(filepath)
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
	return c.cnv(j, out)
}

func FormatError(err error, f Format) error {
	return fmt.Errorf("convert %s: %w", f, err)
}
