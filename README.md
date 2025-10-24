# htmltox

Capture an image or pdf of a webpage using a Chromium (based) browser.

## PDF generation

```shell
NAME:
   htmltox pdf - Convert a webpage to pdf using a Chromium (based) browser.

USAGE:
   htmltox pdf [options]

OPTIONS:
   --url string, -u string           URL to generate PDF from
   --chromiumPath string, -c string  Path to Chrome/Chromium executable
   --output string, -o string        Path to output PDF file (default: "htmltox.pdf")
   --pageSize string, -s string      Page size (A4, Letter, etc.) (default: "A4")
   --authHeader string               Authorization header to use for the requests
   --footer string                   Add a custom string to the left side of the footer
   --pageNumbers                     Add page numbers to the footer (default is true)
   --help, -h                        show help
```

## Image generation

```shell
NAME:
   htmltox img - Capture an image of a webpage using a Chromium (based) browser.

USAGE:
   htmltox img [options]

OPTIONS:
   --url string, -u string           URL to generate PDF from
   --chromiumPath string, -c string  Path to Chrome/Chromium executable
   --selector string, -s string      HTML selector to define what to create an image for (eg. div.tqi-label).
   --output string, -o string        Path to output the image file (default: "htmltox.png")
   --authHeader string               Authorization header to use for the requests
   --scale float                     Device scale factor (1.0 = low, 2.0 = normal, 3.0 = high, 4.0 = ultra high) (default: 3)
   --help, -h                        show help
```
