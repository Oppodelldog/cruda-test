package cruda

import (
	"strconv"

	"github.com/chromedp/chromedp"
)

func ListSelect(itemIndex int) chromedp.Action {
	return prefixError("ListSelect", chromedp.Click(".List li:nth-child("+strconv.Itoa(itemIndex+1)+")", chromedp.ByQuery))
}

func ListItemNum(expected int) chromedp.Action {
	return prefixError("ListItemNum", EqualsAction{
		Expression: "document.querySelectorAll('.List li').length",
		Expected:   expected,
	})
}

func ListItemText(itemIndex int, expected string) chromedp.Action {
	return prefixError("ListItemText", EqualsAction{
		Expression: "document.querySelectorAll('.List li')[" + strconv.Itoa(itemIndex) + "].innerHTML",
		Expected:   expected,
	})
}
