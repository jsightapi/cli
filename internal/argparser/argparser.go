package argparser

import (
	"errors"

	"github.com/jsightapi/cli/internal/statistics"

	"github.com/jsightapi/cli/internal/format"

	"github.com/jsightapi/cli/internal/generator"

	"github.com/urfave/cli/v2"

	"os"
)

var noInputFileErrorString = "no input file"

func DoWork(f format.Format, ctx *cli.Context) error {
	var err error
	var r *os.File
	var specPath *string
	var sendStatFlag bool
	var fileSize int64
	r, specPath, sendStatFlag, fileSize, err = parseArgs(ctx)
	if err == nil {
		var g generator.Generator
		g, err = generator.New(f)
		if err == nil {
			return g.Generate(ctx, *specPath, r, ctx.App.Writer, sendStatFlag, fileSize)
		}
		_ = cli.ShowCommandHelp(ctx, string(f))
		statistics.SendStat(nil, nil, sendStatFlag, fileSize, err)
		return err
	}
	statistics.SendStat(nil, nil, sendStatFlag, fileSize, err)
	if err.Error() == noInputFileErrorString {
		err = cli.ShowCommandHelp(ctx, string(f))
	}
	return err
}

func parseArgs(ctx *cli.Context) (file *os.File, specPath *string, sendStatFlag bool, fileSize int64, e error) {
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
				var r *os.File
				r, err = os.Open(*specPath)
				return r, specPath, sendStatFlag, fileSize, err
			}
		}
	}
	e = err
	if e == nil {
		e = errors.New(noInputFileErrorString)
	}
	return nil, nil, sendStatFlag, 0, e
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
