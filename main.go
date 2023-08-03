package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"

	"bitbucket.org/jsight/cli/internal/generator"
)

var (
	Version    = "development"
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
					},/*
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
			},/*
			{
				Name:   "convert",
				Usage:  "converting JSight to OpenAPI and back",
				Action: convert,
			},*/
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

func generateDocumentation(f generator.Format) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		aa := ctx.Args()
		if aa.Len() != 1 {
			return cli.ShowCommandHelp(ctx, string(f))
		}

		specPath := aa.First()
		if _, err := os.Stat(specPath); err != nil {
			return err
		}

		r, err := os.Open(specPath)
		if err != nil {
			return err
		}

		g, err := generator.New(f)
		if err != nil {
			return err
		}

		return g.Generate(ctx.Context, specPath, r, ctx.App.Writer)
	}
}

/* func convert(*cli.Context) error {
	return errors.New("not implemented yet")
}*/

func printVersion(ctx *cli.Context) error {
	_, err := fmt.Fprintf(ctx.App.Writer, `Version: %s (%s)
Golang version: %s
`, Version, CommitHash, runtime.Version())
	return err
}
