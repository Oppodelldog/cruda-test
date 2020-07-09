package cruda

import "github.com/chromedp/chromedp"

func ToolsItemSelectedFieldShows(v string) chromedp.Action {
	return prefixError(
		"ToolsItemSelectedFieldShows",
		Equals(`document.querySelector("*[data-testid='tools-item-selected']").innerHTML`, v))
}

func ToolsItemCreatedFieldShows(v string) chromedp.Action {
	return prefixError(
		"ToolsItemCreatedFieldShows",
		Equals(`document.querySelector("*[data-testid='tools-item-created']").innerHTML`, v))
}

func ToolsItemUpdateFieldShows(v string) chromedp.Action {
	return prefixError(
		"ToolsItemUpdateFieldShows",
		Equals(`document.querySelector("*[data-testid='tools-item-updated']").innerHTML`, v))
}

func ToolsItemDeleteFieldShows(v string) chromedp.Action {
	return prefixError(
		"ToolsItemDeleteFieldShows",
		Equals(`document.querySelector("*[data-testid='tools-item-deleted']").innerHTML`, v))
}
