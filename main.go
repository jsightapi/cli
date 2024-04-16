package main

import (
	"fmt"

	"github.com/jsightapi/cli/internal/argparser"

	"github.com/jsightapi/cli/internal/format"

	"os"

	"runtime"

	"github.com/urfave/cli/v2"
)

var (
	Version = "1.2.0"
)

var statDisclaimer = "   By default, the application sends anonymous usage and error statistics to JSight to help improve the product.\n   Use the -s option to stop sending the statistics."

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	app := &cli.App{
		Name:  "jsight",
		Usage: fmt.Sprintf("is a tool for working with files in the JSight language.\n\n%s", statDisclaimer),
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
						Action:    makeResult(format.FormatHTML),
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
						Usage: "converts JSight to OpenAPI JSON or YAML",
						Subcommands: []*cli.Command{
							{
								Name:      "json",
								Usage:     "converts JSight to OpenAPI JSON",
								Flags:     []cli.Flag{},
								ArgsUsage: "<input>",
								Action:    makeResult(format.FormatJSON),
							},
							{
								Name:      "yaml",
								Usage:     "converts JSight to OpenAPI YAML",
								Flags:     []cli.Flag{},
								ArgsUsage: "<input>",
								Action:    makeResult(format.FormatYAML),
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
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "s",
				Usage:       "disables the sending of anonymous usage and error statistics to JSight",
				DefaultText: "enabled",
			},
		},
	}

	return app.Run(os.Args)
}

func makeResult(f format.Format) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return argparser.DoWork(ctx, f)
	}
}

func printVersion(ctx *cli.Context) error {
	_, err := fmt.Fprintf(ctx.App.Writer, `Version: %s
Golang version: %s
--
%s
`, Version, runtime.Version(), statDisclaimer)
	return err
}
