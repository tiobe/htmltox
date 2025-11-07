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
	"github.com/urfave/cli/v3"
)

type ActionFunc func(*[]byte) chromedp.Action

func Run(cmd *cli.Command, capture ActionFunc, captureType string, scale float64) error {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.DisableGPU,
		chromedp.ExecPath(cmd.String("chromium-path")),
		chromedp.Flag("headless", cmd.Bool("headless")),
	}

	timer := time.Now()
	log.Printf("Opening browser...")

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Set timeout (adjust as needed)
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Set headers if provided
	headers, err := parseHeaders(cmd.StringSlice("header"))
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
		navigateAndWaitForStatus(cmd.String("url"), "ready"),
		capture(&buffer),
	}

	if err := chromedp.Run(ctx, tasks...); err != nil {
		return fmt.Errorf("chromedp run failed: %w", err)
	}

	if err := os.WriteFile(cmd.String("output"), buffer, 0644); err != nil {
		return fmt.Errorf("failed to save %s: %w", captureType, err)
	}

	log.Printf("%s saved to %s in %.2f seconds\n", captureType, cmd.String("output"), time.Since(timer).Seconds())
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
func navigateAndWaitForStatus(url string, expected string) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timer := time.Now()

		log.Printf("Navigating to: %s", url)
		if err := chromedp.Run(ctx,
			chromedp.Navigate(url),
			chromedp.WaitReady("body", chromedp.ByQuery),
		); err != nil {
			return fmt.Errorf("failed to navigate to %s: %w", url, err)
		}

		log.Printf("Waiting for window.status: %s", expected)
		for {
			var status string
			err := chromedp.Evaluate(`window.status`, &status).Do(ctx)
			if err == nil && status == expected {
				elapsed := time.Since(timer).Seconds()
				log.Printf("Page loaded in %.2f seconds", elapsed)
				return nil
			}
			select {
			case <-ctx.Done():
				return fmt.Errorf("timeout waiting for window.status: %s", expected)
			case <-time.After(100 * time.Millisecond):
			}
		}
	}
}
