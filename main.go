package main

import (
	"bitbucket.org/jsight/cli/internal/converter"

	"fmt"

	"os"
	
	"runtime"

	"github.com/urfave/cli/v2"

	"bitbucket.org/jsight/cli/internal/generator"
)

var (
	Version    = "1.0.0"
	CommitHash = "commit hash not defined"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	app := &cli.App{
		Usage: "is a tool for working with files in the JSight language.",
		Commands: []*cli.Command{
			{
				Name:  "doc",
				Usage: "generate API documentation in various formats",
				Subcommands: []*cli.Command{
					{
						Name:      "html",
						Usage:     "generates documentation in HTML format",
						Flags:     []cli.Flag{},
						ArgsUsage: "<input>",
						Action:    generateDocumentation(generator.FormatHTML),
					}, /*
						{
							Name:      "pdf",
							Usage:     "generates documentation in PDF format",
							ArgsUsage: "<input>",
							Action:    generateDocumentation(generator.FormatPDF),
						},
						{
							Name:      "docx",
							Usage:     "generates documentation in DOCX format",
							ArgsUsage: "<input>",
							Action:    generateDocumentation(generator.FormatDOCX),
						},*/
				},
			},
			{
				Name:  "convert",
				Usage: "converts JSight to other formats",
				Subcommands: []*cli.Command{
					{
						Name:  "openapi",
						Usage: "generate OpenAPI json or yaml from JSight",
						Subcommands: []*cli.Command{
							{
								Name:      "json",
								Usage:     "generate OpenAPI JSON from JSight",
								Flags:     []cli.Flag{},
								ArgsUsage: "<input>",
								Action:    convert(converter.FormatJSON),
							},
							{
								Name:      "yaml",
								Usage:     "generate OpenAPI YAML from JSight",
								Flags:     []cli.Flag{},
								ArgsUsage: "<input>",
								Action:    convert(converter.FormatYAML),
							},
						},
					},
				},
			},
			{
				Name:   "version",
				Usage:  "print tool version",
				Action: printVersion,
			},
		},
		EnableBashCompletion: true,
	}

	return app.Run(os.Args)
}

func parseArgs(ctx *cli.Context, f string) (*os.File, *string, error) {
	aa := ctx.Args()
	if aa.Len() != 1 {
		return nil, nil, cli.ShowCommandHelp(ctx, f)
	}

	specPath := aa.First()
	if _, err := os.Stat(specPath); err != nil {
		return nil, nil, err
	}

	r, err := os.Open(specPath)
	return r, &specPath, err
}

func generateDocumentation(f generator.Format) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		r, specPath, err := parseArgs(ctx, string(f))
		if err != nil {
			return err
		}
		g, err := generator.New(f)
		if err != nil {
			return err
		}
		return g.Generate(ctx.Context, *specPath, r, ctx.App.Writer)
	}
}

func convert(f converter.Format) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		r, specPath, err := parseArgs(ctx, string(f))
		if err != nil {
			return err
		}
		c, err := converter.New(f)
		if err != nil {
			return err
		}
		return c.Convert(ctx.Context, *specPath, r, ctx.App.Writer)
	}
}

func printVersion(ctx *cli.Context) error {
	_, err := fmt.Fprintf(ctx.App.Writer, `Version: %s (%s)
Golang version: %s
`, Version, CommitHash, runtime.Version())
	return err
}
