package crud

import (
	"context"

	"github.com/Oppodelldog/chromedp-test/group"
	"github.com/Oppodelldog/cruda-test/cruda"
	"github.com/chromedp/chromedp"
)

func Case01PageWorks(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		cruda.OpenTestPage(url),

		group.New("expect an empty list",
			cruda.ListItemNum(0),
		),

		group.New("expect an empty form",
			cruda.FormIDEquals(""),
			cruda.FormTextEquals(""),
			cruda.FormDoneEquals(false),
		),
	)
}
