package cruda

import (
	"github.com/Oppodelldog/chromedp-test/group"
	"github.com/chromedp/chromedp"
)

func Group(title string, actions ...chromedp.Action) chromedp.Action {
	return group.New(title, actions...)
}

func OpenTestPage(url string) chromedp.Action {
	return prefixError("OpenTestPage",
		group.New("open CRUD Test Page on "+url,
			OpenWebsite(url),
			NavigateToTestPage(),
			WaitForComponent(),
		),
	)
}

func OpenWebsite(url string) chromedp.Action {
	return prefixError("OpenWebsite",
		group.New("open CRUDA sample application",
			chromedp.Navigate(url),
		),
	)
}

func NavigateToTestPage() chromedp.Action {
	return prefixError("NavigateToTestPage",
		Navigate("/test"),
	)
}

func NavigateToAboutPage() chromedp.Action {
	return prefixError("NavigateToAboutPage",
		Navigate("/about"),
	)
}

func NavigateToTestAndToolsPage() chromedp.Action {
	return prefixError("NavigateToTestAndToolsPage",
		Navigate("/test-and-tools"),
	)
}

func Navigate(path string) chromedp.Action {
	return prefixError("Navigate",
		group.New("navigate to test page",
			chromedp.Click("*[data-testid='app-routes'] a[href='"+path+"']", chromedp.ByQuery),
		),
	)
}
