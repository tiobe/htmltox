package main

import (
	"context"
	"htmltox/internal/img"
	"htmltox/internal/pdf"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var commands = &cli.Command{
	Name:      "htmltox",
	Version:   "0.1.1",
	Copyright: "(c) 2025 TIOBE Software BV",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "url", Required: true, Aliases: []string{"u"}, Usage: "URL to generate PDF from"},
		&cli.StringFlag{Name: "chromium-path", Required: true, Aliases: []string{"c"}, Usage: "Path to Chrome/Chromium executable"},
		&cli.StringFlag{Name: "window-status", Usage: "Wait for the window to reach a certain status"},
		&cli.StringSliceFlag{Name: "header", Aliases: []string{"H"}, Usage: "Custom HTTP headers, e.g. -H 'Authorization: Basic <token>'"},
		&cli.BoolFlag{Name: "headless", Value: false, Usage: "Run htmltox with a headless browser"},
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
