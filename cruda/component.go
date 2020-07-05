package cruda

import (
	"github.com/Oppodelldog/cruda-test/group"
	"github.com/chromedp/chromedp"
)

func WaitForComponent() chromedp.Action {
	return prefixError("WaitForComponent",
		group.Simple(
			group.Text("wait for CRUDAComponent to show"),
			chromedp.WaitVisible("*[data-testid='CRUDComponent']", chromedp.ByQuery),
		),
	)
}

func WaitForJSXPageComponent() chromedp.Action {
	return prefixError("WaitForJSXPageComponent",
		group.Simple(
			group.Text("wait for JSX Element that defines a page"),
			chromedp.WaitVisible("*[data-testid='jsx-page']", chromedp.ByQuery),
		),
	)
}
