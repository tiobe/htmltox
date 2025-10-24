package img

import (
	"context"
	"htmltox/shared"

	"github.com/chromedp/chromedp"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:  "img",
	Usage: "Capture an image of a webpage using a Chromium (based) browser.",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "output", Value: "htmltox.png", Aliases: []string{"o"}, Usage: "Path to output the image file"},
		&cli.StringFlag{Name: "selector", Required: true, Aliases: []string{"S"}, Usage: "HTML selector to define what to create an image for (eg. div.tqi-label)."},
		&cli.Float64Flag{Name: "scale", Value: 3.0, Aliases: []string{"s"}, Usage: "Device scale factor (1.0 = low, 2.0 = normal, 3.0 = high, 4.0 = ultra high)"},
	},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		actionFunc := func(buffer *[]byte) chromedp.Action {
			return chromedp.Screenshot(cmd.String("selector"), buffer)
		}
		return shared.Run(cmd, actionFunc, "image", cmd.Float64("scale"))
	},
}
