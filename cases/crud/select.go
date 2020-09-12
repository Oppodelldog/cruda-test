package crud

import (
	"context"

	"github.com/Oppodelldog/chromedp-test/group"
	"github.com/Oppodelldog/cruda-test/cruda"
	"github.com/chromedp/chromedp"
)

func Case03SelectItem(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		group.New("open Test page with fixtures",
			cruda.OpenWebsite(url),
			cruda.InitAdapterFixtures(testAdapterID),
			cruda.NavigateToTestPage(),
			cruda.WaitForComponent(),
		),

		group.New("test selection of Entry 1",
			group.New("select item 1 in list",
				cruda.ListSelect(0),
			),
			group.New("expect item 1 data in form",
				cruda.FormIDEquals("1"),
				cruda.FormTextEquals("Entry 1"),
				cruda.FormDoneEquals(false),
			),
		),

		group.New("test selection of Entry 2",
			group.New("now select item 2 in list",
				cruda.ListSelect(1),
			),
			group.New("expect item 2 data in form",
				cruda.FormIDEquals("2"),
				cruda.FormTextEquals("Entry 2"),
				cruda.FormDoneEquals(true),
			),
		),
	)
}

func Case03SelectItemServerError(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		group.New("open Test page with fixtures",
			cruda.OpenWebsite(url),
			cruda.InitAdapterFixtures(testAdapterID),
			cruda.NavigateToTestPage(),
			cruda.WaitForComponent(),
		),

		cruda.InitAdapterLoadItemError(testAdapterID, "Server cannot load item"),

		group.New("select item 1 in list",
			cruda.ListSelect(0),
		),

		cruda.FormShowsError("Server cannot load item"),
	)
}
