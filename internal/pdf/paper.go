package pdf

import "strings"

type PaperSize struct{ Width, Height float64 }

var paperSizes = map[string]PaperSize{
	"a3":     {Width: 11.7, Height: 16.5},
	"a4":     {Width: 8.27, Height: 11.7},
	"a5":     {Width: 5.83, Height: 8.27},
	"letter": {Width: 8.5, Height: 11.0},
	"legal":  {Width: 8.5, Height: 14.0},
}

// Simple paper size mapping
func getPaperSize(size string) PaperSize {
	key := strings.ToLower(strings.TrimSpace(size))
	ps, ok := paperSizes[key]
	if !ok {
		return paperSizes["a4"]
	}
	return ps
}
