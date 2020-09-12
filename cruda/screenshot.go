package cruda

import (
	"context"
	"io/ioutil"
	"log"
	"math"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

const quality = 90

func Screenshot(targetFile string) chromedp.Action {
	return prefixError("Screenshot",
		Action{
			targetFile: targetFile,
		},
	)
}

type Action struct {
	targetFile string
}

func (a Action) Do(ctx context.Context) error {
	take(ctx, a.targetFile)

	return nil
}

func take(ctx context.Context, targetFile string) {
	var buf []byte

	if err := chromedp.Run(ctx, fullScreenshot(quality, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(targetFile, buf, 0600); err != nil {
		log.Fatal(err)
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func fullScreenshot(quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}

			return nil
		}),
	}
}
