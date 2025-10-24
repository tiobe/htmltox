package main

import (
	"context"
	"htmltox/img"
	"htmltox/pdf"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var commands = &cli.Command{
	Name:      "htmltox",
	Version:   "0.1.0",
	Copyright: "(c) 2025 TIOBE Software BV",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "url", Required: true, Aliases: []string{"u"}, Usage: "URL to generate PDF from"},
		&cli.StringFlag{Name: "chromium-path", Required: true, Aliases: []string{"c"}, Usage: "Path to Chrome/Chromium executable"},
		&cli.StringSliceFlag{Name: "header", Aliases: []string{"H"}, Usage: "Custom HTTP headers, e.g. -H 'Authorization: Basic <token>'"},
	},
	Commands: []*cli.Command{
		pdf.Command,
		img.Command,
	},
}

func main() {
	log.SetFlags(0)
	if err := commands.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
