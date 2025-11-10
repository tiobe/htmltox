package pdf

import (
	"context"
	"htmltox/internal/shared"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:  "pdf",
	Usage: "Convert a webpage to pdf using a Chromium (based) browser.",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "output", Value: "htmltox.pdf", Aliases: []string{"o"}, Usage: "Path to output PDF file"},
		&cli.StringFlag{Name: "page-size", Value: "A4", Aliases: []string{"s"}, Usage: "Page size (A4, Letter, etc.)"},
		&cli.StringFlag{Name: "footer", Usage: "Add a custom string to the left side of the footer"},
		&cli.BoolFlag{Name: "page-numbers", Value: true, Aliases: []string{"n"}, Usage: "Include page numbers in the footer (default is true)"},
	},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		runArgs := shared.RunArguments{
			ChromiumPath: cmd.String("chromium-path"),
			Headers:      cmd.StringSlice("header"),
			Headless:     cmd.Bool("headless"),
			Output:       cmd.String("output"),
			Url:          cmd.String("url"),
			WindowStatus: cmd.String("window-status"),
		}

		actionFunc := func(buffer *[]byte) chromedp.Action {
			return printToPDF(buffer, cmd.String("page-size"), cmd.String("footer"), cmd.Bool("page-numbers"))
		}

		return shared.Run(runArgs, actionFunc, "PDF", 1.0)
	},
}

// Returns a chromedp action that generates the PDF
func printToPDF(buf *[]byte, size, footer string, pageNumbers bool) chromedp.ActionFunc {
	paperSize := getPaperSize(size)
	return func(ctx context.Context) error {
		var err error
		*buf, _, err = page.PrintToPDF().
			WithPaperWidth(paperSize.Width).
			WithPaperHeight(paperSize.Height).
			WithMarginTop(0.5).
			WithMarginBottom(0.5).
			WithMarginLeft(0.5).
			WithMarginRight(0.5).
			WithDisplayHeaderFooter(true).
			WithHeaderTemplate(HeaderHtml()).
			WithFooterTemplate(FooterHtml(footer, pageNumbers)).
			Do(ctx)
		return err
	}
}
