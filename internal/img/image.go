package img

import (
	"context"
	"htmltox/internal/shared"

	"github.com/chromedp/chromedp"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:  "img",
	Usage: "Capture an image of a webpage using a Chromium (based) browser.",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "output", Value: "htmltox.png", Aliases: []string{"o"}, Usage: "Path to output the image file"},
		&cli.StringFlag{Name: "selector", Aliases: []string{"S"}, Usage: "HTML selector to define what to create an image for (eg. div.tqi-label)."},
		&cli.Float64Flag{Name: "scale", Value: 1.0, Aliases: []string{"s"}, Usage: "Device scale factor"},
	},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		actionFunc := func(buffer *[]byte) chromedp.Action {
			selector := cmd.String("selector")
			if selector != "" {
				return chromedp.Screenshot(selector, buffer)
			}
			return chromedp.FullScreenshot(buffer, 100)
		}
		runArgs := shared.RunArguments{
			ChromiumPath: cmd.String("chromium-path"),
			Headers:      cmd.StringSlice(""),
			Headless:     cmd.Bool("headless"),
			Output:       cmd.String("output"),
			Url:          cmd.String("url"),
			WindowStatus: cmd.String("window-status"),
		}
		return shared.Run(runArgs, actionFunc, "image", cmd.Float64("scale"))
	},
}
