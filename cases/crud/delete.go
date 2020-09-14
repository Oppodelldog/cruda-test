package crud

import (
	"context"

	. "github.com/Oppodelldog/cruda-test/cruda" //nolint:golint
	"github.com/chromedp/chromedp"
)

func Case05DeleteItem(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		Group("open Test page with fixtures",
			OpenWebsite(url),
			InitAdapterFixtures(testAdapterID),
			NavigateToTestPage(),
			WaitForComponent(),

			Group("as a precondition expect 2 items in list",
				ListItemNum(2), //nolint:gomnd
				ListItemText(0, "Entry 1"),
				ListItemText(1, "Entry 2"),
			),
		),

		Group("delete the second item",
			ListSelect(1),
			DeleteItem(),
			FormShowsSuccess(),
			Group("expect 1 item in list",
				ListItemNum(1),
				ListItemText(0, "Entry 1"),
			),
		),

		Group("delete the remaining (first) item",
			ListSelect(0),
			DeleteItem(),
			FormShowsSuccess(),
			Group("expect an empty list",
				ListItemNum(0),
			),
		),
	)
}

func Case05DeleteItemServerError(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		Group("open Test page with fixtures",
			OpenWebsite(url),
			InitAdapterFixtures(testAdapterID),
			NavigateToTestPage(),
			WaitForComponent(),
		),

		InitAdapterDeleteError(testAdapterID, "Server cannot delete item"),

		Group("select item 1 in list",
			ListSelect(0),
			DeleteItem(),
		),

		FormShowsError("Server cannot delete item"),
	)
}
