package shared

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

type ActionFunc func(*[]byte) chromedp.Action

type RunArguments struct {
	ChromiumPath string
	Headers      []string
	Headless     bool
	Output       string
	Url          string
	WindowStatus string
}

func Run(args RunArguments, capture ActionFunc, captureType string, scale float64) error {
	opts := append(chromedp.DefaultExecAllocatorOptions[:], // extend the base options
		chromedp.DisableGPU,
		chromedp.ExecPath(args.ChromiumPath),
		chromedp.Flag("headless", args.Headless),
		chromedp.Flag("no-sandbox", args.Headless),
		chromedp.Flag("disable-dev-shm-usage", args.Headless),
	)

	timer := time.Now()
	log.Printf("Opening browser...")

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Set timeout (adjust as needed)
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	chromedp.ListenTarget(ctx, func(ev any) {
		resp, ok := ev.(*network.EventResponseReceived)
		if ok && resp.Response.Status >= 400 {
			log.Printf("Received %d for %s, cancelling", resp.Response.Status, resp.Response.URL)
			cancel() // aborts chromedp.Run()
		}
	})

	// Set headers if provided
	headers, err := parseHeaders(args.Headers)
	if err != nil {
		return err
	}

	var buffer []byte
	tasks := chromedp.Tasks{
		network.Enable(),
		emulation.SetEmulatedMedia().WithFeatures([]*emulation.MediaFeature{
			{Name: "prefers-color-scheme", Value: "light"},
		}),
		emulation.SetDeviceMetricsOverride(1920, 1080, scale, false),
		network.SetExtraHTTPHeaders(headers),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Printf("Browser opened in %.2f seconds", time.Since(timer).Seconds())
			return nil
		}),
		navigateAndWaitForStatus(args.Url, args.WindowStatus),
		capture(&buffer),
	}

	if err := chromedp.Run(ctx, tasks...); err != nil {
		return fmt.Errorf("chromedp run failed: %w", err)
	}

	if err := os.WriteFile(args.Output, buffer, 0644); err != nil {
		return fmt.Errorf("failed to save %s: %w", captureType, err)
	}

	log.Printf("%s saved to %s in %.2f seconds\n", captureType, args.Output, time.Since(timer).Seconds())
	return nil
}

func parseHeaders(cmdHeaders []string) (network.Headers, error) {
	headers := network.Headers{}
	for _, h := range cmdHeaders {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) == 2 {
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		} else {
			return nil, fmt.Errorf("failed to parse header: %s", h)
		}
	}
	return headers, nil
}

// Polls window.status until it matches the given value.
func navigateAndWaitForStatus(url string, status string) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timer := time.Now()

		log.Printf("Navigating to: %s", url)
		if err := chromedp.Run(ctx,
			chromedp.Navigate(url),
			chromedp.WaitReady("body", chromedp.ByQuery),
		); err != nil {
			return fmt.Errorf("failed to navigate to %s: %w", url, err)
		}

		if status != "" {
			log.Printf("Waiting for window.status: %s", status)
			for {
				var windowStatus string
				err := chromedp.Evaluate(`window.status`, &windowStatus).Do(ctx)
				if err == nil && windowStatus == status {
					log.Printf("Page loaded in %.2f seconds", time.Since(timer).Seconds())
					return nil
				}
				select {
				case <-ctx.Done():
					if ctx.Err() == context.Canceled {
						return fmt.Errorf("navigation canceled after %.2f seconds", time.Since(timer).Seconds())
					}
					return fmt.Errorf("timeout waiting for window.status: %s", status)
				case <-time.After(100 * time.Millisecond):
				}
			}
		}
		return nil
	}
}
