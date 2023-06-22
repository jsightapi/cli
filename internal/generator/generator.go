package generator

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/jsightapi/jsight-api-core/kit"
)

type Format string

const (
	FormatHTML Format = "html"
	FormatPDF  Format = "pdf"
	FormatDOCX Format = "docx"
)

// Generator an abstraction for generating documentation from the specification.
type Generator interface {
	// Generate generates documentation from the specification.
	// Specification will be read from in and print to out.
	Generate(ctx context.Context, filepath string, in io.Reader, out io.Writer) error
}

var ErrUnsupportedFormat = errors.New("unsupported format")

func New(f Format) (Generator, error) {
	if f != FormatHTML {
		return nil, ErrUnsupportedFormat
	}
	return newHTML(), nil
}

type common struct {
	gen func(kit.JApi, io.Writer) error
}

func (c common) Generate(_ context.Context, filepath string, in io.Reader, out io.Writer) error {
	j, err := kit.NewJapi(filepath)
	if err != nil {
		return fmt.Errorf("specification error: %w", err)
	}
	return c.gen(j, out)
}
