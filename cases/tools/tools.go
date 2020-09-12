package tools

import (
	"context"

	"github.com/Oppodelldog/chromedp-test/group"
	"github.com/Oppodelldog/cruda-test/cruda"
	"github.com/chromedp/chromedp"
)

const testAdapterID = "testandtools"

func Case01ToolsExtensionReceiveCreateEvent(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		group.New("Open Test Tools page",
			cruda.OpenWebsite(url),
			cruda.NavigateToTestAndToolsPage(),
			cruda.WaitForToolsContainer(),
		),

		group.New("created item text will show up",
			group.New("initially the tools samples pages field showing the create item is empty",
				cruda.ToolsItemCreatedFieldShows(""),
			),
			group.New("add an item and select it",
				cruda.FormEnterText("Tools-Test"),
				cruda.FormEnterID("id-awesome"),
				cruda.FormSubmit(),
			),
			group.New("the created item show up in the demo field",
				cruda.ToolsItemCreatedFieldShows("id: id-awesome, text: Tools-Test"),
			),
		),
	)
}

func Case02ToolsExtensionReceiveSelectEvent(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		group.New("Open Test Tools page",
			cruda.OpenWebsite(url),
			cruda.NavigateToTestAndToolsPage(),
			cruda.WaitForToolsContainer(),
		),

		group.New("selected item text will show up",
			group.New("initially the tools samples pages field showing the selected item is empty",
				cruda.ToolsItemSelectedFieldShows(""),
			),
			group.New("add an item and select it",
				cruda.FormEnterText("Tools-Test"),
				cruda.FormEnterID("id-awesome"),
				cruda.FormSubmit(),
				cruda.ListSelect(0),
			),
			group.New("the selected item show up in the demo field",
				cruda.ToolsItemSelectedFieldShows("id: id-awesome, text: Tools-Test"),
			),
		),
	)
}

func Case03ToolsExtensionReceiveUpdateEvent(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		group.New("Open Test Tools page",
			cruda.OpenWebsite(url),
			cruda.InitAdapterFixtures("testandtools"),
			cruda.NavigateToTestAndToolsPage(),
			cruda.WaitForToolsContainer(),
		),

		group.New("updated item text will show up",
			group.New("initially the tools samples pages field showing the updated item is empty",
				cruda.ToolsItemUpdateFieldShows(""),
			),
			group.New("change an item",
				cruda.ListSelect(0),
				cruda.FormEnterText(" and some additional text"),
				cruda.FormSubmit(),
			),
			group.New("the updated item show up in the demo field",
				cruda.ToolsItemUpdateFieldShows("id: 1, text: Entry 1 and some additional text"),
			),
		),
	)
}

func Case04ToolsExtensionReceiveDeleteEvent(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		group.New("Open Test Tools page",
			cruda.OpenWebsite(url),
			cruda.InitAdapterFixtures(testAdapterID),
			cruda.NavigateToTestAndToolsPage(),
			cruda.WaitForToolsContainer(),
		),

		group.New("delete item text will show up",
			group.New("initially the tools samples pages field showing the delete item is empty",
				cruda.ToolsItemDeleteFieldShows(""),
			),
			group.New("change an item",
				cruda.ListSelect(0),
				cruda.DeleteItem(),
			),
			group.New("the ddelete item show up in the demo field",
				cruda.ToolsItemDeleteFieldShows("id: 1, text: Entry 1"),
			),
		),
	)
}
