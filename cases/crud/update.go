package crud

import (
	"context"

	. "github.com/Oppodelldog/cruda-test/cruda" //nolint:golint
	"github.com/chromedp/chromedp"
)

func Case04UpdateItem(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		Group("open Test page with fixtures",
			OpenWebsite(url),
			InitAdapterFixtures(testAdapterID),
			NavigateToTestPage(),
			WaitForComponent(),
		),
		Group("now select item 1 in list",
			ListSelect(0),
		),

		Group("change items text",
			FormTextEquals("Entry 1"),
			FormEnterText(" - TEST"),
			FormSubmit(),
		),

		Group("expect item to be updates in list",
			ListItemText(0, "Entry 1 - TEST"),
		),

		Group("now switch selection and ensure new text after reloading",
			ListSelect(1),
			FormTextEquals("Entry 2"),
			ListSelect(0),
			FormTextEquals("Entry 1 - TEST"),
		),
	)
}

func Case04UpdateItemServerError(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		Group("open Test page with fixtures",
			OpenWebsite(url),
			InitAdapterFixtures(testAdapterID),
			NavigateToTestPage(),
			WaitForComponent(),
		),

		InitAdapterUpdateError(testAdapterID, "Server cannot update item"),

		Group("change item 1 text",
			ListSelect(0),
			FormTextEquals("Entry 1"),
			FormEnterText(" - TEST"),
			FormSubmit(),
		),

		FormShowsError("Server cannot update item"),
	)
}
