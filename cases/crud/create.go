package crud

import (
	"context"

	. "github.com/Oppodelldog/cruda-test/cruda" //nolint:golint
	"github.com/chromedp/chromedp"
)

func Case02CreateItem(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		OpenTestPage(url),

		Group("add one entry",
			Group("as a precondition expect an empty list",
				ListItemNum(0),
			),

			Group("add Entry 1",
				FormEnterID("1"),
				FormEnterText("Entry 1"),
				FormSetDone(true),
				FormSubmit(),
				FormShowsSuccess(),
			),

			Group("expect Entry 1 to be in List",
				ListItemNum(1),
				ListItemText(0, "Entry 1"),
			),

			Group("expect item 1 data still in form",
				FormIDEquals("1"),
				FormTextEquals("Entry 1"),
				FormDoneEquals(true),
			),
		),

		Group("add a second entry - needs to press new button to clean form",
			Group("press new button, to clean the form",
				NewItem(),
			),

			Group("expect empty form",
				FormIDEquals(""),
				FormTextEquals(""),
				FormDoneEquals(false),
			),

			Group("add Entry 2",
				FormEnterID("2"),
				FormEnterText("Entry 2"),
				FormSetDone(true),
				FormSubmit(),
				FormShowsSuccess(),
			),

			Group("expect 2 items in list",
				ListItemNum(2), //nolint:gomnd
				ListItemText(0, "Entry 1"),
				ListItemText(1, "Entry 2"),
			),

			Group("expect item 2 data still in form",
				FormIDEquals("2"),
				FormTextEquals("Entry 2"),
				FormDoneEquals(true),
			),
		),
	)
}

func Case02CreateItemFormValidation(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		OpenTestPage(url),

		Group("as a precondition expect an empty list",
			ListItemNum(0),
		),

		Group("expect empty form",
			FormIDEquals(""),
			FormTextEquals(""),
			FormDoneEquals(false),
			FormSubmit(),
		),

		Group("expect list to be unchanged",
			ListItemNum(0),
		),
	)
}

func Case02CreateItemServerError(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		OpenTestPage(url),

		Group("as a precondition expect an empty list",
			ListItemNum(0),
		),

		Group("init data adapter to return an error on create",
			InitAdapterCreationError(testAdapterID, "Server cannot create item"),
		),

		Group("add entry 3",
			FormEnterID("3"),
			FormEnterText("Entry 3"),
			FormSetDone(true),
			FormSubmit(),
		),

		Group("expect list to be unchanged",
			ListItemNum(0),
		),

		FormShowsError("Server cannot create item"),
	)
}
