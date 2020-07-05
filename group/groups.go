package group

import (
	"context"

	"github.com/chromedp/chromedp"
)

func New(title string, action ...chromedp.Action) chromedp.Action {
	return Action{
		title:        title,
		simpleAction: Simple(action...),
	}
}

type Action struct {
	title        string
	simpleAction SimpleAction
}

func (g Action) Do(ctx context.Context) error {
	err := Text(g.title).Do(ctx)
	if err != nil {
		return err
	}

	return g.simpleAction.Do(ctx)
}

func Simple(action ...chromedp.Action) SimpleAction {
	return SimpleAction{
		actions: action,
	}
}

type SimpleAction struct {
	actions []chromedp.Action
}

func (s SimpleAction) Do(ctx context.Context) error {
	var err error

	for _, action := range s.actions {
		err = action.Do(ctx)
		if err != nil {
			break
		}
	}

	return err
}
