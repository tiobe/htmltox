package pdf

import (
	"context"
	"htmltox/shared"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:  "pdf",
	Usage: "Convert a webpage to pdf using a Chromium (based) browser.",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "url", Required: true, Aliases: []string{"u"}, Usage: "URL to generate PDF from"},
		&cli.StringFlag{Name: "chromiumPath", Required: true, Aliases: []string{"c"}, Usage: "Path to Chrome/Chromium executable"},
		&cli.StringFlag{Name: "output", Value: "htmltox.pdf", Aliases: []string{"o"}, Usage: "Path to output PDF file"},
		&cli.StringFlag{Name: "pageSize", Value: "A4", Aliases: []string{"s"}, Usage: "Page size (A4, Letter, etc.)"},
		&cli.StringFlag{Name: "authHeader", Usage: "Authorization header to use for the requests"},
		&cli.StringFlag{Name: "footer", Usage: ""},
		&cli.BoolFlag{Name: "pageNumbers", Value: true, Usage: "Add page numbers to the footer (default is true)"},
	},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		actionFunc := func(buffer *[]byte) chromedp.Action {
			return printToPDF(buffer, cmd.String("pageSize"), cmd.String("footer"), cmd.Bool("pageNumbers"))
		}
		return shared.Run(cmd, actionFunc, "PDF", 1.0)
	},
}

// Returns a chromedp action that generates the PDF
func printToPDF(buf *[]byte, size, footer string, pageNumbers bool) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		var err error
		*buf, _, err = page.PrintToPDF().
			WithPaperWidth(getPaperWidth(size)).
			WithPaperHeight(getPaperHeight(size)).
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

// Simple paper size mapping
func getPaperWidth(size string) float64 {
	switch size {
	case "Letter":
		return 8.5
	case "Legal":
		return 8.5
	default: // A4
		return 8.27
	}
}

func getPaperHeight(size string) float64 {
	switch size {
	case "Letter":
		return 11.0
	case "Legal":
		return 14.0
	default: // A4
		return 11.7
	}
}
