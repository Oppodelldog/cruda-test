package crud

import (
	"context"

	"github.com/Oppodelldog/cruda-test/cruda"
	"github.com/Oppodelldog/cruda-test/group"
	"github.com/chromedp/chromedp"
)

func Case05DeleteItem(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		group.New("open Test page with fixtures",
			cruda.OpenWebsite(url),
			cruda.InitAdapterFixtures(testAdapterID),
			cruda.NavigateToTestPage(),
			cruda.WaitForComponent(),

			group.New("as a precondition expect 2 items in list",
				cruda.ListItemNum(2), //nolint:gomnd
				cruda.ListItemText(0, "Entry 1"),
				cruda.ListItemText(1, "Entry 2"),
			),
		),

		group.New("delete the second item",
			cruda.ListSelect(1),
			cruda.DeleteItem(),
			cruda.FormShowsSuccess(),
			group.New("expect 1 item in list",
				cruda.ListItemNum(1),
				cruda.ListItemText(0, "Entry 1"),
			),
		),

		group.New("delete the remaining (first) item",
			cruda.ListSelect(0),
			cruda.DeleteItem(),
			cruda.FormShowsSuccess(),
			group.New("expect an empty list",
				cruda.ListItemNum(0),
			),
		),
	)
}

func Case05DeleteItemServerError(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		group.New("open Test page with fixtures",
			cruda.OpenWebsite(url),
			cruda.InitAdapterFixtures(testAdapterID),
			cruda.NavigateToTestPage(),
			cruda.WaitForComponent(),
		),

		cruda.InitAdapterDeleteError(testAdapterID, "Server cannot delete item"),

		group.New("select item 1 in list",
			cruda.ListSelect(0),
			cruda.DeleteItem(),
		),

		cruda.FormShowsError("Server cannot delete item"),
	)
}
