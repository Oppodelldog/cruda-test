package cruda

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func NewItem() chromedp.Action {
	return prefixError(
		"NewItem",
		chromedp.Click("*[data-testid='bt-new-item']", chromedp.ByQuery))
}
func DeleteItem() chromedp.Action {
	return prefixError(
		"DeleteItem",
		chromedp.Click("*[data-testid='bt-delete-item']", chromedp.ByQuery))
}

func FormShowsSuccess() chromedp.Action {
	return prefixError(
		"FormShowsSuccess",
		Equals(`document.querySelector("*[data-testid='api-result-success']").innerHTML`, "ok"))
}

func FormShowsError(errorValue string) chromedp.Action {
	return prefixError(
		"FormShowsError",
		Equals(`document.querySelector("*[data-testid='api-result-error']").innerHTML`, errorValue))
}

func FormIDEquals(expected string) chromedp.Action {
	return prefixError(
		"FormIDEquals",
		Equals("document.querySelector('#root_id').value", expected))
}

func FormTextEquals(expected string) chromedp.Action {
	return prefixError(
		"FormTextEquals",
		Equals("document.querySelector('#root_text').value", expected))
}

func FormDoneEquals(expected bool) chromedp.Action {
	return prefixError(
		"FormDoneEquals",
		Equals("document.querySelector('#root_done').checked", expected))
}

func FormSubmit() chromedp.Action {
	return prefixError(
		"FormSubmit",
		chromedp.Click("*[type='submit']", chromedp.ByQuery))
}

func FormEnterID(text string) chromedp.Action {
	return prefixError(
		"FormEnterID",
		formEnter("root_id", text))
}

func FormEnterText(text string) chromedp.Action {
	return prefixError(
		"FormEnterText",
		formEnter("root_text", text))
}

func FormSetDone(checked bool) chromedp.Action {
	return prefixError(
		"FormSetDone",
		FormSetCheckBoxAction{
			checked: checked,
			id:      "root_done",
		})
}

type FormSetCheckBoxAction struct {
	checked bool
	id      string
}

func (f FormSetCheckBoxAction) Do(ctx context.Context) error {
	res, err := clickCheckbox(ctx, f.id)
	if err != nil {
		return err
	}

	if res == f.checked {
		return nil
	}

	res, err = clickCheckbox(ctx, f.id)
	if err != nil {
		return err
	}

	if res == f.checked {
		return nil
	}

	return fmt.Errorf("was not able to click checkbox to be checked=%v", f.checked)
}

func clickCheckbox(ctx context.Context, id string) (bool, error) {
	var sel = "#" + id

	err := chromedp.Click(sel, chromedp.ByID).Do(ctx)
	if err != nil {
		return false, err
	}

	var res bool

	err = chromedp.Evaluate("document.querySelector('"+sel+"').checked", &res, chromedp.EvalAsValue).Do(ctx)
	if err != nil {
		return false, err
	}

	return res, nil
}

func formEnter(id, text string) chromedp.Action {
	return formEnterTextInputAction{
		ID:    id,
		Input: text,
	}
}

type formEnterTextInputAction struct {
	ID    string
	Input string
}

func (f formEnterTextInputAction) Do(ctx context.Context) error {
	return chromedp.SendKeys("#"+f.ID, f.Input, chromedp.ByID).Do(ctx)
}
