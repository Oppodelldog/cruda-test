package cruda

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/chromedp/chromedp"
)

type testItem struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func InitAdapterFixtures() chromedp.Action {
	return prefixError("InitAdapterFixtures", InitAdapterFixturesAction{
		items: []testItem{
			{
				ID:   "1",
				Text: "Entry 1",
				Done: false,
			},
			{
				ID:   "2",
				Text: "Entry 2",
				Done: true,
			},
		},
	})
}

type InitAdapterFixturesAction struct {
	items []testItem
}

func (i InitAdapterFixturesAction) Do(ctx context.Context) error {
	var res bool

	jsonItems, err := json.Marshal(i.items)
	if err != nil {
		return err
	}

	expression := "window.testAdapter.setItems(" + string(jsonItems) + ")"

	err = chromedp.Evaluate(expression, &res, chromedp.EvalAsValue).Do(ctx)
	if err != nil {
		return err
	}

	if !res {
		return fmt.Errorf("init fixtures failed, testAdapter returned %v", res)
	}

	return nil
}

func InitAdapterCreationError(errorValue string) chromedp.Action {
	return prefixError("InitAdapterCreationError", InitAdapterErrorAction{
		operationName: "setCreateOperationError",
		errorValue:    errorValue,
	})
}

func InitAdapterUpdateError(errorValue string) chromedp.Action {
	return prefixError("InitAdapterUpdateError", InitAdapterErrorAction{
		operationName: "setUpdateOperationError",
		errorValue:    errorValue,
	})
}

func InitAdapterDeleteError(errorValue string) chromedp.Action {
	return prefixError("InitAdapterDeleteError", InitAdapterErrorAction{
		operationName: "setDeleteOperationError",
		errorValue:    errorValue,
	})
}

func InitAdapterLoadItemError(errorValue string) chromedp.Action {
	return prefixError("InitAdapterLoadItemError", InitAdapterErrorAction{
		operationName: "setLoadItemOperationError",
		errorValue:    errorValue,
	})
}

type InitAdapterErrorAction struct {
	operationName string
	errorValue    string
}

func (i InitAdapterErrorAction) Do(ctx context.Context) error {
	var res bool

	expression := "window.testAdapter." + i.operationName + "('" + i.errorValue + "')"

	err := chromedp.Evaluate(expression, &res, chromedp.EvalAsValue).Do(ctx)
	if err != nil {
		return err
	}

	if !res {
		return fmt.Errorf("init adapter error failed, testAdapter returned %v", res)
	}

	return nil
}
