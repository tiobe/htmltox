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
	Commands: []*cli.Command{
		pdf.Command,
		img.Command,
	},
}

func main() {
	if err := commands.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
