package crud

import (
	"context"

	"github.com/Oppodelldog/chromedp-test/group"
	"github.com/Oppodelldog/cruda-test/cruda"
	"github.com/chromedp/chromedp"
)

func Case02CreateItem(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		cruda.OpenTestPage(url),

		group.New("add one entry",
			group.New("as a precondition expect an empty list",
				cruda.ListItemNum(0),
			),

			group.New("add Entry 1",
				cruda.FormEnterID("1"),
				cruda.FormEnterText("Entry 1"),
				cruda.FormSetDone(true),
				cruda.FormSubmit(),
				cruda.FormShowsSuccess(),
			),

			group.New("expect Entry 1 to be in List",
				cruda.ListItemNum(1),
				cruda.ListItemText(0, "Entry 1"),
			),

			group.New("expect item 1 data still in form",
				cruda.FormIDEquals("1"),
				cruda.FormTextEquals("Entry 1"),
				cruda.FormDoneEquals(true),
			),
		),

		group.New("add a second entry - needs to press new button to clean form",
			group.New("press new button, to clean the form",
				cruda.NewItem(),
			),

			group.New("expect empty form",
				cruda.FormIDEquals(""),
				cruda.FormTextEquals(""),
				cruda.FormDoneEquals(false),
			),

			group.New("add Entry 2",
				cruda.FormEnterID("2"),
				cruda.FormEnterText("Entry 2"),
				cruda.FormSetDone(true),
				cruda.FormSubmit(),
				cruda.FormShowsSuccess(),
			),

			group.New("expect 2 items in list",
				cruda.ListItemNum(2), //nolint:gomnd
				cruda.ListItemText(0, "Entry 1"),
				cruda.ListItemText(1, "Entry 2"),
			),

			group.New("expect item 2 data still in form",
				cruda.FormIDEquals("2"),
				cruda.FormTextEquals("Entry 2"),
				cruda.FormDoneEquals(true),
			),
		),
	)
}

func Case02CreateItemFormValidation(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		cruda.OpenTestPage(url),

		group.New("as a precondition expect an empty list",
			cruda.ListItemNum(0),
		),

		group.New("expect empty form",
			cruda.FormIDEquals(""),
			cruda.FormTextEquals(""),
			cruda.FormDoneEquals(false),
			cruda.FormSubmit(),
		),

		group.New("expect list to be unchanged",
			cruda.ListItemNum(0),
		),
	)
}

func Case02CreateItemServerError(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		cruda.OpenTestPage(url),

		group.New("as a precondition expect an empty list",
			cruda.ListItemNum(0),
		),

		group.New("init data adapter to return an error on create",
			cruda.InitAdapterCreationError(testAdapterID, "Server cannot create item"),
		),

		group.New("add entry 3",
			cruda.FormEnterID("3"),
			cruda.FormEnterText("Entry 3"),
			cruda.FormSetDone(true),
			cruda.FormSubmit(),
		),

		group.New("expect list to be unchanged",
			cruda.ListItemNum(0),
		),

		cruda.FormShowsError("Server cannot create item"),
	)
}
