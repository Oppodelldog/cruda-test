package crud

import (
	"context"

	. "github.com/Oppodelldog/cruda-test/cruda" //nolint:golint
	"github.com/chromedp/chromedp"
)

func Case01PageWorks(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		OpenTestPage(url),

		Group("expect an empty list",
			ListItemNum(0),
		),

		Group("expect an empty form",
			FormIDEquals(""),
			FormTextEquals(""),
			FormDoneEquals(false),
		),
	)
}
