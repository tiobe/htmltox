package pdf

import (
	"fmt"
	"time"
)

func HeaderHtml() string {
	header := `<div style="width:100%; font-size:8pt; padding:0 20px; box-sizing:border-box; display:flex; justify-content:space-between; color:#555;">`
	header += fmt.Sprintf(`<span class="header-left">%s</span>`, time.Now().Format("2006-01-02 15:04:05"))
	header += `<span class="header-right"></span>`
	header += `</div>`
	return header
}

func FooterHtml(footerText string, pageNumbers bool) string {
	footer := `<div style="width:100%; font-size:8pt; padding:0 20px; box-sizing:border-box; display:flex; justify-content:space-between; color:#555;">`
	if footerText == "" {
		footerText = fmt.Sprintf("(c) %v TIOBE Software BV", time.Now().Year())
	}
	footer += fmt.Sprintf(`<span class="footer-left">%s</span>`, footerText)
	if pageNumbers {
		footer += `<span class="footer-right">Page <span class="pageNumber"></span> / <span class="totalPages"></span></span>`
	}
	footer += `</div>`
	return footer
}
