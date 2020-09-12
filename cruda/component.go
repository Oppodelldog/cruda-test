package cruda

import (
	"github.com/Oppodelldog/chromedp-test/group"
	"github.com/chromedp/chromedp"
)

func WaitForComponent() chromedp.Action {
	return prefixError("WaitForComponent",
		group.New("wait for CRUDAComponent to show",
			chromedp.WaitVisible("*[data-testid='CRUDComponent']", chromedp.ByQuery),
		),
	)
}

func WaitForJSXPageComponent() chromedp.Action {
	return prefixError("WaitForJSXPageComponent",
		group.New("wait for JSX Element that defines a page",
			chromedp.WaitVisible("*[data-testid='jsx-page']", chromedp.ByQuery),
		),
	)
}

func WaitForToolsContainer() chromedp.Action {
	return prefixError("WaitForToolsContainer",
		group.New("wait for tools container that holds the custom tools",
			chromedp.WaitVisible("*[data-testid='crud-and-tools-container']", chromedp.ByQuery),
		),
	)
}
