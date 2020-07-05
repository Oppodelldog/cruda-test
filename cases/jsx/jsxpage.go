package jsx

import (
	"context"

	"github.com/Oppodelldog/cruda-test/cruda"
	"github.com/Oppodelldog/cruda-test/group"
	"github.com/chromedp/chromedp"
)

func Case01CallJSXPage(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		group.New("Open About CRUDA Page",
			cruda.OpenWebsite(url),
			cruda.Navigate("/about"),
			cruda.WaitForJSXPageComponent(),
		),

		group.New("assert headline",
			cruda.Equals(`document.querySelector("h1").innerHTML`, "About CRUDA"),
		),
	)
}
