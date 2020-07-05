package runner

import (
	"context"
	"fmt"

	"github.com/Oppodelldog/cruda-test/cruda"
	"github.com/chromedp/chromedp"
)

func takeFailureScreenshot(ctx context.Context, testID string, err error) {
	if err != nil {
		var fileName = "FAIL-" + testID + ".png"

		errScreenShot := chromedp.Run(ctx, cruda.Screenshot(fileName))
		if errScreenShot != nil {
			fmt.Printf("COULD NOT TAKE SCREENSHOT: %v", errScreenShot)
		}
	}
}
