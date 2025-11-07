# htmltox

Capture an image or pdf of a webpage using a Chromium (based) browser.

## PDF generation

```shell
NAME:
   htmltox pdf - Convert a webpage to pdf using a Chromium (based) browser.

USAGE:
   htmltox pdf [options]

OPTIONS:
   --output string, -o string     Path to output PDF file (default: "htmltox.pdf")
   --page-size string, -s string  Page size (A4, Letter, etc.) (default: "A4")
   --footer string                Add a custom string to the left side of the footer
   --page-numbers, -n             Include page numbers in the footer (default is true)
   --help, -h                     show help

GLOBAL OPTIONS:
   --url string, -u string                                    URL to generate PDF from
   --chromium-path string, -c string                          Path to Chrome/Chromium executable
   --header string, -H string [ --header string, -H string ]  Custom HTTP headers, e.g. -H 'Authorization: Basic <token>'
   --headless                                                 Run htmltox with a headless browser.
```

## Image generation

```shell
NAME:
   htmltox img - Capture an image of a webpage using a Chromium (based) browser.

USAGE:
   htmltox img [options]

OPTIONS:
   --output string, -o string    Path to output the image file (default: "htmltox.png")
   --selector string, -S string  HTML selector to define what to create an image for (eg. div.tqi-label).
   --scale float, -s float       Device scale factor (default: 1)
   --help, -h                    show help

GLOBAL OPTIONS:
   --url string, -u string                                    URL to generate PDF from
   --chromium-path string, -c string                          Path to Chrome/Chromium executable
   --header string, -H string [ --header string, -H string ]  Custom HTTP headers, e.g. -H 'Authorization: Basic <token>'
   --headless                                                 Run htmltox with a headless browser.
```
