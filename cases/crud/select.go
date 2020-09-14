package crud

import (
	"context"

	. "github.com/Oppodelldog/cruda-test/cruda" //nolint:golint
	"github.com/chromedp/chromedp"
)

func Case03SelectItem(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		Group("open Test page with fixtures",
			OpenWebsite(url),
			InitAdapterFixtures(testAdapterID),
			NavigateToTestPage(),
			WaitForComponent(),
		),

		Group("test selection of Entry 1",
			Group("select item 1 in list",
				ListSelect(0),
			),
			Group("expect item 1 data in form",
				FormIDEquals("1"),
				FormTextEquals("Entry 1"),
				FormDoneEquals(false),
			),
		),

		Group("test selection of Entry 2",
			Group("now select item 2 in list",
				ListSelect(1),
			),
			Group("expect item 2 data in form",
				FormIDEquals("2"),
				FormTextEquals("Entry 2"),
				FormDoneEquals(true),
			),
		),
	)
}

func Case03SelectItemServerError(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		Group("open Test page with fixtures",
			OpenWebsite(url),
			InitAdapterFixtures(testAdapterID),
			NavigateToTestPage(),
			WaitForComponent(),
		),

		InitAdapterLoadItemError(testAdapterID, "Server cannot load item"),

		Group("select item 1 in list",
			ListSelect(0),
		),

		FormShowsError("Server cannot load item"),
	)
}
