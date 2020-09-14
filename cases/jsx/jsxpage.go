package jsx

import (
	"context"

	. "github.com/Oppodelldog/cruda-test/cruda" //nolint:golint
	"github.com/chromedp/chromedp"
)

func Case01CallJSXPage(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		Group("Open About CRUDA Page",
			OpenWebsite(url),
			NavigateToAboutPage(),
			WaitForJSXPageComponent(),
		),

		Group("assert headline",
			Equals(`document.querySelector("h1").innerHTML`, "About CRUDA"),
		),
	)
}
