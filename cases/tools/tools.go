package tools

import (
	"context"

	. "github.com/Oppodelldog/cruda-test/cruda" //nolint:golint
	"github.com/chromedp/chromedp"
)

const testAdapterID = "testandtools"

func Case01ToolsExtensionReceiveCreateEvent(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		Group("Open Test Tools page",
			OpenWebsite(url),
			NavigateToTestAndToolsPage(),
			WaitForToolsContainer(),
		),

		Group("created item text will show up",
			Group("initially the tools samples pages field showing the create item is empty",
				ToolsItemCreatedFieldShows(""),
			),
			Group("add an item and select it",
				FormEnterText("Tools-Test"),
				FormEnterID("id-awesome"),
				FormSubmit(),
			),
			Group("the created item show up in the demo field",
				ToolsItemCreatedFieldShows("id: id-awesome, text: Tools-Test"),
			),
		),
	)
}

func Case02ToolsExtensionReceiveSelectEvent(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		Group("Open Test Tools page",
			OpenWebsite(url),
			NavigateToTestAndToolsPage(),
			WaitForToolsContainer(),
		),

		Group("selected item text will show up",
			Group("initially the tools samples pages field showing the selected item is empty",
				ToolsItemSelectedFieldShows(""),
			),
			Group("add an item and select it",
				FormEnterText("Tools-Test"),
				FormEnterID("id-awesome"),
				FormSubmit(),
				ListSelect(0),
			),
			Group("the selected item show up in the demo field",
				ToolsItemSelectedFieldShows("id: id-awesome, text: Tools-Test"),
			),
		),
	)
}

func Case03ToolsExtensionReceiveUpdateEvent(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		Group("Open Test Tools page",
			OpenWebsite(url),
			InitAdapterFixtures("testandtools"),
			NavigateToTestAndToolsPage(),
			WaitForToolsContainer(),
		),

		Group("updated item text will show up",
			Group("initially the tools samples pages field showing the updated item is empty",
				ToolsItemUpdateFieldShows(""),
			),
			Group("change an item",
				ListSelect(0),
				FormEnterText(" and some additional text"),
				FormSubmit(),
			),
			Group("the updated item show up in the demo field",
				ToolsItemUpdateFieldShows("id: 1, text: Entry 1 and some additional text"),
			),
		),
	)
}

func Case04ToolsExtensionReceiveDeleteEvent(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		Group("Open Test Tools page",
			OpenWebsite(url),
			InitAdapterFixtures(testAdapterID),
			NavigateToTestAndToolsPage(),
			WaitForToolsContainer(),
		),

		Group("delete item text will show up",
			Group("initially the tools samples pages field showing the delete item is empty",
				ToolsItemDeleteFieldShows(""),
			),
			Group("change an item",
				ListSelect(0),
				DeleteItem(),
			),
			Group("the delete item show up in the demo field",
				ToolsItemDeleteFieldShows("id: 1, text: Entry 1"),
			),
		),
	)
}
