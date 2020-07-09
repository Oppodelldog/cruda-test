package crud

import (
	"context"

	"github.com/Oppodelldog/cruda-test/cruda"
	"github.com/Oppodelldog/cruda-test/group"
	"github.com/chromedp/chromedp"
)

func Case04UpdateItem(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		group.New("open Test page with fixtures",
			cruda.OpenWebsite(url),
			cruda.InitAdapterFixtures(testAdapterID),
			cruda.NavigateToTestPage(),
			cruda.WaitForComponent(),
		),
		group.New("now select item 1 in list",
			cruda.ListSelect(0),
		),

		group.New("change items text",
			cruda.FormTextEquals("Entry 1"),
			cruda.FormEnterText(" - TEST"),
			cruda.FormSubmit(),
		),

		group.New("expect item to be updates in list",
			cruda.ListItemText(0, "Entry 1 - TEST"),
		),

		group.New("now switch selection and ensure new text after reloading",
			cruda.ListSelect(1),
			cruda.FormTextEquals("Entry 2"),
			cruda.ListSelect(0),
			cruda.FormTextEquals("Entry 1 - TEST"),
		),
	)
}

func Case04UpdateItemServerError(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		group.New("open Test page with fixtures",
			cruda.OpenWebsite(url),
			cruda.InitAdapterFixtures(testAdapterID),
			cruda.NavigateToTestPage(),
			cruda.WaitForComponent(),
		),

		cruda.InitAdapterUpdateError(testAdapterID, "Server cannot update item"),

		group.New("change item 1 text",
			cruda.ListSelect(0),
			cruda.FormTextEquals("Entry 1"),
			cruda.FormEnterText(" - TEST"),
			cruda.FormSubmit(),
		),

		cruda.FormShowsError("Server cannot update item"),
	)
}
