package cruda

import (
	"github.com/Oppodelldog/cruda-test/group"
	"github.com/chromedp/chromedp"
)

func OpenTestPage(url string) chromedp.Action {
	return prefixError("OpenTestPage",
		group.New("open CRUD Test Page on "+url,
			OpenWebsite(url),
			Navigate("/test"),
			WaitForComponent(),
		),
	)
}

func OpenWebsite(url string) chromedp.Action {
	return prefixError("OpenWebsite",
		group.Simple(
			group.Text("open CRUDA sample application"),
			chromedp.Navigate(url),
		),
	)
}

func NavigateToTestPage() chromedp.Action {
	return prefixError("NavigateToTestPage",
		Navigate("/test"),
	)
}

func Navigate(path string) chromedp.Action {
	return prefixError("Navigate",
		group.Simple(
			group.Text("navigate to test page"),
			chromedp.Click("*[data-testid='app-routes'] a[href='"+path+"']", chromedp.ByQuery),
		),
	)
}
