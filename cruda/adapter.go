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

func InitAdapterFixtures(adapterID string) chromedp.Action {
	action := initAdapterFixtures()
	action.testAdapterID = adapterID

	return action
}

func initAdapterFixtures() InitAdapterFixturesAction {
	return InitAdapterFixturesAction{
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
	}
}

type InitAdapterFixturesAction struct {
	testAdapterID string
	items         []testItem
}

func (i InitAdapterFixturesAction) Do(ctx context.Context) error {
	var res bool

	jsonItems, err := json.Marshal(i.items)
	if err != nil {
		return err
	}

	expression := "window.testAdapter" + i.testAdapterID + ".setItems(" + string(jsonItems) + ")"

	err = chromedp.Evaluate(expression, &res, chromedp.EvalAsValue).Do(ctx)
	if err != nil {
		return err
	}

	if !res {
		return fmt.Errorf("init fixtures failed, testAdapter returned %v", res)
	}

	return nil
}

func InitAdapterCreationError(adapterID string, errorValue string) chromedp.Action {
	return prefixError("InitAdapterCreationError", InitAdapterErrorAction{
		testAdapterID: adapterID,
		operationName: "setCreateOperationError",
		errorValue:    errorValue,
	})
}

func InitAdapterUpdateError(adapterID string, errorValue string) chromedp.Action {
	return prefixError("InitAdapterUpdateError", InitAdapterErrorAction{
		testAdapterID: adapterID,
		operationName: "setUpdateOperationError",
		errorValue:    errorValue,
	})
}

func InitAdapterDeleteError(adapterID string, errorValue string) chromedp.Action {
	return prefixError("InitAdapterDeleteError", InitAdapterErrorAction{
		testAdapterID: adapterID,
		operationName: "setDeleteOperationError",
		errorValue:    errorValue,
	})
}

func InitAdapterLoadItemError(adapterID string, errorValue string) chromedp.Action {
	return prefixError("InitAdapterLoadItemError", InitAdapterErrorAction{
		testAdapterID: adapterID,
		operationName: "setLoadItemOperationError",
		errorValue:    errorValue,
	})
}

type InitAdapterErrorAction struct {
	testAdapterID string
	operationName string
	errorValue    string
}

func (i InitAdapterErrorAction) Do(ctx context.Context) error {
	var res bool

	expression := "window.testAdapter" + i.testAdapterID + "." + i.operationName + "('" + i.errorValue + "')"

	err := chromedp.Evaluate(expression, &res, chromedp.EvalAsValue).Do(ctx)
	if err != nil {
		return err
	}

	if !res {
		return fmt.Errorf("init adapter error failed, testAdapter returned %v", res)
	}

	return nil
}
